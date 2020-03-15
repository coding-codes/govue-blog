package main

import (
	"fmt"
	"github.com/coding-codes/govue-blog/config"
	"github.com/coding-codes/govue-blog/router"
	"log"
	"net/http"
	"time"
)

func main() {
	sc := config.Cfg.Server
	server := &http.Server{
		Addr:           sc.Addr,
		Handler:        router.Setup(),
		ReadTimeout:    time.Duration(sc.ReadTimeout * int(time.Second)),
		WriteTimeout:   time.Duration(sc.WriteTimeout * int(time.Second)),
		MaxHeaderBytes: 0,
	}
	fmt.Printf("Server running on %s\n", sc.Addr)
	log.Fatal(server.ListenAndServe())
}
