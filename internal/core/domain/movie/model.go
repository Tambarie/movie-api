package domain

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        string    `sql:"type:uuid; default:uuid_generate_v4();size:100; not null"`
	CreatedAT time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New().String()
	if m.ID == "" {
		err = errors.New("can't save invalid data")
	}
	return
}
