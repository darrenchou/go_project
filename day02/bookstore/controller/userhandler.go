package controller

import (
	"day02/bookstore/dao"
	"html/template"
	"net/http"
)

//Login
func Login(w http.ResponseWriter, r *http.Request) {
	//get user username and password
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	//check username and password
	user, _ := dao.CheckUserNameAndPassword(username, password)
	if user.ID > 0 {

		//if username and password are correct
		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w, "")
	} else {
		//not correct
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "Username or password is not correct!")
	}
}

//Regist
func Regist(w http.ResponseWriter, r *http.Request) {
	//get username and password
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	//check username and password
	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {
		//username already exsits
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "Username already exists!")
	} else {
		//save user info
		dao.SaveUser(username, password, email)

		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w, "")
	}
}

//CheckUserName check username already exists via ajax
func CheckUserName(w http.ResponseWriter, r *http.Request) {
	//get username
	username := r.PostFormValue("username")

	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {

		w.Write([]byte("Username already exists! "))
	} else {

		w.Write([]byte("<font style='color:green'>Username is available!</font>"))
	}
}
