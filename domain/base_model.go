package domain

import "time"

type BaseModel struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}
