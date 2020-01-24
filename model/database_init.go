package model

import (
	"fmt"

	"github.com/princeparmar/telecaller-app/database"
	"github.com/rightjoin/dorm"
)

func init() {
	fmt.Println("!!! model init started !!!")
	tempManager := database.DatabaseManager
	tempManager.Database = ""
	dbo := tempManager.GetORM()
	defer dbo.Close()
	err := dbo.Exec("CREATE DATABASE IF NOT EXISTS " + database.DatabaseManager.Database + ";").Error
	if err != nil {
		panic(err)
	}

	dbo = database.DatabaseManager.GetORM()
	defer dbo.Close()

	dorm.OverrideDB = dbo
	dorm.BuildSchema(
		&Language{},
		&Locality{},
		&Proof{},
		&JobType{},
		&Contact{},
		&ContactJobType{},
		&ContactLanguage{},
		&ContactProof{},
	)
}
