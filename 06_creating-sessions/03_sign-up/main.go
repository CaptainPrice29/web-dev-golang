package main

import (
	"net/http"
	"text/template"

	uuid "github.com/satori/go.uuid"
)
var tpl *template.Template
type user struct{
	UserName, Password, FirstName,LastName string
}
type ex struct{
	Alreadysignedup string
	Alreadyexist string
}
var users = make(map[string]user)
var sessions = make(map[string]string)
func init(){
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func main(){
	http.HandleFunc("/",foo)
	http.HandleFunc("/signup",signup)
	http.HandleFunc("/info",info)
	http.ListenAndServe(":8080",nil)
}
func foo(w http.ResponseWriter,req *http.Request){
	_,err:=req.Cookie("session")
	if err != nil{
		tpl.ExecuteTemplate(w,"index.html",nil)
		return
	}
	p := ex{
		Alreadysignedup: "You are already signed up",
	}
	tpl.ExecuteTemplate(w,"index.html",p)
	
}
func signup(w http.ResponseWriter,req *http.Request){
	//check if already signup,look for cookie
	c,err := req.Cookie("session")
	if err == nil{
		un:=sessions[c.Value]
		u:=users[un]
		tpl.ExecuteTemplate(w,"info.html",u)
		return		
	}
	
	if req.Method=="POST"{
		un := req.FormValue("uname")
		password :=req.FormValue("password")
		fname := req.FormValue("fname")
		lname := req.FormValue("lname")
		
	

	uid := uuid.NewV4()
	c = &http.Cookie{Name: "session",Value: uid.String()}
	http.SetCookie(w,c)

	u := user{un,password,fname,lname}
	sessions[c.Value] = un
	users[un] = u

	http.Redirect(w,req,"/info",http.StatusSeeOther)
    }
	tpl.ExecuteTemplate(w,"signup.html",nil)
}
func info(w http.ResponseWriter,req *http.Request){
	c,err := req.Cookie("session")
	if err == nil{
		un:=sessions[c.Value]
		u:=users[un]
		tpl.ExecuteTemplate(w,"info.html",u)
		return
	}
		tpl.ExecuteTemplate(w,"index.html",nil)
	
}
