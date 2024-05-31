package pgOrm

import (
	"context"
	"database/sql"

	"github.com/naufalfmm/dayatani-farmer-api/utils/orm"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

type postgresOrm struct {
	db *bun.DB
	tx *bun.Tx
}

func (p *postgresOrm) NewSelect() orm.Select {
	sq := p.db.NewSelect()
	if p.tx != nil {
		sq = p.tx.NewSelect()
	}

	return &selectPostgreOrm{
		sq: sq,
	}
}

func (p *postgresOrm) NewValues(model interface{}) orm.Values {
	vq := p.db.NewValues(model)
	if p.tx != nil {
		vq = p.tx.NewValues(model)
	}

	return &valuesPostgreOrm{
		vq: vq,
	}
}

func (p *postgresOrm) NewInsert() orm.Insert {
	iq := p.db.NewInsert()
	if p.tx != nil {
		iq = p.tx.NewInsert()
	}

	return &insertPostgreOrm{
		iq: iq,
	}
}

func (p *postgresOrm) NewUpdate() orm.Update {
	uq := p.db.NewUpdate()
	if p.tx != nil {
		uq = p.tx.NewUpdate()
	}

	return &updatePostgreOrm{
		uq: uq,
	}
}

func (p *postgresOrm) NewDelete() orm.Delete {
	dq := p.db.NewDelete()
	if p.tx != nil {
		dq = p.tx.NewDelete()
	}

	return &deletePostgreOrm{
		dq: dq,
	}
}

func (p *postgresOrm) NewRaw(query string, args ...interface{}) orm.Raw {
	rq := p.db.NewRaw(query, args...)
	if p.tx != nil {
		rq = p.tx.NewRaw(query, args...)
	}

	return &rawPostgreOrm{
		rq: rq,
	}
}

func (p *postgresOrm) NewCreateTable() orm.CreateTable {
	ctq := p.db.NewCreateTable()
	if p.tx != nil {
		ctq = p.tx.NewCreateTable()
	}

	return &createTablePostgreOrm{
		ctq: ctq,
	}
}

func (p *postgresOrm) Begin() (orm.Orm, error) {
	if p.tx != nil {
		return p, nil
	}

	tx, err := p.db.Begin()
	if err != nil {
		return p, err
	}

	p.tx = &tx

	return p, nil
}

func (p *postgresOrm) BeginTx(ctx context.Context, opts *sql.TxOptions) (orm.Orm, error) {
	if p.tx != nil {
		return p, nil
	}

	tx, err := p.db.BeginTx(ctx, opts)
	if err != nil {
		return p, err
	}

	p.tx = &tx

	return p, nil
}

func (p *postgresOrm) Commit() error {
	if p.tx == nil {
		return nil
	}

	if err := p.tx.Commit(); err != nil {
		return err
	}

	p.tx = nil

	return nil
}
func (p *postgresOrm) Rollback() error {
	if p.tx == nil {
		return nil
	}

	if err := p.tx.Rollback(); err != nil {
		return err
	}

	p.tx = nil

	return nil
}

func (p *postgresOrm) RunInTx(ctx context.Context, opts *sql.TxOptions, fn func(ctx context.Context, tx orm.Orm) error) error {
	return p.db.RunInTx(ctx, opts, func(ctx context.Context, tx bun.Tx) error {
		return fn(ctx, &postgresOrm{
			db: p.db,
			tx: &tx,
		})
	})
}

func (p *postgresOrm) Exec(query string, args ...interface{}) (sql.Result, error) {
	return (*p.db).Exec(query, args...)
}

func (p *postgresOrm) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return (*p.db).ExecContext(ctx, query, args...)
}

func (p *postgresOrm) Ping() error {
	return p.db.Ping()
}

func (p *postgresOrm) PingContext(ctx context.Context) error {
	return p.db.PingContext(ctx)
}

type selectPostgreOrm struct {
	sq *bun.SelectQuery
}

func (so *selectPostgreOrm) AppendQuery(fmter schema.Formatter, b []byte) ([]byte, error) {
	return so.sq.AppendQuery(fmter, b)
}

func (so *selectPostgreOrm) With(name string, query orm.QueryAppender) orm.Select {
	sqa, ok := query.(*selectPostgreOrm)
	if ok {
		so.sq = so.sq.With(name, sqa)
	}

	vqa, ok := query.(*valuesPostgreOrm)
	if ok {
		so.sq = so.sq.With(name, vqa)
	}

	return so
}

func (so *selectPostgreOrm) Union(other orm.Select) orm.Select {
	so.sq = so.sq.Union(other.(*selectPostgreOrm).sq)

	return so
}

func (so *selectPostgreOrm) UnionAll(other orm.Select) orm.Select {
	so.sq = so.sq.UnionAll(other.(*selectPostgreOrm).sq)

	return so
}

func (so *selectPostgreOrm) Model(model interface{}) orm.Select {
	so.sq = so.sq.Model(model)

	return so
}

func (so *selectPostgreOrm) Column(columns ...string) orm.Select {
	so.sq = so.sq.Column(columns...)

	return so
}

func (so *selectPostgreOrm) ColumnExpr(query string, args ...interface{}) orm.Select {
	so.sq = so.sq.ColumnExpr(query, args...)

	return so
}

func (so *selectPostgreOrm) ExcludeColumn(columns ...string) orm.Select {
	so.sq = so.sq.ExcludeColumn(columns...)

	return so
}

func (so *selectPostgreOrm) Table(tables ...string) orm.Select {
	so.sq = so.sq.Table(tables...)

	return so
}

func (so *selectPostgreOrm) TableExpr(query string, args ...interface{}) orm.Select {
	so.sq = so.sq.TableExpr(query, args...)

	return so
}

func (so *selectPostgreOrm) ModelTableExpr(query string, args ...interface{}) orm.Select {
	so.sq = so.sq.ModelTableExpr(query, args...)

	return so
}

func (so *selectPostgreOrm) Join(join string, args ...interface{}) orm.Select {
	so.sq = so.sq.Join(join, args...)

	return so
}

func (so *selectPostgreOrm) JoinOn(cond string, args ...interface{}) orm.Select {
	so.sq = so.sq.JoinOn(cond, args...)

	return so
}

func (so *selectPostgreOrm) JoinOnOr(cond string, args ...interface{}) orm.Select {
	so.sq = so.sq.JoinOnOr(cond, args...)

	return so
}

func (so *selectPostgreOrm) Where(query string, args ...interface{}) orm.Select {
	so.sq = so.sq.Where(query, args...)

	return so
}

func (so *selectPostgreOrm) WhereAllWithDeleted() orm.Select {
	so.sq = so.sq.WhereAllWithDeleted()

	return so
}

func (so *selectPostgreOrm) WhereDeleted() orm.Select {
	so.sq = so.sq.WhereDeleted()

	return so
}

func (so *selectPostgreOrm) WhereGroup(sep string, fn func(orm.Select) orm.Select) orm.Select {
	so.sq = so.sq.WhereGroup(sep, func(sq *bun.SelectQuery) *bun.SelectQuery {
		return fn(so).(*selectPostgreOrm).sq
	})

	return so
}

func (so *selectPostgreOrm) WhereOr(query string, args ...interface{}) orm.Select {
	so.sq = so.sq.WhereOr(query, args...)

	return so
}

func (so *selectPostgreOrm) WherePK(cols ...string) orm.Select {
	so.sq = so.sq.WherePK(cols...)

	return so
}

func (so *selectPostgreOrm) Group(columns ...string) orm.Select {
	so.sq = so.sq.Group(columns...)

	return so
}

func (so *selectPostgreOrm) GroupExpr(group string, args ...interface{}) orm.Select {
	so.sq = so.sq.GroupExpr(group, args...)

	return so
}

func (so *selectPostgreOrm) Order(orders ...string) orm.Select {
	so.sq = so.sq.Order(orders...)

	return so
}

func (so *selectPostgreOrm) OrderExpr(query string, args ...interface{}) orm.Select {
	so.sq = so.sq.OrderExpr(query, args...)

	return so
}

func (so *selectPostgreOrm) Having(having string, args ...interface{}) orm.Select {
	so.sq = so.sq.Having(having, args...)

	return so
}

func (so *selectPostgreOrm) Limit(n int) orm.Select {
	so.sq = so.sq.Limit(n)

	return so
}

func (so *selectPostgreOrm) Offset(n int) orm.Select {
	so.sq = so.sq.Offset(n)

	return so
}

func (so *selectPostgreOrm) Relation(name string, apply ...func(orm.Select) orm.Select) orm.Select {
	sqapply := []func(sq *bun.SelectQuery) *bun.SelectQuery{}
	for _, ap := range apply {
		sqapply = append(sqapply, func(sq *bun.SelectQuery) *bun.SelectQuery {
			return ap(so).(*selectPostgreOrm).sq
		})
	}

	so.sq = so.sq.Relation(name, sqapply...)

	return so
}

func (so *selectPostgreOrm) UseIndex(indexes ...string) orm.Select {
	so.sq = so.sq.UseIndex(indexes...)

	return so
}

func (so *selectPostgreOrm) UseIndexForGroupBy(indexes ...string) orm.Select {
	so.sq = so.sq.UseIndexForGroupBy(indexes...)

	return so
}

func (so *selectPostgreOrm) UseIndexForJoin(indexes ...string) orm.Select {
	so.sq = so.sq.UseIndexForJoin(indexes...)

	return so
}

func (so *selectPostgreOrm) UseIndexForOrderBy(indexes ...string) orm.Select {
	so.sq = so.sq.UseIndexForOrderBy(indexes...)

	return so
}

func (so *selectPostgreOrm) For(s string, args ...interface{}) orm.Select {
	so.sq = so.sq.For(s, args...)

	return so
}

func (so *selectPostgreOrm) Scan(ctx context.Context, dest ...interface{}) error {
	err := so.sq.Scan(ctx, dest...)

	so.sq = nil

	return err
}

func (so *selectPostgreOrm) Exec(ctx context.Context, dest ...interface{}) (sql.Result, error) {
	res, err := so.sq.Exec(ctx, dest...)

	so.sq = nil

	return res, err
}

func (so *selectPostgreOrm) Count(ctx context.Context) (int, error) {
	c, err := so.sq.Count(ctx)

	so.sq = nil

	return c, err
}

func (so *selectPostgreOrm) ScanAndCount(ctx context.Context, dest ...interface{}) (int, error) {
	c, err := so.sq.ScanAndCount(ctx, dest...)

	so.sq = nil

	return c, err
}

type valuesPostgreOrm struct {
	vq *bun.ValuesQuery
}

func (vo *valuesPostgreOrm) AppendQuery(fmter schema.Formatter, b []byte) ([]byte, error) {
	return vo.vq.AppendQuery(fmter, b)
}

func (vo *valuesPostgreOrm) Column(columns ...string) orm.Values {
	vo.vq = vo.vq.Column(columns...)

	return vo
}

func (vo *valuesPostgreOrm) Value(column string, expr string, args ...interface{}) orm.Values {
	vo.vq = vo.vq.Value(column, expr, args...)

	return vo
}

func (vo *valuesPostgreOrm) WithOrder() orm.Values {
	vo.vq = vo.vq.WithOrder()

	return vo
}

type insertPostgreOrm struct {
	iq *bun.InsertQuery
}

func (io *insertPostgreOrm) With(name string, query orm.QueryAppender) orm.Insert {
	sqa, ok := query.(*selectPostgreOrm)
	if ok {
		io.iq = io.iq.With(name, sqa)
	}

	vqa, ok := query.(*valuesPostgreOrm)
	if ok {
		io.iq = io.iq.With(name, vqa)
	}

	return io
}

func (io *insertPostgreOrm) Model(model interface{}) orm.Insert {
	io.iq = io.iq.Model(model)

	return io
}

func (io *insertPostgreOrm) Column(columns ...string) orm.Insert {
	io.iq = io.iq.Column(columns...)

	return io
}

func (io *insertPostgreOrm) ExcludeColumn(columns ...string) orm.Insert {
	io.iq = io.iq.ExcludeColumn(columns...)

	return io
}

func (io *insertPostgreOrm) Table(tables ...string) orm.Insert {
	io.iq = io.iq.Table(tables...)

	return io
}

func (io *insertPostgreOrm) TableExpr(query string, args ...interface{}) orm.Insert {
	io.iq = io.iq.TableExpr(query, args...)

	return io
}

func (io *insertPostgreOrm) ModelTableExpr(query string, args ...interface{}) orm.Insert {
	io.iq = io.iq.ModelTableExpr(query, args...)

	return io
}

func (io *insertPostgreOrm) Value(column string, expr string, args ...interface{}) orm.Insert {
	io.iq = io.iq.Value(column, expr, args...)

	return io
}

func (io *insertPostgreOrm) On(s string, args ...interface{}) orm.Insert {
	io.iq = io.iq.On(s, args...)

	return io
}

func (io *insertPostgreOrm) Set(query string, args ...interface{}) orm.Insert {
	io.iq = io.iq.Set(query, args...)

	return io
}

func (io *insertPostgreOrm) Returning(query string, args ...interface{}) orm.Insert {
	io.iq = io.iq.Returning(query, args...)

	return io
}

func (io *insertPostgreOrm) Ignore() orm.Insert {
	io.iq = io.iq.Ignore()

	return io
}

func (io *insertPostgreOrm) Exec(ctx context.Context, dest ...interface{}) (sql.Result, error) {
	res, err := io.iq.Exec(ctx, dest...)

	io.iq = nil

	return res, err
}

type updatePostgreOrm struct {
	uq *bun.UpdateQuery
}

func (uo *updatePostgreOrm) With(name string, query schema.QueryAppender) orm.Update {
	sqa, ok := query.(*selectPostgreOrm)
	if ok {
		uo.uq = uo.uq.With(name, sqa)
	}

	vqa, ok := query.(*valuesPostgreOrm)
	if ok {
		uo.uq = uo.uq.With(name, vqa)
	}

	return uo
}

func (uo *updatePostgreOrm) Bulk() orm.Update {
	uo.uq = uo.uq.Bulk()

	return uo
}

func (uo *updatePostgreOrm) Model(model interface{}) orm.Update {
	uo.uq = uo.uq.Model(model)

	return uo
}

func (uo *updatePostgreOrm) Column(columns ...string) orm.Update {
	uo.uq = uo.uq.Column(columns...)

	return uo
}

func (uo *updatePostgreOrm) ExcludeColumn(columns ...string) orm.Update {
	uo.uq = uo.uq.ExcludeColumn(columns...)

	return uo
}

func (uo *updatePostgreOrm) Table(tables ...string) orm.Update {
	uo.uq = uo.uq.Table(tables...)

	return uo
}

func (uo *updatePostgreOrm) TableExpr(query string, args ...interface{}) orm.Update {
	uo.uq = uo.uq.TableExpr(query, args...)

	return uo
}

func (uo *updatePostgreOrm) ModelTableExpr(query string, args ...interface{}) orm.Update {
	uo.uq = uo.uq.ModelTableExpr(query, args...)

	return uo
}

func (uo *updatePostgreOrm) Value(column string, query string, args ...interface{}) orm.Update {
	uo.uq = uo.uq.Value(column, query, args...)

	return uo
}

func (uo *updatePostgreOrm) Set(query string, args ...interface{}) orm.Update {
	uo.uq = uo.uq.Set(query, args...)

	return uo
}

func (uo *updatePostgreOrm) SetColumn(column string, query string, args ...interface{}) orm.Update {
	uo.uq = uo.uq.SetColumn(column, query, args...)

	return uo
}

func (uo *updatePostgreOrm) OmitZero() orm.Update {
	uo.uq = uo.uq.OmitZero()

	return uo
}

func (uo *updatePostgreOrm) Where(query string, args ...interface{}) orm.Update {
	uo.uq = uo.uq.Where(query, args...)

	return uo
}

func (uo *updatePostgreOrm) WhereAllWithDeleted() orm.Update {
	uo.uq = uo.uq.WhereAllWithDeleted()

	return uo
}

func (uo *updatePostgreOrm) WhereDeleted() orm.Update {
	uo.uq = uo.uq.WhereDeleted()

	return uo
}

func (uo *updatePostgreOrm) WhereGroup(sep string, fn func(orm.Update) orm.Update) orm.Update {
	uo.uq = uo.uq.WhereGroup(sep, func(sq *bun.UpdateQuery) *bun.UpdateQuery {
		return fn(uo).(*updatePostgreOrm).uq
	})

	return uo
}

func (uo *updatePostgreOrm) WhereOr(query string, args ...interface{}) orm.Update {
	uo.uq = uo.uq.WhereOr(query, args...)

	return uo
}

func (uo *updatePostgreOrm) WherePK(cols ...string) orm.Update {
	uo.uq = uo.uq.WherePK(cols...)

	return uo
}

func (uo *updatePostgreOrm) Returning(query string, args ...interface{}) orm.Update {
	uo.uq = uo.uq.Returning(query, args...)

	return uo
}

func (uo *updatePostgreOrm) Exec(ctx context.Context, dest ...interface{}) (sql.Result, error) {
	res, err := uo.uq.Exec(ctx, dest...)

	uo.uq = nil

	return res, err
}

type deletePostgreOrm struct {
	dq *bun.DeleteQuery
}

func (do *deletePostgreOrm) With(name string, query schema.QueryAppender) orm.Delete {
	sqa, ok := query.(*selectPostgreOrm)
	if ok {
		do.dq = do.dq.With(name, sqa)
	}

	vqa, ok := query.(*valuesPostgreOrm)
	if ok {
		do.dq = do.dq.With(name, vqa)
	}

	return do
}

func (do *deletePostgreOrm) ForceDelete() orm.Delete {
	do.dq = do.dq.ForceDelete()

	return do
}

func (do *deletePostgreOrm) Model(model interface{}) orm.Delete {
	do.dq = do.dq.Model(model)

	return do
}

func (do *deletePostgreOrm) Table(tables ...string) orm.Delete {
	do.dq = do.dq.Table(tables...)

	return do
}

func (do *deletePostgreOrm) TableExpr(query string, args ...interface{}) orm.Delete {
	do.dq = do.dq.TableExpr(query, args...)

	return do
}

func (do *deletePostgreOrm) ModelTableExpr(query string, args ...interface{}) orm.Delete {
	do.dq = do.dq.ModelTableExpr(query, args...)

	return do
}

func (do *deletePostgreOrm) Where(query string, args ...interface{}) orm.Delete {
	do.dq = do.dq.Where(query, args...)

	return do
}

func (do *deletePostgreOrm) WhereAllWithDeleted() orm.Delete {
	do.dq = do.dq.WhereAllWithDeleted()

	return do
}

func (do *deletePostgreOrm) WhereDeleted() orm.Delete {
	do.dq = do.dq.WhereDeleted()

	return do
}

func (do *deletePostgreOrm) WhereGroup(sep string, fn func(orm.Delete) orm.Delete) orm.Delete {
	do.dq = do.dq.WhereGroup(sep, func(dq *bun.DeleteQuery) *bun.DeleteQuery {
		return fn(do).(*deletePostgreOrm).dq
	})

	return do
}

func (do *deletePostgreOrm) WhereOr(query string, args ...interface{}) orm.Delete {
	do.dq = do.dq.WhereOr(query, args...)

	return do
}

func (do *deletePostgreOrm) WherePK(cols ...string) orm.Delete {
	do.dq = do.dq.WherePK(cols...)

	return do
}

func (do *deletePostgreOrm) Returning(query string, args ...interface{}) orm.Delete {
	do.dq = do.dq.Returning(query, args...)

	return do
}

func (do *deletePostgreOrm) Exec(ctx context.Context, dest ...interface{}) (sql.Result, error) {
	res, err := do.dq.Exec(ctx, dest...)

	do.dq = nil

	return res, err
}

type rawPostgreOrm struct {
	rq *bun.RawQuery
}

func (ro *rawPostgreOrm) Exec(ctx context.Context, dest ...interface{}) (sql.Result, error) {
	res, err := ro.rq.Exec(ctx, dest...)

	ro.rq = nil

	return res, err
}

func (ro *rawPostgreOrm) Scan(ctx context.Context, dest ...interface{}) error {
	err := ro.rq.Scan(ctx, dest...)

	ro.rq = nil

	return err
}

type createTablePostgreOrm struct {
	ctq *bun.CreateTableQuery
}

func (cto *createTablePostgreOrm) Model(model interface{}) orm.CreateTable {
	cto.ctq = cto.ctq.Model(model)

	return cto
}

func (cto *createTablePostgreOrm) Table(tables ...string) orm.CreateTable {
	cto.ctq = cto.ctq.Table(tables...)

	return cto
}

func (cto *createTablePostgreOrm) TableExpr(query string, args ...interface{}) orm.CreateTable {
	cto.ctq = cto.ctq.TableExpr(query, args...)

	return cto
}

func (cto *createTablePostgreOrm) ModelTableExpr(query string, args ...interface{}) orm.CreateTable {
	cto.ctq = cto.ctq.ModelTableExpr(query, args...)

	return cto
}

func (cto *createTablePostgreOrm) Temp() orm.CreateTable {
	cto.ctq = cto.ctq.Temp()

	return cto
}

func (cto *createTablePostgreOrm) IfNotExists() orm.CreateTable {
	cto.ctq = cto.ctq.IfNotExists()

	return cto
}

func (cto *createTablePostgreOrm) Varchar(n int) orm.CreateTable {
	cto.ctq = cto.ctq.Varchar(n)

	return cto
}

func (cto *createTablePostgreOrm) WithForeignKeys() orm.CreateTable {
	cto.ctq = cto.ctq.WithForeignKeys()

	return cto
}

func (cto *createTablePostgreOrm) ForeignKey(query string, args ...interface{}) orm.CreateTable {
	cto.ctq = cto.ctq.ForeignKey(query, args...)

	return cto
}

func (cto *createTablePostgreOrm) PartitionBy(query string, args ...interface{}) orm.CreateTable {
	cto.ctq = cto.ctq.PartitionBy(query, args...)

	return cto
}

func (cto *createTablePostgreOrm) TableSpace(tablespace string) orm.CreateTable {
	cto.ctq = cto.ctq.TableSpace(tablespace)

	return cto
}

func (cto *createTablePostgreOrm) Exec(ctx context.Context, dest ...interface{}) (sql.Result, error) {
	res, err := cto.ctq.Exec(ctx, dest...)

	cto.ctq = nil

	return res, err
}
