// config.go
package config

import (
	"goWeb/entities"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*func GetDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "crudgolang"

	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	return
}*/

var DB *gorm.DB

func GetDb() {

	connection, err := gorm.Open(mysql.Open("root:@/crudgolang"))

	if err != nil {
		panic("could not connect to the database")
	}
	DB = connection
	connection.AutoMigrate(&entities.Product{})
	return
}
