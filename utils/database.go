package utils

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// SaveWithDuplicateIgnore :
func SaveWithDuplicateIgnore(dbo *gorm.DB, i interface{}) error {
	err := dbo.Save(i).Error
	if err != nil {
		if mysqlErr, _ := err.(*mysql.MySQLError); mysqlErr != nil && mysqlErr.Number == 1062 {
			return nil
		}
	}
	return err
}
