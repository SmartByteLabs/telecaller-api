package model

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

// User is user login details.
type User struct {
	gorm.Model
	Name sql.NullInt64
}
