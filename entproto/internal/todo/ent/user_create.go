// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/contrib/entproto/internal/todo/ent/attachment"
	"entgo.io/contrib/entproto/internal/todo/ent/group"
	"entgo.io/contrib/entproto/internal/todo/ent/user"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetUserName sets the "user_name" field.
func (uc *UserCreate) SetUserName(s string) *UserCreate {
	uc.mutation.SetUserName(s)
	return uc
}

// SetJoined sets the "joined" field.
func (uc *UserCreate) SetJoined(t time.Time) *UserCreate {
	uc.mutation.SetJoined(t)
	return uc
}

// SetPoints sets the "points" field.
func (uc *UserCreate) SetPoints(u uint) *UserCreate {
	uc.mutation.SetPoints(u)
	return uc
}

// SetExp sets the "exp" field.
func (uc *UserCreate) SetExp(u uint64) *UserCreate {
	uc.mutation.SetExp(u)
	return uc
}

// SetStatus sets the "status" field.
func (uc *UserCreate) SetStatus(u user.Status) *UserCreate {
	uc.mutation.SetStatus(u)
	return uc
}

// SetExternalID sets the "external_id" field.
func (uc *UserCreate) SetExternalID(i int) *UserCreate {
	uc.mutation.SetExternalID(i)
	return uc
}

// SetCrmID sets the "crm_id" field.
func (uc *UserCreate) SetCrmID(u uuid.UUID) *UserCreate {
	uc.mutation.SetCrmID(u)
	return uc
}

// SetBanned sets the "banned" field.
func (uc *UserCreate) SetBanned(b bool) *UserCreate {
	uc.mutation.SetBanned(b)
	return uc
}

// SetNillableBanned sets the "banned" field if the given value is not nil.
func (uc *UserCreate) SetNillableBanned(b *bool) *UserCreate {
	if b != nil {
		uc.SetBanned(*b)
	}
	return uc
}

// SetCustomPb sets the "custom_pb" field.
func (uc *UserCreate) SetCustomPb(u uint8) *UserCreate {
	uc.mutation.SetCustomPb(u)
	return uc
}

// SetOptNum sets the "opt_num" field.
func (uc *UserCreate) SetOptNum(i int) *UserCreate {
	uc.mutation.SetOptNum(i)
	return uc
}

// SetNillableOptNum sets the "opt_num" field if the given value is not nil.
func (uc *UserCreate) SetNillableOptNum(i *int) *UserCreate {
	if i != nil {
		uc.SetOptNum(*i)
	}
	return uc
}

// SetOptStr sets the "opt_str" field.
func (uc *UserCreate) SetOptStr(s string) *UserCreate {
	uc.mutation.SetOptStr(s)
	return uc
}

// SetNillableOptStr sets the "opt_str" field if the given value is not nil.
func (uc *UserCreate) SetNillableOptStr(s *string) *UserCreate {
	if s != nil {
		uc.SetOptStr(*s)
	}
	return uc
}

// SetOptBool sets the "opt_bool" field.
func (uc *UserCreate) SetOptBool(b bool) *UserCreate {
	uc.mutation.SetOptBool(b)
	return uc
}

// SetNillableOptBool sets the "opt_bool" field if the given value is not nil.
func (uc *UserCreate) SetNillableOptBool(b *bool) *UserCreate {
	if b != nil {
		uc.SetOptBool(*b)
	}
	return uc
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (uc *UserCreate) SetGroupID(id int) *UserCreate {
	uc.mutation.SetGroupID(id)
	return uc
}

// SetNillableGroupID sets the "group" edge to the Group entity by ID if the given value is not nil.
func (uc *UserCreate) SetNillableGroupID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetGroupID(*id)
	}
	return uc
}

// SetGroup sets the "group" edge to the Group entity.
func (uc *UserCreate) SetGroup(g *Group) *UserCreate {
	return uc.SetGroupID(g.ID)
}

// SetAttachmentID sets the "attachment" edge to the Attachment entity by ID.
func (uc *UserCreate) SetAttachmentID(id uuid.UUID) *UserCreate {
	uc.mutation.SetAttachmentID(id)
	return uc
}

// SetNillableAttachmentID sets the "attachment" edge to the Attachment entity by ID if the given value is not nil.
func (uc *UserCreate) SetNillableAttachmentID(id *uuid.UUID) *UserCreate {
	if id != nil {
		uc = uc.SetAttachmentID(*id)
	}
	return uc
}

// SetAttachment sets the "attachment" edge to the Attachment entity.
func (uc *UserCreate) SetAttachment(a *Attachment) *UserCreate {
	return uc.SetAttachmentID(a.ID)
}

// AddReceivedIDs adds the "received" edge to the Attachment entity by IDs.
func (uc *UserCreate) AddReceivedIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddReceivedIDs(ids...)
	return uc
}

// AddReceived adds the "received" edges to the Attachment entity.
func (uc *UserCreate) AddReceived(a ...*Attachment) *UserCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uc.AddReceivedIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	uc.defaults()
	if len(uc.hooks) == 0 {
		if err = uc.check(); err != nil {
			return nil, err
		}
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uc.check(); err != nil {
				return nil, err
			}
			uc.mutation = mutation
			if node, err = uc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			mut = uc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.Banned(); !ok {
		v := user.DefaultBanned
		uc.mutation.SetBanned(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.UserName(); !ok {
		return &ValidationError{Name: "user_name", err: errors.New("ent: missing required field \"user_name\"")}
	}
	if _, ok := uc.mutation.Joined(); !ok {
		return &ValidationError{Name: "joined", err: errors.New("ent: missing required field \"joined\"")}
	}
	if _, ok := uc.mutation.Points(); !ok {
		return &ValidationError{Name: "points", err: errors.New("ent: missing required field \"points\"")}
	}
	if _, ok := uc.mutation.Exp(); !ok {
		return &ValidationError{Name: "exp", err: errors.New("ent: missing required field \"exp\"")}
	}
	if _, ok := uc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New("ent: missing required field \"status\"")}
	}
	if v, ok := uc.mutation.Status(); ok {
		if err := user.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	if _, ok := uc.mutation.ExternalID(); !ok {
		return &ValidationError{Name: "external_id", err: errors.New("ent: missing required field \"external_id\"")}
	}
	if _, ok := uc.mutation.CrmID(); !ok {
		return &ValidationError{Name: "crm_id", err: errors.New("ent: missing required field \"crm_id\"")}
	}
	if _, ok := uc.mutation.Banned(); !ok {
		return &ValidationError{Name: "banned", err: errors.New("ent: missing required field \"banned\"")}
	}
	if _, ok := uc.mutation.CustomPb(); !ok {
		return &ValidationError{Name: "custom_pb", err: errors.New("ent: missing required field \"custom_pb\"")}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		}
	)
	if value, ok := uc.mutation.UserName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUserName,
		})
		_node.UserName = value
	}
	if value, ok := uc.mutation.Joined(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldJoined,
		})
		_node.Joined = value
	}
	if value, ok := uc.mutation.Points(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: user.FieldPoints,
		})
		_node.Points = value
	}
	if value, ok := uc.mutation.Exp(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: user.FieldExp,
		})
		_node.Exp = value
	}
	if value, ok := uc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: user.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := uc.mutation.ExternalID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldExternalID,
		})
		_node.ExternalID = value
	}
	if value, ok := uc.mutation.CrmID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: user.FieldCrmID,
		})
		_node.CrmID = value
	}
	if value, ok := uc.mutation.Banned(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldBanned,
		})
		_node.Banned = value
	}
	if value, ok := uc.mutation.CustomPb(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: user.FieldCustomPb,
		})
		_node.CustomPb = value
	}
	if value, ok := uc.mutation.OptNum(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldOptNum,
		})
		_node.OptNum = value
	}
	if value, ok := uc.mutation.OptStr(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldOptStr,
		})
		_node.OptStr = value
	}
	if value, ok := uc.mutation.OptBool(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldOptBool,
		})
		_node.OptBool = value
	}
	if nodes := uc.mutation.GroupIDs(); len(nodes) > 0 {
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
		_node.user_group = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.AttachmentIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.ReceivedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.ReceivedTable,
			Columns: user.ReceivedPrimaryKey,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
