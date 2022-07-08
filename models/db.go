package models

import (
	"bytes"
	"database/sql"
	"log"

	"github.com/hako/branca"
	"github.com/tools/viper"

	// "./../config"
	// CARGA DEL DRIVER SQL
	_ "github.com/go-sql-driver/mysql"
)

var (
	database   Database
	taskStatus map[string]int
	err        error
)

//Database encapsulates database
type Database struct {
	db    *sql.DB
	codec *branca.Branca
}

func init() {

	database.db, err = sql.Open("mysql", urlCreate())

	// }
	if err != nil {
		log.Fatal(err)
	}
}

func urlCreate() (url string) {
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

//Close function closes this database connection
func Close() {
	database.db.Close()

}
