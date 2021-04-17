// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/contrib/entproto/internal/entprototest/ent/messagewithoptionals"
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageWithOptionalsQuery is the builder for querying MessageWithOptionals entities.
type MessageWithOptionalsQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.MessageWithOptionals
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MessageWithOptionalsQuery builder.
func (mwoq *MessageWithOptionalsQuery) Where(ps ...predicate.MessageWithOptionals) *MessageWithOptionalsQuery {
	mwoq.predicates = append(mwoq.predicates, ps...)
	return mwoq
}

// Limit adds a limit step to the query.
func (mwoq *MessageWithOptionalsQuery) Limit(limit int) *MessageWithOptionalsQuery {
	mwoq.limit = &limit
	return mwoq
}

// Offset adds an offset step to the query.
func (mwoq *MessageWithOptionalsQuery) Offset(offset int) *MessageWithOptionalsQuery {
	mwoq.offset = &offset
	return mwoq
}

// Order adds an order step to the query.
func (mwoq *MessageWithOptionalsQuery) Order(o ...OrderFunc) *MessageWithOptionalsQuery {
	mwoq.order = append(mwoq.order, o...)
	return mwoq
}

// First returns the first MessageWithOptionals entity from the query.
// Returns a *NotFoundError when no MessageWithOptionals was found.
func (mwoq *MessageWithOptionalsQuery) First(ctx context.Context) (*MessageWithOptionals, error) {
	nodes, err := mwoq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{messagewithoptionals.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mwoq *MessageWithOptionalsQuery) FirstX(ctx context.Context) *MessageWithOptionals {
	node, err := mwoq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MessageWithOptionals ID from the query.
// Returns a *NotFoundError when no MessageWithOptionals ID was found.
func (mwoq *MessageWithOptionalsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mwoq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{messagewithoptionals.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mwoq *MessageWithOptionalsQuery) FirstIDX(ctx context.Context) int {
	id, err := mwoq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MessageWithOptionals entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one MessageWithOptionals entity is not found.
// Returns a *NotFoundError when no MessageWithOptionals entities are found.
func (mwoq *MessageWithOptionalsQuery) Only(ctx context.Context) (*MessageWithOptionals, error) {
	nodes, err := mwoq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{messagewithoptionals.Label}
	default:
		return nil, &NotSingularError{messagewithoptionals.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mwoq *MessageWithOptionalsQuery) OnlyX(ctx context.Context) *MessageWithOptionals {
	node, err := mwoq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MessageWithOptionals ID in the query.
// Returns a *NotSingularError when exactly one MessageWithOptionals ID is not found.
// Returns a *NotFoundError when no entities are found.
func (mwoq *MessageWithOptionalsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mwoq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{messagewithoptionals.Label}
	default:
		err = &NotSingularError{messagewithoptionals.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mwoq *MessageWithOptionalsQuery) OnlyIDX(ctx context.Context) int {
	id, err := mwoq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MessageWithOptionalsSlice.
func (mwoq *MessageWithOptionalsQuery) All(ctx context.Context) ([]*MessageWithOptionals, error) {
	if err := mwoq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mwoq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mwoq *MessageWithOptionalsQuery) AllX(ctx context.Context) []*MessageWithOptionals {
	nodes, err := mwoq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MessageWithOptionals IDs.
func (mwoq *MessageWithOptionalsQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := mwoq.Select(messagewithoptionals.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mwoq *MessageWithOptionalsQuery) IDsX(ctx context.Context) []int {
	ids, err := mwoq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mwoq *MessageWithOptionalsQuery) Count(ctx context.Context) (int, error) {
	if err := mwoq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mwoq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mwoq *MessageWithOptionalsQuery) CountX(ctx context.Context) int {
	count, err := mwoq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mwoq *MessageWithOptionalsQuery) Exist(ctx context.Context) (bool, error) {
	if err := mwoq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mwoq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mwoq *MessageWithOptionalsQuery) ExistX(ctx context.Context) bool {
	exist, err := mwoq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MessageWithOptionalsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mwoq *MessageWithOptionalsQuery) Clone() *MessageWithOptionalsQuery {
	if mwoq == nil {
		return nil
	}
	return &MessageWithOptionalsQuery{
		config:     mwoq.config,
		limit:      mwoq.limit,
		offset:     mwoq.offset,
		order:      append([]OrderFunc{}, mwoq.order...),
		predicates: append([]predicate.MessageWithOptionals{}, mwoq.predicates...),
		// clone intermediate query.
		sql:  mwoq.sql.Clone(),
		path: mwoq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		StrOptional string `json:"str_optional,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MessageWithOptionals.Query().
//		GroupBy(messagewithoptionals.FieldStrOptional).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (mwoq *MessageWithOptionalsQuery) GroupBy(field string, fields ...string) *MessageWithOptionalsGroupBy {
	group := &MessageWithOptionalsGroupBy{config: mwoq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mwoq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mwoq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		StrOptional string `json:"str_optional,omitempty"`
//	}
//
//	client.MessageWithOptionals.Query().
//		Select(messagewithoptionals.FieldStrOptional).
//		Scan(ctx, &v)
//
func (mwoq *MessageWithOptionalsQuery) Select(field string, fields ...string) *MessageWithOptionalsSelect {
	mwoq.fields = append([]string{field}, fields...)
	return &MessageWithOptionalsSelect{MessageWithOptionalsQuery: mwoq}
}

func (mwoq *MessageWithOptionalsQuery) prepareQuery(ctx context.Context) error {
	for _, f := range mwoq.fields {
		if !messagewithoptionals.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mwoq.path != nil {
		prev, err := mwoq.path(ctx)
		if err != nil {
			return err
		}
		mwoq.sql = prev
	}
	return nil
}

func (mwoq *MessageWithOptionalsQuery) sqlAll(ctx context.Context) ([]*MessageWithOptionals, error) {
	var (
		nodes = []*MessageWithOptionals{}
		_spec = mwoq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &MessageWithOptionals{config: mwoq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, mwoq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (mwoq *MessageWithOptionalsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mwoq.querySpec()
	return sqlgraph.CountNodes(ctx, mwoq.driver, _spec)
}

func (mwoq *MessageWithOptionalsQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := mwoq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (mwoq *MessageWithOptionalsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   messagewithoptionals.Table,
			Columns: messagewithoptionals.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: messagewithoptionals.FieldID,
			},
		},
		From:   mwoq.sql,
		Unique: true,
	}
	if fields := mwoq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, messagewithoptionals.FieldID)
		for i := range fields {
			if fields[i] != messagewithoptionals.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mwoq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mwoq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mwoq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mwoq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, messagewithoptionals.ValidColumn)
			}
		}
	}
	return _spec
}

func (mwoq *MessageWithOptionalsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mwoq.driver.Dialect())
	t1 := builder.Table(messagewithoptionals.Table)
	selector := builder.Select(t1.Columns(messagewithoptionals.Columns...)...).From(t1)
	if mwoq.sql != nil {
		selector = mwoq.sql
		selector.Select(selector.Columns(messagewithoptionals.Columns...)...)
	}
	for _, p := range mwoq.predicates {
		p(selector)
	}
	for _, p := range mwoq.order {
		p(selector, messagewithoptionals.ValidColumn)
	}
	if offset := mwoq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mwoq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MessageWithOptionalsGroupBy is the group-by builder for MessageWithOptionals entities.
type MessageWithOptionalsGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mwogb *MessageWithOptionalsGroupBy) Aggregate(fns ...AggregateFunc) *MessageWithOptionalsGroupBy {
	mwogb.fns = append(mwogb.fns, fns...)
	return mwogb
}

// Scan applies the group-by query and scans the result into the given value.
func (mwogb *MessageWithOptionalsGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := mwogb.path(ctx)
	if err != nil {
		return err
	}
	mwogb.sql = query
	return mwogb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mwogb *MessageWithOptionalsGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := mwogb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (mwogb *MessageWithOptionalsGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(mwogb.fields) > 1 {
		return nil, errors.New("ent: MessageWithOptionalsGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := mwogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mwogb *MessageWithOptionalsGroupBy) StringsX(ctx context.Context) []string {
	v, err := mwogb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mwogb *MessageWithOptionalsGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mwogb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{messagewithoptionals.Label}
	default:
		err = fmt.Errorf("ent: MessageWithOptionalsGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mwogb *MessageWithOptionalsGroupBy) StringX(ctx context.Context) string {
	v, err := mwogb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (mwogb *MessageWithOptionalsGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(mwogb.fields) > 1 {
		return nil, errors.New("ent: MessageWithOptionalsGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := mwogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mwogb *MessageWithOptionalsGroupBy) IntsX(ctx context.Context) []int {
	v, err := mwogb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mwogb *MessageWithOptionalsGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mwogb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{messagewithoptionals.Label}
	default:
		err = fmt.Errorf("ent: MessageWithOptionalsGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mwogb *MessageWithOptionalsGroupBy) IntX(ctx context.Context) int {
	v, err := mwogb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (mwogb *MessageWithOptionalsGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(mwogb.fields) > 1 {
		return nil, errors.New("ent: MessageWithOptionalsGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := mwogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mwogb *MessageWithOptionalsGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := mwogb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mwogb *MessageWithOptionalsGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mwogb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{messagewithoptionals.Label}
	default:
		err = fmt.Errorf("ent: MessageWithOptionalsGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mwogb *MessageWithOptionalsGroupBy) Float64X(ctx context.Context) float64 {
	v, err := mwogb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (mwogb *MessageWithOptionalsGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(mwogb.fields) > 1 {
		return nil, errors.New("ent: MessageWithOptionalsGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := mwogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mwogb *MessageWithOptionalsGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := mwogb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mwogb *MessageWithOptionalsGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mwogb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{messagewithoptionals.Label}
	default:
		err = fmt.Errorf("ent: MessageWithOptionalsGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mwogb *MessageWithOptionalsGroupBy) BoolX(ctx context.Context) bool {
	v, err := mwogb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mwogb *MessageWithOptionalsGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range mwogb.fields {
		if !messagewithoptionals.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := mwogb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mwogb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mwogb *MessageWithOptionalsGroupBy) sqlQuery() *sql.Selector {
	selector := mwogb.sql
	columns := make([]string, 0, len(mwogb.fields)+len(mwogb.fns))
	columns = append(columns, mwogb.fields...)
	for _, fn := range mwogb.fns {
		columns = append(columns, fn(selector, messagewithoptionals.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(mwogb.fields...)
}

// MessageWithOptionalsSelect is the builder for selecting fields of MessageWithOptionals entities.
type MessageWithOptionalsSelect struct {
	*MessageWithOptionalsQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (mwos *MessageWithOptionalsSelect) Scan(ctx context.Context, v interface{}) error {
	if err := mwos.prepareQuery(ctx); err != nil {
		return err
	}
	mwos.sql = mwos.MessageWithOptionalsQuery.sqlQuery(ctx)
	return mwos.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mwos *MessageWithOptionalsSelect) ScanX(ctx context.Context, v interface{}) {
	if err := mwos.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (mwos *MessageWithOptionalsSelect) Strings(ctx context.Context) ([]string, error) {
	if len(mwos.fields) > 1 {
		return nil, errors.New("ent: MessageWithOptionalsSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := mwos.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mwos *MessageWithOptionalsSelect) StringsX(ctx context.Context) []string {
	v, err := mwos.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (mwos *MessageWithOptionalsSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mwos.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{messagewithoptionals.Label}
	default:
		err = fmt.Errorf("ent: MessageWithOptionalsSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mwos *MessageWithOptionalsSelect) StringX(ctx context.Context) string {
	v, err := mwos.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (mwos *MessageWithOptionalsSelect) Ints(ctx context.Context) ([]int, error) {
	if len(mwos.fields) > 1 {
		return nil, errors.New("ent: MessageWithOptionalsSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := mwos.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mwos *MessageWithOptionalsSelect) IntsX(ctx context.Context) []int {
	v, err := mwos.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (mwos *MessageWithOptionalsSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mwos.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{messagewithoptionals.Label}
	default:
		err = fmt.Errorf("ent: MessageWithOptionalsSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mwos *MessageWithOptionalsSelect) IntX(ctx context.Context) int {
	v, err := mwos.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (mwos *MessageWithOptionalsSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(mwos.fields) > 1 {
		return nil, errors.New("ent: MessageWithOptionalsSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := mwos.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mwos *MessageWithOptionalsSelect) Float64sX(ctx context.Context) []float64 {
	v, err := mwos.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (mwos *MessageWithOptionalsSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mwos.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{messagewithoptionals.Label}
	default:
		err = fmt.Errorf("ent: MessageWithOptionalsSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mwos *MessageWithOptionalsSelect) Float64X(ctx context.Context) float64 {
	v, err := mwos.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (mwos *MessageWithOptionalsSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(mwos.fields) > 1 {
		return nil, errors.New("ent: MessageWithOptionalsSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := mwos.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mwos *MessageWithOptionalsSelect) BoolsX(ctx context.Context) []bool {
	v, err := mwos.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (mwos *MessageWithOptionalsSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mwos.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{messagewithoptionals.Label}
	default:
		err = fmt.Errorf("ent: MessageWithOptionalsSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mwos *MessageWithOptionalsSelect) BoolX(ctx context.Context) bool {
	v, err := mwos.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mwos *MessageWithOptionalsSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := mwos.sqlQuery().Query()
	if err := mwos.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mwos *MessageWithOptionalsSelect) sqlQuery() sql.Querier {
	selector := mwos.sql
	selector.Select(selector.Columns(mwos.fields...)...)
	return selector
}
