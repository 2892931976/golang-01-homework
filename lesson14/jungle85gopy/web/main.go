package main

import (
	"log"
	"net/http"
	"os"

	_ "net/http/pprof"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

// NeedLogin for check login
func NeedLogin(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("user")
		if err != nil {
			render(w, "login", "login out of time")
			return
		}
		h(w, r)
	}
}

func main() {
	var err error
	db, err = sqlx.Open("mysql", "golang:golang@tcp(59.110.12.72:3306)/go")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	loginCounter := NewCounter(http.HandlerFunc(Login))
	http.Handle("/login", loginCounter)
	http.Handle("/", loginCounter)
	http.HandleFunc("/loginCounter", loginCounter.GetCounter)
	http.HandleFunc("/checkLogin", CheckLogin)

	// /add render() html; post action to /create
	http.HandleFunc("/add", NeedLogin(RenderAdd))
	http.HandleFunc("/create", NeedLogin(Create))

	// /delete
	http.HandleFunc("/delete", NeedLogin(Delete))

	// list
	http.HandleFunc("/list", NeedLogin(List))

	// modify
	http.HandleFunc("/update", NeedLogin(RenderUpdate))
	http.HandleFunc("/modify", NeedLogin(Modify))

	// file server
	http.Handle("/static/", http.FileServer(http.Dir(".")))

	// api
	http.HandleFunc("/users", Users)

	// middleware
	mux := handlers.LoggingHandler(os.Stderr, http.DefaultServeMux)
	c := NewCounter(mux)
	http.HandleFunc("/counter", c.GetCounter)

	log.Fatal(http.ListenAndServe(":8090", c))
}
