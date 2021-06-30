// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/contrib/schemast/internal/mutatetest/ent/predicate"
	"entgo.io/contrib/schemast/internal/mutatetest/ent/withfields"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WithFieldsUpdate is the builder for updating WithFields entities.
type WithFieldsUpdate struct {
	config
	hooks    []Hook
	mutation *WithFieldsMutation
}

// Where adds a new predicate for the WithFieldsUpdate builder.
func (wfu *WithFieldsUpdate) Where(ps ...predicate.WithFields) *WithFieldsUpdate {
	wfu.mutation.predicates = append(wfu.mutation.predicates, ps...)
	return wfu
}

// SetExisting sets the "existing" field.
func (wfu *WithFieldsUpdate) SetExisting(s string) *WithFieldsUpdate {
	wfu.mutation.SetExisting(s)
	return wfu
}

// Mutation returns the WithFieldsMutation object of the builder.
func (wfu *WithFieldsUpdate) Mutation() *WithFieldsMutation {
	return wfu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wfu *WithFieldsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wfu.hooks) == 0 {
		affected, err = wfu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WithFieldsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wfu.mutation = mutation
			affected, err = wfu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wfu.hooks) - 1; i >= 0; i-- {
			mut = wfu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wfu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wfu *WithFieldsUpdate) SaveX(ctx context.Context) int {
	affected, err := wfu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wfu *WithFieldsUpdate) Exec(ctx context.Context) error {
	_, err := wfu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wfu *WithFieldsUpdate) ExecX(ctx context.Context) {
	if err := wfu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wfu *WithFieldsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   withfields.Table,
			Columns: withfields.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: withfields.FieldID,
			},
		},
	}
	if ps := wfu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wfu.mutation.Existing(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: withfields.FieldExisting,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wfu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{withfields.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// WithFieldsUpdateOne is the builder for updating a single WithFields entity.
type WithFieldsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WithFieldsMutation
}

// SetExisting sets the "existing" field.
func (wfuo *WithFieldsUpdateOne) SetExisting(s string) *WithFieldsUpdateOne {
	wfuo.mutation.SetExisting(s)
	return wfuo
}

// Mutation returns the WithFieldsMutation object of the builder.
func (wfuo *WithFieldsUpdateOne) Mutation() *WithFieldsMutation {
	return wfuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wfuo *WithFieldsUpdateOne) Select(field string, fields ...string) *WithFieldsUpdateOne {
	wfuo.fields = append([]string{field}, fields...)
	return wfuo
}

// Save executes the query and returns the updated WithFields entity.
func (wfuo *WithFieldsUpdateOne) Save(ctx context.Context) (*WithFields, error) {
	var (
		err  error
		node *WithFields
	)
	if len(wfuo.hooks) == 0 {
		node, err = wfuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WithFieldsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wfuo.mutation = mutation
			node, err = wfuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wfuo.hooks) - 1; i >= 0; i-- {
			mut = wfuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wfuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wfuo *WithFieldsUpdateOne) SaveX(ctx context.Context) *WithFields {
	node, err := wfuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wfuo *WithFieldsUpdateOne) Exec(ctx context.Context) error {
	_, err := wfuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wfuo *WithFieldsUpdateOne) ExecX(ctx context.Context) {
	if err := wfuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wfuo *WithFieldsUpdateOne) sqlSave(ctx context.Context) (_node *WithFields, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   withfields.Table,
			Columns: withfields.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: withfields.FieldID,
			},
		},
	}
	id, ok := wfuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing WithFields.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := wfuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, withfields.FieldID)
		for _, f := range fields {
			if !withfields.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != withfields.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wfuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wfuo.mutation.Existing(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: withfields.FieldExisting,
		})
	}
	_node = &WithFields{config: wfuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wfuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{withfields.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
