package libraries

import (
	"fmt"
	"log"
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func getMySQLConnectionString() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Env("DB_USER"),
		Env("DB_PWD"),
		Env("DB_HOST"),
		Env("DB_PORT"),
		Env("DB_NAME"))

	return dsn
}

func getSqlServerConnectionString() string {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		Env("DB_USER"),
		url.QueryEscape(Env("DB_PWD")),
		Env("DB_HOST"),
		Env("DB_PORT"),
		Env("DB_NAME"))

	return dsn
}

func NewDB(params ...string) *gorm.DB {
	var err error

	if Env("DB_DRIVER") == "sqlsrv" {
		DB, err = gorm.Open(sqlserver.Open(getSqlServerConnectionString()), &gorm.Config{})
	} else {
		DB, err = gorm.Open(mysql.Open(getMySQLConnectionString()), &gorm.Config{})
	}

	if err != nil {
		log.Panic(err)
	}

	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}
