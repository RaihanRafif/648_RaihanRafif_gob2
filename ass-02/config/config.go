package config

import (
	"ass2/structs"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	username = "root"
	password = ""
	hostname = "127.0.0.1:3306"
	dbname   = "ass2_orders"
)

func DBInit() *gorm.DB {

	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local", username, password, hostname, dbname))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(structs.Item{}, structs.Order{})
	return db
}
