package configs

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// mysql database konektor
func MySQLConn() *gorm.DB {

	configDB := map[string]string{
		"DB_Username": Config("DB_USERNAME"),
		"DB_Password": Config("DB_PASSWORD"),
		"DB_Port":     Config("DB_PORT"),
		"DB_Host":     Config("DB_ADDRESS"),
		"DB_Name":     Config("DB_NAME"),
	}
	fmt.Println(configDB)
	fmt.Println("Server running well...")

	openConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configDB["DB_Username"],
		configDB["DB_Password"],
		configDB["DB_Host"],
		configDB["DB_Port"],
		configDB["DB_Name"])
	db, e := gorm.Open(mysql.Open(openConnection), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	return db

}
