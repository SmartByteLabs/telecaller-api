package main

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/princeparmar/telecaller-app/database"
)

func main() {
	fmt.Println("package is under construction")
	dbo := database.DatabaseManager.GetORM()
	fmt.Println(dbo.DB().Ping())
}
