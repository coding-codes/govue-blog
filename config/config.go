package config

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

type config struct {
	Server   server
	Database database
}

type server struct {
	Mode         string
	Addr         string
	ReadTimeout  int
	WriteTimeout int
}

type database struct {
	Dialect      string
	User         string
	Password     string
	Host         string
	Name         string
	Protocol     string
	Charset      string
	ParseTime    string
	Loc          string
	MaxIdleConns int
	MaxOpenConns int
}

var Cfg config

func init() {
	if _, err := toml.DecodeFile("./config.toml", &Cfg); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Init configure successful.")
}
