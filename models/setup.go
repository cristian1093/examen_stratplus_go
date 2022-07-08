package models

import (
	"bytes"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tools/viper"
)

var (
	databaseGorm DatabaseGorm
	// taskStatusGorm map[string]int
	// err            error
)

type DatabaseGorm struct {
	db *gorm.DB
	// codec *branca.Branca
}

func ConectGorm() {
	// Please define your username and password for MySQL.
	databaseGorm.db, err = gorm.Open("mysql", urlCreate2())
	// NOTE: See we’re using = to assign the global var
	// instead of := which would assign it only in this function

	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
	}
}

func urlCreate2() (url string) {
	var u bytes.Buffer

	u.WriteString(viper.GetEnv("MYSQL_USER"))
	u.WriteString(":")
	u.WriteString(viper.GetEnv("MYSQL_PASSWORD"))
	u.WriteString("@tcp(")
	u.WriteString(viper.GetEnv("MYSQL_HOST"))
	u.WriteString(":")
	u.WriteString(viper.GetEnv("MYSQL_PORT"))
	u.WriteString(")/")
	u.WriteString(viper.GetEnv("MYSQL_DB"))

	url = u.String()

	return
}
