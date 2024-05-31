package dao

import (
	"time"

	"github.com/uptrace/bun"
)

type (
	Farmer struct {
		bun.BaseModel `bun:"farmers"`
		ID            uint64    `bun:"id,pk,autoincrement"`
		Name          string    `bun:"name,notnull"`
		BirthDate     time.Time `bun:"birth_date,notnull"`
		CreatedAt     time.Time `bun:"created_at,notnull"`
		UpdatedAt     time.Time `bun:"updated_at,notnull"`
	}

	Farmers []Farmer
)

type FarmerPaging struct {
	Limit  int
	Offset int
	Count  int
	Sorts  []string

	Items Farmers
}
