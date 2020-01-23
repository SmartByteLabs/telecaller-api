package database

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/rightjoin/fig"
)

// Manager is the handling the connection values.
type Manager struct {
	Host     string
	User     string
	Port     int
	Password string
	Database string
	Params   []string
	Debug    bool
}

// DatabaseManager handles the database connection and cred
var DatabaseManager Manager

func init() {
	fmt.Println("!!! database manager setup started !!!")
	// read database config from the config and env both
	DatabaseManager.Host = fig.String("database.host")
	DatabaseManager.User = fig.String("database.user")
	DatabaseManager.Port = fig.IntOr(3306, "database.post")
	DatabaseManager.Password = os.Getenv("DB_PASS")
	// if database password in not available in env the read from config
	if DatabaseManager.Password == "" {
		DatabaseManager.Password = fig.String("database.password")
	}
	DatabaseManager.Database = fig.String("database.database")
	DatabaseManager.Debug = fig.BoolOr(false, "database.debug")
	DatabaseManager.Params = []string{
		"charset=utf8",
		"parseTime=True",
		"loc=Local",
	}
}

// ConnectionString generates the connection string using dbinfo.
func (db Manager) ConnectionString(params ...string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", db.User, db.Password, db.Host,
		db.Port, db.Database, strings.Join(append(db.Params, params...), "&"))
}

// GetORM gives the pointer of the database.
func (db Manager) GetORM(params ...string) *gorm.DB {
	dbo := GetORMByConnectionString(db.ConnectionString(params...))
	if db.Debug {
		dbo.Debug()
	}
	return dbo
}

// GetORMByConnectionString connect database using connection string.
func GetORMByConnectionString(str string) *gorm.DB {
	dbo, err := gorm.Open("mysql", str)
	if err != nil {
		panic(err)
	}
	return dbo
}

// Exec executed the query correspond to the params.
func Exec(db *gorm.DB, query string, params ...interface{}) *mysql.MySQLError {
	return nil
}
