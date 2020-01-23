package model

import "time"

// Timestamps is handling the created at and updated at fields for each model.
type Timestamps struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

// PrimaryKey is handling the auto-inc primary id.
type PrimaryKey struct {
	ID uint `gorm:"id"`
}
