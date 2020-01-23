package database

import (
	"fmt"

	"github.com/princeparmar/telecaller-app/model"
)

func init() {
	fmt.Println("!!! model init started !!!")
	tempManager := DatabaseManager
	tempManager.Database = ""
	dbo := tempManager.GetORM()
	defer dbo.Close()
	err := dbo.Exec("CREATE DATABASE IF NOT EXISTS " + DatabaseManager.Database + ";").Error
	if err != nil {
		panic(err)
	}

	dbo = DatabaseManager.GetORM()
	defer dbo.Close()
	err = dbo.AutoMigrate(
		&model.Language{},
		&model.Locality{},
		&model.Proof{},
		&model.JobType{},
		&model.Contact{},
		&model.ContactJobType{},
		&model.ContactLanguage{},
		&model.ContactProof{},
	).Error
	if err != nil {
		panic(err)
	}
}
