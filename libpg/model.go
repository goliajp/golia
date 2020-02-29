package libpg

import (
	"time"
)

type Model struct {
	Id        Id         `json:"id" gorm:"primary_key; index"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
}
