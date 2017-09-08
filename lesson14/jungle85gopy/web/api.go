package main

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

// Response for resp
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Users api
func Users(w http.ResponseWriter, r *http.Request) {
	var users []User
	var resp Response
	var buf []byte

	r.ParseForm()
	format := r.FormValue("f")

	err := db.Select(&users, "SELECT * FROM user")
	if err != nil {
		resp.Code = 500
		resp.Msg = err.Error()
	} else {
		resp.Code = 200
		resp.Data = users
	}
	switch format {
	case "xml":
		w.Header().Set("Content-Type", "text/xml")
		buf, _ = xml.Marshal(&resp)
	case "json":
	default:
		w.Header().Set("Content-Type", "application/json")
		buf, _ = json.Marshal(&resp)
	}
	w.Write(buf)
}
