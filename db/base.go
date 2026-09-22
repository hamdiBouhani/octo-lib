package db

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _IObject = new(BaseObject)

type IObject interface {
	GetBaseObject() *BaseObject
	GetID() uuid.UUID
}
type BaseObject struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (b *BaseObject) GetBaseObject() *BaseObject {
	return b
}

func (b *BaseObject) GetID() uuid.UUID {
	return b.ID
}

func (b *BaseObject) MarkUpdate(tx *gorm.DB) (err error) {
	b.UpdatedAt = time.Now()
	return nil
}

func (b *BaseObject) MarkDeleted() {
	b.DeletedAt.Time = time.Now()
	b.DeletedAt.Valid = true
}

func NewBaseObject[T string | uuid.UUID](id T) *BaseObject {
	switch value := any(id).(type) {
	case string:
		return &BaseObject{ID: uuid.MustParse(value)}
	case uuid.UUID:
		return &BaseObject{ID: value}
	default:
		return &BaseObject{}
	}
}
