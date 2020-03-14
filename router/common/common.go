package common

import (
	"encoding/json"
	"log"
	"net/http"
)

type ReturnMsg struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func SuccessReturn(data interface{}, w http.ResponseWriter) {
	ret, err := json.Marshal(ReturnMsg{
		Code: 0,
		Data: data,
		Msg:  "success",
	})
	if err != nil {
		log.Println(err)
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(200)
	_, _ = w.Write(ret)
}

func ErrorReturn(code int, httpCode int, data interface{}, msg string, w http.ResponseWriter) {
	ret, err := json.Marshal(ReturnMsg{
		Code: code,
		Data: data,
		Msg:  msg,
	})
	if err != nil {
		log.Println(err)
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(httpCode)
	_, _ = w.Write(ret)
}
