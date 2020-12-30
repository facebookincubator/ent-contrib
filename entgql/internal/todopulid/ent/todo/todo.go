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

package todo

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

const (
	// Label holds the string label denoting the todo type in the database.
	Label = "todo"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldPriority holds the string denoting the priority field in the database.
	FieldPriority = "priority"
	// FieldText holds the string denoting the text field in the database.
	FieldText = "text"

	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"

	// Table holds the table name of the todo in the database.
	Table = "todos"
	// ParentTable is the table the holds the parent relation/edge.
	ParentTable = "todos"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "todo_children"
	// ChildrenTable is the table the holds the children relation/edge.
	ChildrenTable = "todos"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "todo_children"
)

// Columns holds all SQL columns for todo fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldStatus,
	FieldPriority,
	FieldText,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Todo type.
var ForeignKeys = []string{
	"todo_children",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
	// DefaultPriority holds the default value on creation for the priority field.
	DefaultPriority int
	// TextValidator is a validator for the "text" field. It is called by the builders before save.
	TextValidator func(string) error
	// DefaultID holds the default value on creation for the id field.
	DefaultID func() string
)

// Status defines the type for the status enum field.
type Status string

// Status values.
const (
	StatusInProgress Status = "IN_PROGRESS"
	StatusCompleted  Status = "COMPLETED"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusInProgress, StatusCompleted:
		return nil
	default:
		return fmt.Errorf("todo: invalid enum value for status field: %q", s)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (s Status) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(s.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (s *Status) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*s = Status(str)
	if err := StatusValidator(*s); err != nil {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}
