// Code generated by ent, DO NOT EDIT.

package openuem_ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/doncicuto/openuem_ent/agent"
	"github.com/doncicuto/openuem_ent/predicate"
	"github.com/doncicuto/openuem_ent/share"
)

// ShareQuery is the builder for querying Share entities.
type ShareQuery struct {
	config
	ctx        *QueryContext
	order      []share.OrderOption
	inters     []Interceptor
	predicates []predicate.Share
	withOwner  *AgentQuery
	withFKs    bool
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ShareQuery builder.
func (sq *ShareQuery) Where(ps ...predicate.Share) *ShareQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit the number of records to be returned by this query.
func (sq *ShareQuery) Limit(limit int) *ShareQuery {
	sq.ctx.Limit = &limit
	return sq
}

// Offset to start from.
func (sq *ShareQuery) Offset(offset int) *ShareQuery {
	sq.ctx.Offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *ShareQuery) Unique(unique bool) *ShareQuery {
	sq.ctx.Unique = &unique
	return sq
}

// Order specifies how the records should be ordered.
func (sq *ShareQuery) Order(o ...share.OrderOption) *ShareQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryOwner chains the current query on the "owner" edge.
func (sq *ShareQuery) QueryOwner() *AgentQuery {
	query := (&AgentClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(share.Table, share.FieldID, selector),
			sqlgraph.To(agent.Table, agent.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, share.OwnerTable, share.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Share entity from the query.
// Returns a *NotFoundError when no Share was found.
func (sq *ShareQuery) First(ctx context.Context) (*Share, error) {
	nodes, err := sq.Limit(1).All(setContextOp(ctx, sq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{share.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *ShareQuery) FirstX(ctx context.Context) *Share {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Share ID from the query.
// Returns a *NotFoundError when no Share ID was found.
func (sq *ShareQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(1).IDs(setContextOp(ctx, sq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{share.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *ShareQuery) FirstIDX(ctx context.Context) int {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Share entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Share entity is found.
// Returns a *NotFoundError when no Share entities are found.
func (sq *ShareQuery) Only(ctx context.Context) (*Share, error) {
	nodes, err := sq.Limit(2).All(setContextOp(ctx, sq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{share.Label}
	default:
		return nil, &NotSingularError{share.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *ShareQuery) OnlyX(ctx context.Context) *Share {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Share ID in the query.
// Returns a *NotSingularError when more than one Share ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *ShareQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(2).IDs(setContextOp(ctx, sq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{share.Label}
	default:
		err = &NotSingularError{share.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *ShareQuery) OnlyIDX(ctx context.Context) int {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Shares.
func (sq *ShareQuery) All(ctx context.Context) ([]*Share, error) {
	ctx = setContextOp(ctx, sq.ctx, ent.OpQueryAll)
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Share, *ShareQuery]()
	return withInterceptors[[]*Share](ctx, sq, qr, sq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sq *ShareQuery) AllX(ctx context.Context) []*Share {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Share IDs.
func (sq *ShareQuery) IDs(ctx context.Context) (ids []int, err error) {
	if sq.ctx.Unique == nil && sq.path != nil {
		sq.Unique(true)
	}
	ctx = setContextOp(ctx, sq.ctx, ent.OpQueryIDs)
	if err = sq.Select(share.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *ShareQuery) IDsX(ctx context.Context) []int {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *ShareQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sq.ctx, ent.OpQueryCount)
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sq, querierCount[*ShareQuery](), sq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sq *ShareQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *ShareQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sq.ctx, ent.OpQueryExist)
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("openuem_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *ShareQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ShareQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *ShareQuery) Clone() *ShareQuery {
	if sq == nil {
		return nil
	}
	return &ShareQuery{
		config:     sq.config,
		ctx:        sq.ctx.Clone(),
		order:      append([]share.OrderOption{}, sq.order...),
		inters:     append([]Interceptor{}, sq.inters...),
		predicates: append([]predicate.Share{}, sq.predicates...),
		withOwner:  sq.withOwner.Clone(),
		// clone intermediate query.
		sql:       sq.sql.Clone(),
		path:      sq.path,
		modifiers: append([]func(*sql.Selector){}, sq.modifiers...),
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *ShareQuery) WithOwner(opts ...func(*AgentQuery)) *ShareQuery {
	query := (&AgentClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withOwner = query
	return sq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Share.Query().
//		GroupBy(share.FieldName).
//		Aggregate(openuem_ent.Count()).
//		Scan(ctx, &v)
func (sq *ShareQuery) GroupBy(field string, fields ...string) *ShareGroupBy {
	sq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ShareGroupBy{build: sq}
	grbuild.flds = &sq.ctx.Fields
	grbuild.label = share.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Share.Query().
//		Select(share.FieldName).
//		Scan(ctx, &v)
func (sq *ShareQuery) Select(fields ...string) *ShareSelect {
	sq.ctx.Fields = append(sq.ctx.Fields, fields...)
	sbuild := &ShareSelect{ShareQuery: sq}
	sbuild.label = share.Label
	sbuild.flds, sbuild.scan = &sq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ShareSelect configured with the given aggregations.
func (sq *ShareQuery) Aggregate(fns ...AggregateFunc) *ShareSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *ShareQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sq.inters {
		if inter == nil {
			return fmt.Errorf("openuem_ent: uninitialized interceptor (forgotten import openuem_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sq); err != nil {
				return err
			}
		}
	}
	for _, f := range sq.ctx.Fields {
		if !share.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("openuem_ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *ShareQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Share, error) {
	var (
		nodes       = []*Share{}
		withFKs     = sq.withFKs
		_spec       = sq.querySpec()
		loadedTypes = [1]bool{
			sq.withOwner != nil,
		}
	)
	if sq.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, share.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Share).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Share{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(sq.modifiers) > 0 {
		_spec.Modifiers = sq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sq.withOwner; query != nil {
		if err := sq.loadOwner(ctx, query, nodes, nil,
			func(n *Share, e *Agent) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sq *ShareQuery) loadOwner(ctx context.Context, query *AgentQuery, nodes []*Share, init func(*Share), assign func(*Share, *Agent)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Share)
	for i := range nodes {
		if nodes[i].agent_shares == nil {
			continue
		}
		fk := *nodes[i].agent_shares
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(agent.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "agent_shares" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (sq *ShareQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	if len(sq.modifiers) > 0 {
		_spec.Modifiers = sq.modifiers
	}
	_spec.Node.Columns = sq.ctx.Fields
	if len(sq.ctx.Fields) > 0 {
		_spec.Unique = sq.ctx.Unique != nil && *sq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *ShareQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(share.Table, share.Columns, sqlgraph.NewFieldSpec(share.FieldID, field.TypeInt))
	_spec.From = sq.sql
	if unique := sq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sq.path != nil {
		_spec.Unique = true
	}
	if fields := sq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, share.FieldID)
		for i := range fields {
			if fields[i] != share.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *ShareQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(share.Table)
	columns := sq.ctx.Fields
	if len(columns) == 0 {
		columns = share.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.ctx.Unique != nil && *sq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range sq.modifiers {
		m(selector)
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (sq *ShareQuery) Modify(modifiers ...func(s *sql.Selector)) *ShareSelect {
	sq.modifiers = append(sq.modifiers, modifiers...)
	return sq.Select()
}

// ShareGroupBy is the group-by builder for Share entities.
type ShareGroupBy struct {
	selector
	build *ShareQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *ShareGroupBy) Aggregate(fns ...AggregateFunc) *ShareGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the selector query and scans the result into the given value.
func (sgb *ShareGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sgb.build.ctx, ent.OpQueryGroupBy)
	if err := sgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ShareQuery, *ShareGroupBy](ctx, sgb.build, sgb, sgb.build.inters, v)
}

func (sgb *ShareGroupBy) sqlScan(ctx context.Context, root *ShareQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sgb.flds)+len(sgb.fns))
		for _, f := range *sgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ShareSelect is the builder for selecting fields of Share entities.
type ShareSelect struct {
	*ShareQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *ShareSelect) Aggregate(fns ...AggregateFunc) *ShareSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *ShareSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ss.ctx, ent.OpQuerySelect)
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ShareQuery, *ShareSelect](ctx, ss.ShareQuery, ss, ss.inters, v)
}

func (ss *ShareSelect) sqlScan(ctx context.Context, root *ShareQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ss.fns))
	for _, fn := range ss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ss *ShareSelect) Modify(modifiers ...func(s *sql.Selector)) *ShareSelect {
	ss.modifiers = append(ss.modifiers, modifiers...)
	return ss
}
