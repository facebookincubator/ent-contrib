// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"

	"entgo.io/contrib/entproto"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	entSchemaPath *string
	snake         = gen.Funcs["snake"].(func(string) string)
	camel         = gen.Funcs["camel"].(func(string) string)
	contextImp    = protogen.GoImportPath("context")
	grpcStatImp   = protogen.GoImportPath("google.golang.org/grpc/status")
	codesImp      = protogen.GoImportPath("google.golang.org/grpc/codes")
)

func main() {
	var flags flag.FlagSet
	entSchemaPath = flags.String("schema_path", "", "ent schema path")
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(plg *protogen.Plugin) error {
		g, err := entc.LoadGraph(*entSchemaPath, &gen.Config{})
		if err != nil {
			return err
		}
		for _, f := range plg.Files {
			if !f.Generate {
				continue
			}
			if err := processFile(plg, f, g); err != nil {
				return err
			}
		}
		return nil
	})
}

// processFile generates service implementations from all services defined in the file.
func processFile(gen *protogen.Plugin, file *protogen.File, graph *gen.Graph) error {
	if len(file.Services) == 0 {
		return nil
	}
	for _, s := range file.Services {
		sg, err := newServiceGenerator(gen, file, graph, s)
		if err != nil {
			return err
		}
		if err := sg.generate(); err != nil {
			return err
		}
	}
	return nil
}

func newServiceGenerator(plugin *protogen.Plugin, file *protogen.File, graph *gen.Graph, service *protogen.Service) (*serviceGenerator, error) {
	adapter, err := entproto.LoadAdapter(graph)
	if err != nil {
		return nil, err
	}
	typeName, err := extractEntTypeName(service, graph)
	if err != nil {
		return nil, err
	}
	filename := file.GeneratedFilenamePrefix + "_" + snake(service.GoName) + ".go"
	g := plugin.NewGeneratedFile(filename, file.GoImportPath)
	fieldMap, err := adapter.FieldMap(typeName)
	if err != nil {
		return nil, err
	}
	return &serviceGenerator{
		GeneratedFile: g,
		entPackage:    protogen.GoImportPath(graph.Config.Package),
		file:          file,
		service:       service,
		typeName:      typeName,
		fieldMap:      fieldMap,
	}, nil
}

func (g *serviceGenerator) generate() error {
	g.P("// Code generated by protoc-gen-entgrpc. DO NOT EDIT.")
	g.P("package ", g.file.GoPackageName)
	g.P()
	g.generateConstructor()
	g.P()
	if err := g.generateToProtoMapper(); err != nil {
		return err
	}
	g.P()

	for _, method := range g.service.Methods {
		if err := g.generateMethod(method); err != nil {
			return err
		}
	}
	return nil
}

func (g *serviceGenerator) generateConstructor() {
	g.Tmpl(`
	// %(svcName) implements %(svcName)Server
	type %(svcName) struct {
		client *%(entClient)
		Unimplemented%(svcName)Server
	}

	// New%(svcName) returns a new %(svcName)
	func New%(svcName)(client *%(entClient)) *%(svcName) {
		return &%(svcName){
			client: client,
		}
	}`, tmplValues{
		"svcName":   g.service.GoName,
		"entClient": g.entPackage.Ident("Client"),
	})
}

func (g *serviceGenerator) generateToProtoMapper() error {
	// Mapper from the ent type to the proto type.
	g.Tmpl(`
	// toProto%(typeName) transforms the ent type to the pb type (TODO: complete implementation)
	func toProto%(typeName)(e *%(entTypeIdent)) *%(typeName){
		return &%(typeName) {`, tmplValues{
		"typeName":     g.typeName,
		"entTypeIdent": g.entPackage.Ident(g.typeName),
	})

	// TODO: impl in next PR
	castToProtoFunc := func(fld *entproto.FieldMappingDescriptor) (interface{}, error) {
		return "placeholder", nil
	}
	for _, fld := range g.fieldMap.Fields() {
		protoFunc, err := castToProtoFunc(fld)
		if err != nil {
			return err
		}
		g.Tmpl("// %(pbStructField): %(castFunc)(e.%(entStructField))", tmplValues{
			"pbStructField":  fld.PbStructField(),
			"entStructField": fld.EntField.StructField(),
			"castFunc":       protoFunc,
		})
	}
	g.Tmpl(`}
	}`, nil)
	return nil
}

type serviceGenerator struct {
	*protogen.GeneratedFile
	entPackage protogen.GoImportPath
	file       *protogen.File
	service    *protogen.Service
	typeName   string
	fieldMap   entproto.FieldMap
}

func (g *serviceGenerator) Tmpl(s string, values tmplValues) {
	if err := printTemplate(g.GeneratedFile, s, values); err != nil {
		panic(err)
	}
}

func (g *serviceGenerator) generateMethod(me *protogen.Method) error {
	g.Tmpl(`
	// %(name) implements %(svcName)Server.%(name)
	func (svc *%(svcName)) %(name)(ctx %(ctx), req *%(inputIdent)) (*%(outputIdent), error) {`, tmplValues{
		"name":        me.GoName,
		"svcName":     g.service.GoName,
		"ctx":         contextImp.Ident("Context"),
		"inputIdent":  me.Input.GoIdent,
		"outputIdent": me.Output.GoIdent,
	})

	switch me.GoName {
	// TODO: specific method implementations
	default:
		g.Tmpl(`return nil, %(grpcStatusError)(%(notImplemented), "error")`, tmplValues{
			"grpcStatusError": grpcStatImp.Ident("Error"),
			"notImplemented":  codesImp.Ident("Unimplemented"),
		})
	}
	g.P("}")
	return nil
}

func extractEntTypeName(s *protogen.Service, g *gen.Graph) (string, error) {
	typeName := s.GoName[0 : len(s.GoName)-len("Service")] // Get the type name from "UserService" -> "User"
	for _, gt := range g.Nodes {
		if gt.Name == typeName {
			return typeName, nil
		}
	}
	return "", fmt.Errorf("entproto: type %q of service %q not found in graph", typeName, s.GoName)
}
