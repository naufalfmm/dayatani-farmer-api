package db

import (
	"context"

	"github.com/naufalfmm/dayatani-farmer-api/utils/orm"
)

type DB struct {
	o   orm.Orm
	trx *Trx
}

func (db *DB) StartTrx(ctx context.Context) {
	if db.trx == nil {
		db.trx = &Trx{}
	}

	db.trx.Begin(ctx, db.o)
}

func (db *DB) CommitTrx(ctx context.Context) {
	if db.trx == nil {
		return
	}

	db.trx.Commit()
	db.trx = nil
}

func (db *DB) RollbackTrx(ctx context.Context) {
	if db.trx == nil {
		return
	}

	db.trx.Rollback()
	db.trx = nil
}

func (db *DB) GetDB() orm.Orm {
	if db.trx == nil {
		return db.o
	}

	return db.trx.trx
}
