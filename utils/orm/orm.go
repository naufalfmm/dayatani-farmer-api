package orm

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun/schema"
)

type (
	QueryAppender interface {
		AppendQuery(fmter schema.Formatter, b []byte) ([]byte, error)
	}

	Select interface {
		AppendQuery(fmter schema.Formatter, b []byte) ([]byte, error)
		With(name string, query QueryAppender) Select
		Union(other Select) Select
		UnionAll(other Select) Select
		Model(model interface{}) Select
		Column(columns ...string) Select
		ColumnExpr(query string, args ...interface{}) Select
		ExcludeColumn(columns ...string) Select
		Table(tables ...string) Select
		TableExpr(query string, args ...interface{}) Select
		ModelTableExpr(query string, args ...interface{}) Select
		Join(join string, args ...interface{}) Select
		JoinOn(cond string, args ...interface{}) Select
		JoinOnOr(cond string, args ...interface{}) Select
		Where(query string, args ...interface{}) Select
		WhereAllWithDeleted() Select
		WhereDeleted() Select
		WhereGroup(sep string, fn func(Select) Select) Select
		WhereOr(query string, args ...interface{}) Select
		WherePK(cols ...string) Select
		Group(columns ...string) Select
		GroupExpr(group string, args ...interface{}) Select
		Order(orders ...string) Select
		OrderExpr(query string, args ...interface{}) Select
		Having(having string, args ...interface{}) Select
		Limit(n int) Select
		Offset(n int) Select
		Relation(name string, apply ...func(Select) Select) Select
		UseIndex(indexes ...string) Select
		UseIndexForGroupBy(indexes ...string) Select
		UseIndexForJoin(indexes ...string) Select
		UseIndexForOrderBy(indexes ...string) Select
		For(s string, args ...interface{}) Select
		Scan(ctx context.Context, dest ...interface{}) error
		Exec(ctx context.Context, dest ...interface{}) (sql.Result, error)
		Count(ctx context.Context) (int, error)
		ScanAndCount(ctx context.Context, dest ...interface{}) (int, error)
	}

	Values interface {
		AppendQuery(fmter schema.Formatter, b []byte) ([]byte, error)
		Column(columns ...string) Values
		Value(column string, expr string, args ...interface{}) Values
		WithOrder() Values
	}

	Insert interface {
		With(name string, query QueryAppender) Insert
		Model(model interface{}) Insert
		Column(columns ...string) Insert
		ExcludeColumn(columns ...string) Insert
		Table(tables ...string) Insert
		TableExpr(query string, args ...interface{}) Insert
		ModelTableExpr(query string, args ...interface{}) Insert
		Value(column string, expr string, args ...interface{}) Insert
		On(s string, args ...interface{}) Insert
		Set(query string, args ...interface{}) Insert
		Returning(query string, args ...interface{}) Insert
		Ignore() Insert
		Exec(ctx context.Context, dest ...interface{}) (sql.Result, error)
	}

	Update interface {
		With(name string, query schema.QueryAppender) Update
		Bulk() Update
		Model(model interface{}) Update
		Column(columns ...string) Update
		ExcludeColumn(columns ...string) Update
		Table(tables ...string) Update
		TableExpr(query string, args ...interface{}) Update
		ModelTableExpr(query string, args ...interface{}) Update
		Value(column string, query string, args ...interface{}) Update
		Set(query string, args ...interface{}) Update
		SetColumn(column string, query string, args ...interface{}) Update
		OmitZero() Update
		Where(query string, args ...interface{}) Update
		WhereAllWithDeleted() Update
		WhereDeleted() Update
		WhereGroup(sep string, fn func(Update) Update) Update
		WhereOr(query string, args ...interface{}) Update
		WherePK(cols ...string) Update
		Returning(query string, args ...interface{}) Update
		Exec(ctx context.Context, dest ...interface{}) (sql.Result, error)
	}

	Delete interface {
		With(name string, query schema.QueryAppender) Delete
		ForceDelete() Delete
		Model(model interface{}) Delete
		Table(tables ...string) Delete
		TableExpr(query string, args ...interface{}) Delete
		ModelTableExpr(query string, args ...interface{}) Delete
		Where(query string, args ...interface{}) Delete
		WhereAllWithDeleted() Delete
		WhereDeleted() Delete
		WhereGroup(sep string, fn func(Delete) Delete) Delete
		WhereOr(query string, args ...interface{}) Delete
		WherePK(cols ...string) Delete
		Returning(query string, args ...interface{}) Delete
		Exec(ctx context.Context, dest ...interface{}) (sql.Result, error)
	}

	Raw interface {
		Exec(ctx context.Context, dest ...interface{}) (sql.Result, error)
		Scan(ctx context.Context, dest ...interface{}) error
	}

	CreateTable interface {
		Model(model interface{}) CreateTable
		Table(tables ...string) CreateTable
		TableExpr(query string, args ...interface{}) CreateTable
		ModelTableExpr(query string, args ...interface{}) CreateTable
		Temp() CreateTable
		IfNotExists() CreateTable
		Varchar(n int) CreateTable
		WithForeignKeys() CreateTable
		ForeignKey(query string, args ...interface{}) CreateTable
		PartitionBy(query string, args ...interface{}) CreateTable
		TableSpace(tablespace string) CreateTable
		Exec(ctx context.Context, dest ...interface{}) (sql.Result, error)
	}

	Orm interface {
		NewSelect() Select
		NewValues(model interface{}) Values
		NewInsert() Insert
		NewUpdate() Update
		NewDelete() Delete
		NewRaw(query string, args ...interface{}) Raw
		NewCreateTable() CreateTable

		Begin() (Orm, error)
		BeginTx(ctx context.Context, opts *sql.TxOptions) (Orm, error)
		Commit() error
		Rollback() error
		RunInTx(ctx context.Context, opts *sql.TxOptions, fn func(ctx context.Context, tx Orm) error) error

		Exec(query string, args ...interface{}) (sql.Result, error)
		ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)

		Ping() error
		PingContext(ctx context.Context) error
	}
)
