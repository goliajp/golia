package libgorm

import (
	"time"
)

type Model struct {
	ID        uint       `json:"id" gorm:"primary_key; index"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
}
