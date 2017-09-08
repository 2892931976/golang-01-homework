package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"html/template"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"sync"
)

var (
	_     http.Handler = &Counter{}
	dbx   *sqlx.DB
	store = sessions.NewCookieStore([]byte("something-very-secret"))
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"-"`
	Note     string `json:"note"`
	IsAdmin  int    `json:"isadmin"`
}

type Counter struct {
	h     http.Handler
	count map[string]int
	mutex sync.Mutex
}

func NewCounter(h http.Handler) *Counter {
	return &Counter{
		h:     h,
		count: make(map[string]int),
	}
}

func (c *Counter) GetCounter(w http.ResponseWriter, r *http.Request) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for path, count := range c.count {
		fmt.Fprintf(w, "%s\t%d\n", path, count)
	}
}

func (c *Counter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.mutex.Lock()
	c.count[r.URL.Path]++
	c.mutex.Unlock()
	c.h.ServeHTTP(w, r)
}

func render(w http.ResponseWriter, name string, data interface{}) {
	path := filepath.Join("template", name)
	tpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	render(w, "login.html", nil)
}

func CheckLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	password := r.FormValue("password")

	var user User
	err := dbx.Get(&user, "select password from user where name = ?", name)
	if err != nil {
		render(w, "login.html", "user not found")
		return
	}

	if fmt.Sprintf("%x", md5.Sum([]byte(password))) == user.Password {
		session, _ := store.New(r, "test")
		session.Values["user"] = name
		session.Save(r, w)
		http.Redirect(w, r, "/list/", http.StatusFound)
	} else {
		render(w, "login.html", "bad password or username")
	}
}

func List(w http.ResponseWriter, r *http.Request) {
	var users []User
	err := dbx.Select(&users, "select * from user")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	render(w, "list.html", users)
}

func Add(w http.ResponseWriter, r *http.Request) {
	render(w, "add.html", nil)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	password := fmt.Sprintf("%x", md5.Sum([]byte(r.FormValue("password"))))
	note := r.FormValue("note")
	isadmin := r.FormValue("isadmin")
	if isadmin == "" {
		isadmin = "0"
	}
	res, err := dbx.Exec("INSERT INTO user (name, password, note, isadmin) "+
		"				VALUES(?, ?, ?, ?)", name, password, note, isadmin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(res.LastInsertId())
	log.Println(res.RowsAffected())
	http.Redirect(w, r, "/list/", http.StatusFound)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")
	_, err := dbx.Exec("DELETE FROM user WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/list/", http.StatusFound)
}

func Change(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")
	var user User
	err := dbx.Get(&user, "select * from user where id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(user)
	render(w, "update.html", user)
}

func ChangeUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")
	name := r.FormValue("name")
	note := r.FormValue("note")
	_, err := dbx.Exec("update user set name = ?, note = ? where id = ?",
		name, note, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/list/", http.StatusFound)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello http")
}

func NeedLogin(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "test")
		_, ok := session.Values["user"]
		if !ok {
			render(w, "login.html", "登陆过期")
			return
		}
		h(w, r)
	}
}

func Users(w http.ResponseWriter, r *http.Request) {
	var resp Response
	var users []User
	err := dbx.Select(&users, "select * from user")
	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Msg = err.Error()
	} else {
		resp.Code = http.StatusOK
		resp.Data = users
	}
	buf, _ := json.Marshal(&resp)
	w.Write(buf)
}

func main() {
	var err error
	dbx, err = sqlx.Open("mysql", "root:123456@tcp(192.168.3.50:3306)/go")
	if err != nil {
		log.Fatalln(err)
	}
	err = dbx.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	loginCounter := NewCounter(http.HandlerFunc(Login))
	http.Handle("/login/", loginCounter)
	http.HandleFunc("/lcount/", loginCounter.GetCounter)

	http.Handle("/checklogin/", http.HandlerFunc(CheckLogin))
	http.Handle("/hello/", NeedLogin(http.HandlerFunc(Hello)))
	http.Handle("/list/", NeedLogin(http.HandlerFunc(List)))
	http.Handle("/add/", NeedLogin(http.HandlerFunc(Add)))
	http.Handle("/adduser/", NeedLogin(http.HandlerFunc(AddUser)))
	http.Handle("/change/", NeedLogin(http.HandlerFunc(Change)))
	http.Handle("/changeuser/", NeedLogin(http.HandlerFunc(ChangeUser)))
	http.Handle("/delete/", NeedLogin(http.HandlerFunc(Delete)))
	http.Handle("/users/", NeedLogin(http.HandlerFunc(Users)))
	http.Handle("/static/", http.FileServer(http.Dir(".")))

	h := handlers.LoggingHandler(os.Stderr, http.DefaultServeMux)
	c := NewCounter(h)
	http.HandleFunc("/count/", c.GetCounter)
	log.Fatal(http.ListenAndServe(":8000", c))
}
