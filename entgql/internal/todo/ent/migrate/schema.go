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
//
// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "text", Type: field.TypeString, Size: 2147483647},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"ENABLED", "DISABLED"}},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
	}
	// TodosColumns holds the columns for the "todos" table.
	TodosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"IN_PROGRESS", "COMPLETED"}},
		{Name: "priority", Type: field.TypeInt, Default: 0},
		{Name: "text", Type: field.TypeString, Size: 2147483647},
		{Name: "blob", Type: field.TypeBytes, Nullable: true},
		{Name: "category_todos", Type: field.TypeInt, Nullable: true},
		{Name: "todo_children", Type: field.TypeInt, Nullable: true},
		{Name: "todo_secret", Type: field.TypeInt, Nullable: true},
	}
	// TodosTable holds the schema information for the "todos" table.
	TodosTable = &schema.Table{
		Name:       "todos",
		Columns:    TodosColumns,
		PrimaryKey: []*schema.Column{TodosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "todos_categories_todos",
				Columns:    []*schema.Column{TodosColumns[6]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "todos_todos_children",
				Columns:    []*schema.Column{TodosColumns[7]},
				RefColumns: []*schema.Column{TodosColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "todos_very_secrets_secret",
				Columns:    []*schema.Column{TodosColumns[8]},
				RefColumns: []*schema.Column{VerySecretsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// VerySecretsColumns holds the columns for the "very_secrets" table.
	VerySecretsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "password", Type: field.TypeString},
	}
	// VerySecretsTable holds the schema information for the "very_secrets" table.
	VerySecretsTable = &schema.Table{
		Name:       "very_secrets",
		Columns:    VerySecretsColumns,
		PrimaryKey: []*schema.Column{VerySecretsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CategoriesTable,
		TodosTable,
		VerySecretsTable,
	}
)

func init() {
	TodosTable.ForeignKeys[0].RefTable = CategoriesTable
	TodosTable.ForeignKeys[1].RefTable = TodosTable
	TodosTable.ForeignKeys[2].RefTable = VerySecretsTable
}
