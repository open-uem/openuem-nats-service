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
	"github.com/doncicuto/openuem_ent/monitor"
	"github.com/doncicuto/openuem_ent/predicate"
)

// MonitorQuery is the builder for querying Monitor entities.
type MonitorQuery struct {
	config
	ctx        *QueryContext
	order      []monitor.OrderOption
	inters     []Interceptor
	predicates []predicate.Monitor
	withOwner  *AgentQuery
	withFKs    bool
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MonitorQuery builder.
func (mq *MonitorQuery) Where(ps ...predicate.Monitor) *MonitorQuery {
	mq.predicates = append(mq.predicates, ps...)
	return mq
}

// Limit the number of records to be returned by this query.
func (mq *MonitorQuery) Limit(limit int) *MonitorQuery {
	mq.ctx.Limit = &limit
	return mq
}

// Offset to start from.
func (mq *MonitorQuery) Offset(offset int) *MonitorQuery {
	mq.ctx.Offset = &offset
	return mq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mq *MonitorQuery) Unique(unique bool) *MonitorQuery {
	mq.ctx.Unique = &unique
	return mq
}

// Order specifies how the records should be ordered.
func (mq *MonitorQuery) Order(o ...monitor.OrderOption) *MonitorQuery {
	mq.order = append(mq.order, o...)
	return mq
}

// QueryOwner chains the current query on the "owner" edge.
func (mq *MonitorQuery) QueryOwner() *AgentQuery {
	query := (&AgentClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(monitor.Table, monitor.FieldID, selector),
			sqlgraph.To(agent.Table, agent.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, monitor.OwnerTable, monitor.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Monitor entity from the query.
// Returns a *NotFoundError when no Monitor was found.
func (mq *MonitorQuery) First(ctx context.Context) (*Monitor, error) {
	nodes, err := mq.Limit(1).All(setContextOp(ctx, mq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{monitor.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mq *MonitorQuery) FirstX(ctx context.Context) *Monitor {
	node, err := mq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Monitor ID from the query.
// Returns a *NotFoundError when no Monitor ID was found.
func (mq *MonitorQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mq.Limit(1).IDs(setContextOp(ctx, mq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{monitor.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mq *MonitorQuery) FirstIDX(ctx context.Context) int {
	id, err := mq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Monitor entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Monitor entity is found.
// Returns a *NotFoundError when no Monitor entities are found.
func (mq *MonitorQuery) Only(ctx context.Context) (*Monitor, error) {
	nodes, err := mq.Limit(2).All(setContextOp(ctx, mq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{monitor.Label}
	default:
		return nil, &NotSingularError{monitor.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mq *MonitorQuery) OnlyX(ctx context.Context) *Monitor {
	node, err := mq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Monitor ID in the query.
// Returns a *NotSingularError when more than one Monitor ID is found.
// Returns a *NotFoundError when no entities are found.
func (mq *MonitorQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mq.Limit(2).IDs(setContextOp(ctx, mq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{monitor.Label}
	default:
		err = &NotSingularError{monitor.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mq *MonitorQuery) OnlyIDX(ctx context.Context) int {
	id, err := mq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Monitors.
func (mq *MonitorQuery) All(ctx context.Context) ([]*Monitor, error) {
	ctx = setContextOp(ctx, mq.ctx, ent.OpQueryAll)
	if err := mq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Monitor, *MonitorQuery]()
	return withInterceptors[[]*Monitor](ctx, mq, qr, mq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mq *MonitorQuery) AllX(ctx context.Context) []*Monitor {
	nodes, err := mq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Monitor IDs.
func (mq *MonitorQuery) IDs(ctx context.Context) (ids []int, err error) {
	if mq.ctx.Unique == nil && mq.path != nil {
		mq.Unique(true)
	}
	ctx = setContextOp(ctx, mq.ctx, ent.OpQueryIDs)
	if err = mq.Select(monitor.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mq *MonitorQuery) IDsX(ctx context.Context) []int {
	ids, err := mq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mq *MonitorQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mq.ctx, ent.OpQueryCount)
	if err := mq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mq, querierCount[*MonitorQuery](), mq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mq *MonitorQuery) CountX(ctx context.Context) int {
	count, err := mq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mq *MonitorQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mq.ctx, ent.OpQueryExist)
	switch _, err := mq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("openuem_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mq *MonitorQuery) ExistX(ctx context.Context) bool {
	exist, err := mq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MonitorQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mq *MonitorQuery) Clone() *MonitorQuery {
	if mq == nil {
		return nil
	}
	return &MonitorQuery{
		config:     mq.config,
		ctx:        mq.ctx.Clone(),
		order:      append([]monitor.OrderOption{}, mq.order...),
		inters:     append([]Interceptor{}, mq.inters...),
		predicates: append([]predicate.Monitor{}, mq.predicates...),
		withOwner:  mq.withOwner.Clone(),
		// clone intermediate query.
		sql:       mq.sql.Clone(),
		path:      mq.path,
		modifiers: append([]func(*sql.Selector){}, mq.modifiers...),
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MonitorQuery) WithOwner(opts ...func(*AgentQuery)) *MonitorQuery {
	query := (&AgentClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withOwner = query
	return mq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Manufacturer string `json:"manufacturer,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Monitor.Query().
//		GroupBy(monitor.FieldManufacturer).
//		Aggregate(openuem_ent.Count()).
//		Scan(ctx, &v)
func (mq *MonitorQuery) GroupBy(field string, fields ...string) *MonitorGroupBy {
	mq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MonitorGroupBy{build: mq}
	grbuild.flds = &mq.ctx.Fields
	grbuild.label = monitor.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Manufacturer string `json:"manufacturer,omitempty"`
//	}
//
//	client.Monitor.Query().
//		Select(monitor.FieldManufacturer).
//		Scan(ctx, &v)
func (mq *MonitorQuery) Select(fields ...string) *MonitorSelect {
	mq.ctx.Fields = append(mq.ctx.Fields, fields...)
	sbuild := &MonitorSelect{MonitorQuery: mq}
	sbuild.label = monitor.Label
	sbuild.flds, sbuild.scan = &mq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MonitorSelect configured with the given aggregations.
func (mq *MonitorQuery) Aggregate(fns ...AggregateFunc) *MonitorSelect {
	return mq.Select().Aggregate(fns...)
}

func (mq *MonitorQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mq.inters {
		if inter == nil {
			return fmt.Errorf("openuem_ent: uninitialized interceptor (forgotten import openuem_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mq); err != nil {
				return err
			}
		}
	}
	for _, f := range mq.ctx.Fields {
		if !monitor.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("openuem_ent: invalid field %q for query", f)}
		}
	}
	if mq.path != nil {
		prev, err := mq.path(ctx)
		if err != nil {
			return err
		}
		mq.sql = prev
	}
	return nil
}

func (mq *MonitorQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Monitor, error) {
	var (
		nodes       = []*Monitor{}
		withFKs     = mq.withFKs
		_spec       = mq.querySpec()
		loadedTypes = [1]bool{
			mq.withOwner != nil,
		}
	)
	if mq.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, monitor.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Monitor).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Monitor{config: mq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(mq.modifiers) > 0 {
		_spec.Modifiers = mq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mq.withOwner; query != nil {
		if err := mq.loadOwner(ctx, query, nodes, nil,
			func(n *Monitor, e *Agent) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mq *MonitorQuery) loadOwner(ctx context.Context, query *AgentQuery, nodes []*Monitor, init func(*Monitor), assign func(*Monitor, *Agent)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Monitor)
	for i := range nodes {
		if nodes[i].agent_monitors == nil {
			continue
		}
		fk := *nodes[i].agent_monitors
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
			return fmt.Errorf(`unexpected foreign-key "agent_monitors" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (mq *MonitorQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mq.querySpec()
	if len(mq.modifiers) > 0 {
		_spec.Modifiers = mq.modifiers
	}
	_spec.Node.Columns = mq.ctx.Fields
	if len(mq.ctx.Fields) > 0 {
		_spec.Unique = mq.ctx.Unique != nil && *mq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mq.driver, _spec)
}

func (mq *MonitorQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(monitor.Table, monitor.Columns, sqlgraph.NewFieldSpec(monitor.FieldID, field.TypeInt))
	_spec.From = mq.sql
	if unique := mq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mq.path != nil {
		_spec.Unique = true
	}
	if fields := mq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, monitor.FieldID)
		for i := range fields {
			if fields[i] != monitor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mq *MonitorQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mq.driver.Dialect())
	t1 := builder.Table(monitor.Table)
	columns := mq.ctx.Fields
	if len(columns) == 0 {
		columns = monitor.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mq.sql != nil {
		selector = mq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mq.ctx.Unique != nil && *mq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range mq.modifiers {
		m(selector)
	}
	for _, p := range mq.predicates {
		p(selector)
	}
	for _, p := range mq.order {
		p(selector)
	}
	if offset := mq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (mq *MonitorQuery) Modify(modifiers ...func(s *sql.Selector)) *MonitorSelect {
	mq.modifiers = append(mq.modifiers, modifiers...)
	return mq.Select()
}

// MonitorGroupBy is the group-by builder for Monitor entities.
type MonitorGroupBy struct {
	selector
	build *MonitorQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mgb *MonitorGroupBy) Aggregate(fns ...AggregateFunc) *MonitorGroupBy {
	mgb.fns = append(mgb.fns, fns...)
	return mgb
}

// Scan applies the selector query and scans the result into the given value.
func (mgb *MonitorGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mgb.build.ctx, ent.OpQueryGroupBy)
	if err := mgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MonitorQuery, *MonitorGroupBy](ctx, mgb.build, mgb, mgb.build.inters, v)
}

func (mgb *MonitorGroupBy) sqlScan(ctx context.Context, root *MonitorQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mgb.fns))
	for _, fn := range mgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mgb.flds)+len(mgb.fns))
		for _, f := range *mgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MonitorSelect is the builder for selecting fields of Monitor entities.
type MonitorSelect struct {
	*MonitorQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ms *MonitorSelect) Aggregate(fns ...AggregateFunc) *MonitorSelect {
	ms.fns = append(ms.fns, fns...)
	return ms
}

// Scan applies the selector query and scans the result into the given value.
func (ms *MonitorSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ms.ctx, ent.OpQuerySelect)
	if err := ms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MonitorQuery, *MonitorSelect](ctx, ms.MonitorQuery, ms, ms.inters, v)
}

func (ms *MonitorSelect) sqlScan(ctx context.Context, root *MonitorQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ms.fns))
	for _, fn := range ms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ms *MonitorSelect) Modify(modifiers ...func(s *sql.Selector)) *MonitorSelect {
	ms.modifiers = append(ms.modifiers, modifiers...)
	return ms
}
