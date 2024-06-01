package model

import "time"

type MigrationLog struct {
	ID        string    `gorm:"primaryKey;not null;type:varchar(15)"`
	Name      string    `gorm:"not null;type:varchar(255)"`
	MigrateAt time.Time `gorm:"not null;type:datetime"`
}

func (log MigrationLog) TableName() string {
	return "migration_logs"
}
