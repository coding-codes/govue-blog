package main

import (
	"github.com/coding-codes/router"
	"github.com/coding-codes/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	gin.SetMode(utils.ServerInfo.RunMode)

	r := router.InitRouter()

	s := &http.Server{
		Addr:           utils.ServerInfo.ServerAddr,
		Handler:        r,
		ReadTimeout:    utils.ServerInfo.ReadTimeout,
		WriteTimeout:   utils.ServerInfo.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	if e := s.ListenAndServe(); e != nil && e != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", e)
	}

}
