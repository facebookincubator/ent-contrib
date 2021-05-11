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

package entgql

import (
	"fmt"
	"github.com/99designs/gqlgen/plugin"
	"github.com/vektah/gqlparser/v2/ast"
	"strings"
)

type entgqlgen struct {
	connections []string
}

var relayCoreTypes = `
scalar Cursor

interface Node {
    id: ID!
}

type PageInfo {
  hasNextPage: Boolean!
  hasPreviousPage: Boolean!
  startCursor: Cursor
  endCursor: Cursor
}
`

var connectionTemplate = `type %[1]sConnection {
    totalCount: Int!
    pageInfo: PageInfo!
    edges: [%[1]sEdge]
}

type %[1]sEdge {
    node: %[1]s
    cursor: Cursor!
}`

func generateConnectionType(name string) string {
	return fmt.Sprintf(connectionTemplate, name)
}

func (r *entgqlgen) InjectSourceEarly() *ast.Source {
	var sources = []string{relayCoreTypes}
	for _, c := range r.connections {
		sources = append(sources, generateConnectionType(c))
	}
	return &ast.Source{
		Name:    "entgqlgen/types.graphql",
		Input:   strings.Join(sources, "\n\n"),
		BuiltIn: true,
	}
}

func NewPlugin(connections []string) plugin.Plugin {
	return &entgqlgen{connections: connections}
}

func (r *entgqlgen) Name() string {
	return "entgqlgen"
}

var _ plugin.EarlySourceInjector = &entgqlgen{}