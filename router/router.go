package router

import (
	"net/http"
)

func Setup() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/login", login)

	return mux
}

func index(w http.ResponseWriter, r *http.Request) {
}
