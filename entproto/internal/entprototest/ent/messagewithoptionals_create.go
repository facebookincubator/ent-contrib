// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/contrib/entproto/internal/entprototest/ent/messagewithoptionals"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// MessageWithOptionalsCreate is the builder for creating a MessageWithOptionals entity.
type MessageWithOptionalsCreate struct {
	config
	mutation *MessageWithOptionalsMutation
	hooks    []Hook
}

// SetStrOptional sets the "str_optional" field.
func (mwoc *MessageWithOptionalsCreate) SetStrOptional(s string) *MessageWithOptionalsCreate {
	mwoc.mutation.SetStrOptional(s)
	return mwoc
}

// SetNillableStrOptional sets the "str_optional" field if the given value is not nil.
func (mwoc *MessageWithOptionalsCreate) SetNillableStrOptional(s *string) *MessageWithOptionalsCreate {
	if s != nil {
		mwoc.SetStrOptional(*s)
	}
	return mwoc
}

// SetIntOptional sets the "int_optional" field.
func (mwoc *MessageWithOptionalsCreate) SetIntOptional(i int8) *MessageWithOptionalsCreate {
	mwoc.mutation.SetIntOptional(i)
	return mwoc
}

// SetNillableIntOptional sets the "int_optional" field if the given value is not nil.
func (mwoc *MessageWithOptionalsCreate) SetNillableIntOptional(i *int8) *MessageWithOptionalsCreate {
	if i != nil {
		mwoc.SetIntOptional(*i)
	}
	return mwoc
}

// SetUintOptional sets the "uint_optional" field.
func (mwoc *MessageWithOptionalsCreate) SetUintOptional(u uint8) *MessageWithOptionalsCreate {
	mwoc.mutation.SetUintOptional(u)
	return mwoc
}

// SetNillableUintOptional sets the "uint_optional" field if the given value is not nil.
func (mwoc *MessageWithOptionalsCreate) SetNillableUintOptional(u *uint8) *MessageWithOptionalsCreate {
	if u != nil {
		mwoc.SetUintOptional(*u)
	}
	return mwoc
}

// SetFloatOptional sets the "float_optional" field.
func (mwoc *MessageWithOptionalsCreate) SetFloatOptional(f float32) *MessageWithOptionalsCreate {
	mwoc.mutation.SetFloatOptional(f)
	return mwoc
}

// SetNillableFloatOptional sets the "float_optional" field if the given value is not nil.
func (mwoc *MessageWithOptionalsCreate) SetNillableFloatOptional(f *float32) *MessageWithOptionalsCreate {
	if f != nil {
		mwoc.SetFloatOptional(*f)
	}
	return mwoc
}

// SetBoolOptional sets the "bool_optional" field.
func (mwoc *MessageWithOptionalsCreate) SetBoolOptional(b bool) *MessageWithOptionalsCreate {
	mwoc.mutation.SetBoolOptional(b)
	return mwoc
}

// SetNillableBoolOptional sets the "bool_optional" field if the given value is not nil.
func (mwoc *MessageWithOptionalsCreate) SetNillableBoolOptional(b *bool) *MessageWithOptionalsCreate {
	if b != nil {
		mwoc.SetBoolOptional(*b)
	}
	return mwoc
}

// SetBytesOptional sets the "bytes_optional" field.
func (mwoc *MessageWithOptionalsCreate) SetBytesOptional(b []byte) *MessageWithOptionalsCreate {
	mwoc.mutation.SetBytesOptional(b)
	return mwoc
}

// SetUUIDOptional sets the "uuid_optional" field.
func (mwoc *MessageWithOptionalsCreate) SetUUIDOptional(u uuid.UUID) *MessageWithOptionalsCreate {
	mwoc.mutation.SetUUIDOptional(u)
	return mwoc
}

// SetTimeOptional sets the "time_optional" field.
func (mwoc *MessageWithOptionalsCreate) SetTimeOptional(t time.Time) *MessageWithOptionalsCreate {
	mwoc.mutation.SetTimeOptional(t)
	return mwoc
}

// SetNillableTimeOptional sets the "time_optional" field if the given value is not nil.
func (mwoc *MessageWithOptionalsCreate) SetNillableTimeOptional(t *time.Time) *MessageWithOptionalsCreate {
	if t != nil {
		mwoc.SetTimeOptional(*t)
	}
	return mwoc
}

// Mutation returns the MessageWithOptionalsMutation object of the builder.
func (mwoc *MessageWithOptionalsCreate) Mutation() *MessageWithOptionalsMutation {
	return mwoc.mutation
}

// Save creates the MessageWithOptionals in the database.
func (mwoc *MessageWithOptionalsCreate) Save(ctx context.Context) (*MessageWithOptionals, error) {
	var (
		err  error
		node *MessageWithOptionals
	)
	if len(mwoc.hooks) == 0 {
		if err = mwoc.check(); err != nil {
			return nil, err
		}
		node, err = mwoc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageWithOptionalsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mwoc.check(); err != nil {
				return nil, err
			}
			mwoc.mutation = mutation
			if node, err = mwoc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mwoc.hooks) - 1; i >= 0; i-- {
			mut = mwoc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mwoc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mwoc *MessageWithOptionalsCreate) SaveX(ctx context.Context) *MessageWithOptionals {
	v, err := mwoc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (mwoc *MessageWithOptionalsCreate) check() error {
	return nil
}

func (mwoc *MessageWithOptionalsCreate) sqlSave(ctx context.Context) (*MessageWithOptionals, error) {
	_node, _spec := mwoc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mwoc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (mwoc *MessageWithOptionalsCreate) createSpec() (*MessageWithOptionals, *sqlgraph.CreateSpec) {
	var (
		_node = &MessageWithOptionals{config: mwoc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: messagewithoptionals.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: messagewithoptionals.FieldID,
			},
		}
	)
	if value, ok := mwoc.mutation.StrOptional(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: messagewithoptionals.FieldStrOptional,
		})
		_node.StrOptional = value
	}
	if value, ok := mwoc.mutation.IntOptional(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: messagewithoptionals.FieldIntOptional,
		})
		_node.IntOptional = value
	}
	if value, ok := mwoc.mutation.UintOptional(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: messagewithoptionals.FieldUintOptional,
		})
		_node.UintOptional = value
	}
	if value, ok := mwoc.mutation.FloatOptional(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: messagewithoptionals.FieldFloatOptional,
		})
		_node.FloatOptional = value
	}
	if value, ok := mwoc.mutation.BoolOptional(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: messagewithoptionals.FieldBoolOptional,
		})
		_node.BoolOptional = value
	}
	if value, ok := mwoc.mutation.BytesOptional(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: messagewithoptionals.FieldBytesOptional,
		})
		_node.BytesOptional = value
	}
	if value, ok := mwoc.mutation.UUIDOptional(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: messagewithoptionals.FieldUUIDOptional,
		})
		_node.UUIDOptional = value
	}
	if value, ok := mwoc.mutation.TimeOptional(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: messagewithoptionals.FieldTimeOptional,
		})
		_node.TimeOptional = value
	}
	return _node, _spec
}

// MessageWithOptionalsCreateBulk is the builder for creating many MessageWithOptionals entities in bulk.
type MessageWithOptionalsCreateBulk struct {
	config
	builders []*MessageWithOptionalsCreate
}

// Save creates the MessageWithOptionals entities in the database.
func (mwocb *MessageWithOptionalsCreateBulk) Save(ctx context.Context) ([]*MessageWithOptionals, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mwocb.builders))
	nodes := make([]*MessageWithOptionals, len(mwocb.builders))
	mutators := make([]Mutator, len(mwocb.builders))
	for i := range mwocb.builders {
		func(i int, root context.Context) {
			builder := mwocb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MessageWithOptionalsMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mwocb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mwocb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mwocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mwocb *MessageWithOptionalsCreateBulk) SaveX(ctx context.Context) []*MessageWithOptionals {
	v, err := mwocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
