package models

import (
	"github.com/golangkit/formatime"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Base struct {
	ID        uuid.UUID           `gorm:"type:uuid;primary_key;"`
	CreatedAt formatime.Timestamp `gorm:"timestamp"`
	UpdatedAt formatime.Timestamp
	DeletedAt *formatime.Timestamp `sql:"index"`
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", uid)
}
