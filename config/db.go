package config

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// buat variable untuk diakses diluar package
var Db *gorm.DB

func init() {
	var err error
	godotenv.Load(".env")
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	Db, err = gorm.Open(os.Getenv("DB_DRIVER"), DBURL)

	if err != nil {
		panic(err)
	}

	// set this to 'true' to see sql logs
	Db.LogMode(true)

}
