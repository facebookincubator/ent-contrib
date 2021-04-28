// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/contrib/entproto/cmd/protoc-gen-ent/internal/todo/ent/attachment"
	"entgo.io/ent/dialect/sql"
)

// Attachment is the model entity for the Attachment schema.
type Attachment struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Contents holds the value of the "contents" field.
	Contents string `json:"contents,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Attachment) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case attachment.FieldID, attachment.FieldContents:
			values[i] = &sql.NullString{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Attachment", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Attachment fields.
func (a *Attachment) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case attachment.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				a.ID = value.String
			}
		case attachment.FieldContents:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field contents", values[i])
			} else if value.Valid {
				a.Contents = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Attachment.
// Note that you need to call Attachment.Unwrap() before calling this method if this Attachment
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Attachment) Update() *AttachmentUpdateOne {
	return (&AttachmentClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Attachment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Attachment) Unwrap() *Attachment {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Attachment is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Attachment) String() string {
	var builder strings.Builder
	builder.WriteString("Attachment(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", contents=")
	builder.WriteString(a.Contents)
	builder.WriteByte(')')
	return builder.String()
}

// Attachments is a parsable slice of Attachment.
type Attachments []*Attachment

func (a Attachments) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
