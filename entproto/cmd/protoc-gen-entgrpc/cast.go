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
	"fmt"

	"entgo.io/contrib/entproto"
	"google.golang.org/protobuf/compiler/protogen"
	dpb "google.golang.org/protobuf/types/descriptorpb"
)

func (g *serviceGenerator) castToProtoFunc(fld *entproto.FieldMappingDescriptor) (interface{}, error) {
	// TODO(rotemtam): don't wrap if the ent type == the pb type
	pbd := fld.PbFieldDescriptor
	switch pbd.GetType() {
	case dpb.FieldDescriptorProto_TYPE_INT32:
		return "int32", nil
	case dpb.FieldDescriptorProto_TYPE_INT64:
		return "int64", nil
	case dpb.FieldDescriptorProto_TYPE_STRING:
		return "string", nil
	case dpb.FieldDescriptorProto_TYPE_UINT32:
		return "uint32", nil
	case dpb.FieldDescriptorProto_TYPE_UINT64:
		return "uint64", nil
	case dpb.FieldDescriptorProto_TYPE_ENUM:
		ident := g.pbEnumIdent(fld)
		methodName := "toProto" + ident.GoName
		return methodName, nil
	case dpb.FieldDescriptorProto_TYPE_MESSAGE:
		if pbd.GetMessageType().GetFullyQualifiedName() == "google.protobuf.Timestamp" {
			newTS := protogen.GoImportPath("google.golang.org/protobuf/types/known/timestamppb").Ident("New")
			return newTS, nil
		} else {
			return nil, fmt.Errorf("entproto: no mapping for pb message type %q", pbd.GetFullyQualifiedName())
		}
	default:
		return nil, fmt.Errorf("entproto: no mapping for pb field type %q", pbd.GetType())
	}
}

func (g *serviceGenerator) castToEntFunc(fld *entproto.FieldMappingDescriptor) (interface{}, error) {
	et := fld.EntField
	switch {
	case et.IsBool(), et.IsBytes(), et.IsString(), et.Type.Numeric():
		return et.Type.String(), nil
	case et.IsTime():
		return protogen.GoImportPath("entgo.io/contrib/entproto").Ident("ExtractTime"), nil
	case et.IsEnum():
		ident := g.pbEnumIdent(fld)
		methodName := "toEnt" + ident.GoName
		return methodName, nil
	// case field.TypeJSON:
	// case field.TypeUUID:
	// case field.TypeOther:
	default:
		return nil, fmt.Errorf("entproto: no mapping to ent field type %q", et.Type.Type.ConstName())
	}
}