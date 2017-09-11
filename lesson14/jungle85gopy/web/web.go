package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

const age = 300

// User for sql row
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"-" xml:"-"`
	Note     string `json:"note"`
	Isadmin  bool   `json:"isadmin"`
}

// Login usage
func Login(w http.ResponseWriter, r *http.Request) {
	render(w, "login", nil)
}

// ------ RenderAdd and Add func ------//

// RenderAdd usage
func RenderAdd(w http.ResponseWriter, r *http.Request) {
	render(w, "add", nil)
}

// Create to add a user
func Create(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	passwd := fmt.Sprintf("%x", md5.Sum([]byte(r.FormValue("password"))))
	note := r.FormValue("note")
	isad := r.FormValue("isadmin")
	isadmin, _ := strconv.ParseBool(isad)

	log.Printf("add:%s,%s,%s,%v", name, passwd, note, isadmin)
	res, err := db.Exec("INSERT INTO user VALUES(NULL, ?, ?, ?, ?)", name, passwd, note, isadmin)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	log.Print(res.LastInsertId())
	log.Print(res.RowsAffected())
	http.Redirect(w, r, "/list", 302)
}

// ------ Delete func ------//

// Delete delete a user
func Delete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")

	db.Exec("DELETE FROM user WHERE id = ?", id)
	http.Redirect(w, r, "/list", 302)
}

// ------ RenderUpdate and Modify func ------//

// RenderUpdate render update html
func RenderUpdate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")

	var user User
	err := db.Get(&user, "SELECT * FROM user WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	render(w, "update", user)
}

// Modify update info of given user
func Modify(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")
	name := r.FormValue("name")
	password := fmt.Sprintf("%x", md5.Sum([]byte(r.FormValue("password"))))
	note := r.FormValue("note")
	isadminStr := r.FormValue("isadmin")
	isadmin, _ := strconv.ParseBool(isadminStr)
	log.Printf("mod:%s,%s,%s,%s,%v", id, name, password, note, isadmin)

	stmt, err := db.Prepare(`UPDATE user SET name=?,password=?,note=?,isadmin=? WHERE id=?`)
	res, err := stmt.Exec(name, password, note, isadmin, id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	log.Print(res.RowsAffected())
	http.Redirect(w, r, "/list", 302)
}

// List all user as table
func List(w http.ResponseWriter, r *http.Request) {
	var users []User
	err := db.Select(&users, "SELECT * FROM user")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	render(w, "list", users)
}

// ------ common func ------//

func render(w http.ResponseWriter, name string, data interface{}) {
	tplFile := filepath.Join("template", name+".tpl")
	tpl, err := template.ParseFiles(tplFile)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	err = tpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// CheckLogin usage
func CheckLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("user")
	passwd := r.FormValue("password")

	var user User
	err := db.Get(&user, "SELECT password FROM user WHERE name = ?", name)
	if err != nil {
		render(w, "login", "user not found")
		return
	}
	// 计算passwd的md5与刚才获取的值是否相同。不同则失败。
	if fmt.Sprintf("%x", md5.Sum([]byte(passwd))) != user.Password {
		render(w, "login", "bad password or username")
	}
	http.SetCookie(w, &http.Cookie{
		Name:   "user",
		Value:  name,
		MaxAge: age,
	})
	http.Redirect(w, r, "/list", 302)
}
