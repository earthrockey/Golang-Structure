package config

import (
	"fmt"

	"github.com/earthrockey/Golang-Structure/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	user     = "root"
	password = ""
	host     = "localhost"
	dbname   = "echo_gorm"
)

func ConnectDB() (*gorm.DB, error) {
	var err error
	var client *gorm.DB
	str := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbname)
	if client, err = gorm.Open("mysql", str); err != nil {
		fmt.Println("...ConnectDB-Error...")
		fmt.Println(err.Error())
		return nil, err
	}
	return client, nil
}

func CreateTable() {
	db, err := ConnectDB()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&model.User{}, &model.Book{}, &model.Achievement{})

}
