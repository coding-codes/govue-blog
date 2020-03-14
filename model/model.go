package model

import (
	"database/sql"
	"fmt"
	"github.com/coding-codes/govue-blog/config"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var dbc = config.Cfg.Database

func mysql() string {
	return fmt.Sprintf("%s:%s@%s(%s)/%s?charset=%s&parseTime=%s&loc=%s",
		dbc.User, dbc.Password, dbc.Protocol, dbc.Host, dbc.Name, dbc.Charset, dbc.ParseTime, dbc.Loc)
}

var db *sql.DB

func init() {
	var err error
	db, err := sql.Open(dbc.Dialect, mysql())
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(dbc.MaxIdleConns)
	db.SetMaxOpenConns(dbc.MaxOpenConns)
	fmt.Println(err)
	fmt.Printf("Connect %s successful.", dbc.Dialect)

}

func hash(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hash)
}
