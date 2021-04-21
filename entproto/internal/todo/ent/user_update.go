// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/contrib/entproto/internal/todo/ent/attachment"
	"entgo.io/contrib/entproto/internal/todo/ent/group"
	"entgo.io/contrib/entproto/internal/todo/ent/predicate"
	"entgo.io/contrib/entproto/internal/todo/ent/user"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where adds a new predicate for the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.predicates = append(uu.mutation.predicates, ps...)
	return uu
}

// SetUserName sets the "user_name" field.
func (uu *UserUpdate) SetUserName(s string) *UserUpdate {
	uu.mutation.SetUserName(s)
	return uu
}

// SetPoints sets the "points" field.
func (uu *UserUpdate) SetPoints(u uint) *UserUpdate {
	uu.mutation.ResetPoints()
	uu.mutation.SetPoints(u)
	return uu
}

// AddPoints adds u to the "points" field.
func (uu *UserUpdate) AddPoints(u uint) *UserUpdate {
	uu.mutation.AddPoints(u)
	return uu
}

// SetExp sets the "exp" field.
func (uu *UserUpdate) SetExp(u uint64) *UserUpdate {
	uu.mutation.ResetExp()
	uu.mutation.SetExp(u)
	return uu
}

// AddExp adds u to the "exp" field.
func (uu *UserUpdate) AddExp(u uint64) *UserUpdate {
	uu.mutation.AddExp(u)
	return uu
}

// SetStatus sets the "status" field.
func (uu *UserUpdate) SetStatus(u user.Status) *UserUpdate {
	uu.mutation.SetStatus(u)
	return uu
}

// SetExternalID sets the "external_id" field.
func (uu *UserUpdate) SetExternalID(i int) *UserUpdate {
	uu.mutation.ResetExternalID()
	uu.mutation.SetExternalID(i)
	return uu
}

// AddExternalID adds i to the "external_id" field.
func (uu *UserUpdate) AddExternalID(i int) *UserUpdate {
	uu.mutation.AddExternalID(i)
	return uu
}

// SetCrmID sets the "crm_id" field.
func (uu *UserUpdate) SetCrmID(u uuid.UUID) *UserUpdate {
	uu.mutation.SetCrmID(u)
	return uu
}

// SetBanned sets the "banned" field.
func (uu *UserUpdate) SetBanned(b bool) *UserUpdate {
	uu.mutation.SetBanned(b)
	return uu
}

// SetNillableBanned sets the "banned" field if the given value is not nil.
func (uu *UserUpdate) SetNillableBanned(b *bool) *UserUpdate {
	if b != nil {
		uu.SetBanned(*b)
	}
	return uu
}

// SetCustomPb sets the "custom_pb" field.
func (uu *UserUpdate) SetCustomPb(u uint8) *UserUpdate {
	uu.mutation.ResetCustomPb()
	uu.mutation.SetCustomPb(u)
	return uu
}

// AddCustomPb adds u to the "custom_pb" field.
func (uu *UserUpdate) AddCustomPb(u uint8) *UserUpdate {
	uu.mutation.AddCustomPb(u)
	return uu
}

// SetOptNum sets the "opt_num" field.
func (uu *UserUpdate) SetOptNum(i int) *UserUpdate {
	uu.mutation.ResetOptNum()
	uu.mutation.SetOptNum(i)
	return uu
}

// SetNillableOptNum sets the "opt_num" field if the given value is not nil.
func (uu *UserUpdate) SetNillableOptNum(i *int) *UserUpdate {
	if i != nil {
		uu.SetOptNum(*i)
	}
	return uu
}

// AddOptNum adds i to the "opt_num" field.
func (uu *UserUpdate) AddOptNum(i int) *UserUpdate {
	uu.mutation.AddOptNum(i)
	return uu
}

// ClearOptNum clears the value of the "opt_num" field.
func (uu *UserUpdate) ClearOptNum() *UserUpdate {
	uu.mutation.ClearOptNum()
	return uu
}

// SetOptStr sets the "opt_str" field.
func (uu *UserUpdate) SetOptStr(s string) *UserUpdate {
	uu.mutation.SetOptStr(s)
	return uu
}

// SetNillableOptStr sets the "opt_str" field if the given value is not nil.
func (uu *UserUpdate) SetNillableOptStr(s *string) *UserUpdate {
	if s != nil {
		uu.SetOptStr(*s)
	}
	return uu
}

// ClearOptStr clears the value of the "opt_str" field.
func (uu *UserUpdate) ClearOptStr() *UserUpdate {
	uu.mutation.ClearOptStr()
	return uu
}

// SetOptBool sets the "opt_bool" field.
func (uu *UserUpdate) SetOptBool(s string) *UserUpdate {
	uu.mutation.SetOptBool(s)
	return uu
}

// SetNillableOptBool sets the "opt_bool" field if the given value is not nil.
func (uu *UserUpdate) SetNillableOptBool(s *string) *UserUpdate {
	if s != nil {
		uu.SetOptBool(*s)
	}
	return uu
}

// ClearOptBool clears the value of the "opt_bool" field.
func (uu *UserUpdate) ClearOptBool() *UserUpdate {
	uu.mutation.ClearOptBool()
	return uu
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (uu *UserUpdate) SetGroupID(id int) *UserUpdate {
	uu.mutation.SetGroupID(id)
	return uu
}

// SetNillableGroupID sets the "group" edge to the Group entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableGroupID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetGroupID(*id)
	}
	return uu
}

// SetGroup sets the "group" edge to the Group entity.
func (uu *UserUpdate) SetGroup(g *Group) *UserUpdate {
	return uu.SetGroupID(g.ID)
}

// SetAttachmentID sets the "attachment" edge to the Attachment entity by ID.
func (uu *UserUpdate) SetAttachmentID(id uuid.UUID) *UserUpdate {
	uu.mutation.SetAttachmentID(id)
	return uu
}

// SetNillableAttachmentID sets the "attachment" edge to the Attachment entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableAttachmentID(id *uuid.UUID) *UserUpdate {
	if id != nil {
		uu = uu.SetAttachmentID(*id)
	}
	return uu
}

// SetAttachment sets the "attachment" edge to the Attachment entity.
func (uu *UserUpdate) SetAttachment(a *Attachment) *UserUpdate {
	return uu.SetAttachmentID(a.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearGroup clears the "group" edge to the Group entity.
func (uu *UserUpdate) ClearGroup() *UserUpdate {
	uu.mutation.ClearGroup()
	return uu
}

// ClearAttachment clears the "attachment" edge to the Attachment entity.
func (uu *UserUpdate) ClearAttachment() *UserUpdate {
	uu.mutation.ClearAttachment()
	return uu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		if err = uu.check(); err != nil {
			return 0, err
		}
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uu.check(); err != nil {
				return 0, err
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Status(); ok {
		if err := user.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UserName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUserName,
		})
	}
	if value, ok := uu.mutation.Points(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: user.FieldPoints,
		})
	}
	if value, ok := uu.mutation.AddedPoints(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: user.FieldPoints,
		})
	}
	if value, ok := uu.mutation.Exp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: user.FieldExp,
		})
	}
	if value, ok := uu.mutation.AddedExp(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: user.FieldExp,
		})
	}
	if value, ok := uu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: user.FieldStatus,
		})
	}
	if value, ok := uu.mutation.ExternalID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldExternalID,
		})
	}
	if value, ok := uu.mutation.AddedExternalID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldExternalID,
		})
	}
	if value, ok := uu.mutation.CrmID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: user.FieldCrmID,
		})
	}
	if value, ok := uu.mutation.Banned(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldBanned,
		})
	}
	if value, ok := uu.mutation.CustomPb(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: user.FieldCustomPb,
		})
	}
	if value, ok := uu.mutation.AddedCustomPb(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: user.FieldCustomPb,
		})
	}
	if value, ok := uu.mutation.OptNum(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldOptNum,
		})
	}
	if value, ok := uu.mutation.AddedOptNum(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldOptNum,
		})
	}
	if uu.mutation.OptNumCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: user.FieldOptNum,
		})
	}
	if value, ok := uu.mutation.OptStr(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldOptStr,
		})
	}
	if uu.mutation.OptStrCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldOptStr,
		})
	}
	if value, ok := uu.mutation.OptBool(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldOptBool,
		})
	}
	if uu.mutation.OptBoolCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldOptBool,
		})
	}
	if uu.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.GroupTable,
			Columns: []string{user.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.GroupTable,
			Columns: []string{user.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.AttachmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.AttachmentTable,
			Columns: []string{user.AttachmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: attachment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.AttachmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.AttachmentTable,
			Columns: []string{user.AttachmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: attachment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// SetUserName sets the "user_name" field.
func (uuo *UserUpdateOne) SetUserName(s string) *UserUpdateOne {
	uuo.mutation.SetUserName(s)
	return uuo
}

// SetPoints sets the "points" field.
func (uuo *UserUpdateOne) SetPoints(u uint) *UserUpdateOne {
	uuo.mutation.ResetPoints()
	uuo.mutation.SetPoints(u)
	return uuo
}

// AddPoints adds u to the "points" field.
func (uuo *UserUpdateOne) AddPoints(u uint) *UserUpdateOne {
	uuo.mutation.AddPoints(u)
	return uuo
}

// SetExp sets the "exp" field.
func (uuo *UserUpdateOne) SetExp(u uint64) *UserUpdateOne {
	uuo.mutation.ResetExp()
	uuo.mutation.SetExp(u)
	return uuo
}

// AddExp adds u to the "exp" field.
func (uuo *UserUpdateOne) AddExp(u uint64) *UserUpdateOne {
	uuo.mutation.AddExp(u)
	return uuo
}

// SetStatus sets the "status" field.
func (uuo *UserUpdateOne) SetStatus(u user.Status) *UserUpdateOne {
	uuo.mutation.SetStatus(u)
	return uuo
}

// SetExternalID sets the "external_id" field.
func (uuo *UserUpdateOne) SetExternalID(i int) *UserUpdateOne {
	uuo.mutation.ResetExternalID()
	uuo.mutation.SetExternalID(i)
	return uuo
}

// AddExternalID adds i to the "external_id" field.
func (uuo *UserUpdateOne) AddExternalID(i int) *UserUpdateOne {
	uuo.mutation.AddExternalID(i)
	return uuo
}

// SetCrmID sets the "crm_id" field.
func (uuo *UserUpdateOne) SetCrmID(u uuid.UUID) *UserUpdateOne {
	uuo.mutation.SetCrmID(u)
	return uuo
}

// SetBanned sets the "banned" field.
func (uuo *UserUpdateOne) SetBanned(b bool) *UserUpdateOne {
	uuo.mutation.SetBanned(b)
	return uuo
}

// SetNillableBanned sets the "banned" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableBanned(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetBanned(*b)
	}
	return uuo
}

// SetCustomPb sets the "custom_pb" field.
func (uuo *UserUpdateOne) SetCustomPb(u uint8) *UserUpdateOne {
	uuo.mutation.ResetCustomPb()
	uuo.mutation.SetCustomPb(u)
	return uuo
}

// AddCustomPb adds u to the "custom_pb" field.
func (uuo *UserUpdateOne) AddCustomPb(u uint8) *UserUpdateOne {
	uuo.mutation.AddCustomPb(u)
	return uuo
}

// SetOptNum sets the "opt_num" field.
func (uuo *UserUpdateOne) SetOptNum(i int) *UserUpdateOne {
	uuo.mutation.ResetOptNum()
	uuo.mutation.SetOptNum(i)
	return uuo
}

// SetNillableOptNum sets the "opt_num" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableOptNum(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetOptNum(*i)
	}
	return uuo
}

// AddOptNum adds i to the "opt_num" field.
func (uuo *UserUpdateOne) AddOptNum(i int) *UserUpdateOne {
	uuo.mutation.AddOptNum(i)
	return uuo
}

// ClearOptNum clears the value of the "opt_num" field.
func (uuo *UserUpdateOne) ClearOptNum() *UserUpdateOne {
	uuo.mutation.ClearOptNum()
	return uuo
}

// SetOptStr sets the "opt_str" field.
func (uuo *UserUpdateOne) SetOptStr(s string) *UserUpdateOne {
	uuo.mutation.SetOptStr(s)
	return uuo
}

// SetNillableOptStr sets the "opt_str" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableOptStr(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetOptStr(*s)
	}
	return uuo
}

// ClearOptStr clears the value of the "opt_str" field.
func (uuo *UserUpdateOne) ClearOptStr() *UserUpdateOne {
	uuo.mutation.ClearOptStr()
	return uuo
}

// SetOptBool sets the "opt_bool" field.
func (uuo *UserUpdateOne) SetOptBool(s string) *UserUpdateOne {
	uuo.mutation.SetOptBool(s)
	return uuo
}

// SetNillableOptBool sets the "opt_bool" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableOptBool(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetOptBool(*s)
	}
	return uuo
}

// ClearOptBool clears the value of the "opt_bool" field.
func (uuo *UserUpdateOne) ClearOptBool() *UserUpdateOne {
	uuo.mutation.ClearOptBool()
	return uuo
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (uuo *UserUpdateOne) SetGroupID(id int) *UserUpdateOne {
	uuo.mutation.SetGroupID(id)
	return uuo
}

// SetNillableGroupID sets the "group" edge to the Group entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableGroupID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetGroupID(*id)
	}
	return uuo
}

// SetGroup sets the "group" edge to the Group entity.
func (uuo *UserUpdateOne) SetGroup(g *Group) *UserUpdateOne {
	return uuo.SetGroupID(g.ID)
}

// SetAttachmentID sets the "attachment" edge to the Attachment entity by ID.
func (uuo *UserUpdateOne) SetAttachmentID(id uuid.UUID) *UserUpdateOne {
	uuo.mutation.SetAttachmentID(id)
	return uuo
}

// SetNillableAttachmentID sets the "attachment" edge to the Attachment entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAttachmentID(id *uuid.UUID) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetAttachmentID(*id)
	}
	return uuo
}

// SetAttachment sets the "attachment" edge to the Attachment entity.
func (uuo *UserUpdateOne) SetAttachment(a *Attachment) *UserUpdateOne {
	return uuo.SetAttachmentID(a.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearGroup clears the "group" edge to the Group entity.
func (uuo *UserUpdateOne) ClearGroup() *UserUpdateOne {
	uuo.mutation.ClearGroup()
	return uuo
}

// ClearAttachment clears the "attachment" edge to the Attachment entity.
func (uuo *UserUpdateOne) ClearAttachment() *UserUpdateOne {
	uuo.mutation.ClearAttachment()
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if len(uuo.hooks) == 0 {
		if err = uuo.check(); err != nil {
			return nil, err
		}
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uuo.check(); err != nil {
				return nil, err
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Status(); ok {
		if err := user.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing User.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UserName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUserName,
		})
	}
	if value, ok := uuo.mutation.Points(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: user.FieldPoints,
		})
	}
	if value, ok := uuo.mutation.AddedPoints(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: user.FieldPoints,
		})
	}
	if value, ok := uuo.mutation.Exp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: user.FieldExp,
		})
	}
	if value, ok := uuo.mutation.AddedExp(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: user.FieldExp,
		})
	}
	if value, ok := uuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: user.FieldStatus,
		})
	}
	if value, ok := uuo.mutation.ExternalID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldExternalID,
		})
	}
	if value, ok := uuo.mutation.AddedExternalID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldExternalID,
		})
	}
	if value, ok := uuo.mutation.CrmID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: user.FieldCrmID,
		})
	}
	if value, ok := uuo.mutation.Banned(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldBanned,
		})
	}
	if value, ok := uuo.mutation.CustomPb(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: user.FieldCustomPb,
		})
	}
	if value, ok := uuo.mutation.AddedCustomPb(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: user.FieldCustomPb,
		})
	}
	if value, ok := uuo.mutation.OptNum(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldOptNum,
		})
	}
	if value, ok := uuo.mutation.AddedOptNum(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldOptNum,
		})
	}
	if uuo.mutation.OptNumCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: user.FieldOptNum,
		})
	}
	if value, ok := uuo.mutation.OptStr(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldOptStr,
		})
	}
	if uuo.mutation.OptStrCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldOptStr,
		})
	}
	if value, ok := uuo.mutation.OptBool(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldOptBool,
		})
	}
	if uuo.mutation.OptBoolCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldOptBool,
		})
	}
	if uuo.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.GroupTable,
			Columns: []string{user.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.GroupTable,
			Columns: []string{user.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.AttachmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.AttachmentTable,
			Columns: []string{user.AttachmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: attachment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.AttachmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.AttachmentTable,
			Columns: []string{user.AttachmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: attachment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
