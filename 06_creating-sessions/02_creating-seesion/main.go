package main

import (
	"net/http"
	"text/template"

	uuid "github.com/satori/go.uuid"
)
type user struct{
	UserName string
	FirstName string
	LastName string
}
var tpl *template.Template
var users =make(map[string]user)
var sessions = make(map[string]string)
func init(){
	tpl = template.Must(template.ParseGlob("template/*"))
}
func main(){
	http.HandleFunc("/",foo)
	http.HandleFunc("/info",info)
	http.ListenAndServe(":8080",nil)
}
func foo(w http.ResponseWriter,req *http.Request){
	c,err:=req.Cookie("session")
	if err != nil{
		id := uuid.NewV4()
		c = &http.Cookie{Name: "session",Value: id.String()}
		http.SetCookie(w,c)
	}
	var u user
	if un,ok:=sessions[c.Value];ok{
		u = users[un]
	}
	if req.Method == "POST"{
		un := req.FormValue("uname")
		fn := req.FormValue("fname")
		ln := req.FormValue("lname")
		u = user{un,fn,ln}
		sessions[c.Value] = un
		users[un] = u
	}
	tpl.ExecuteTemplate(w,"index.html",u)
}
func info(w http.ResponseWriter,req *http.Request){
	 c,err := req.Cookie("session")
	 if err != nil{
		 http.Redirect(w,req,"/",http.StatusSeeOther)
		 return
	 }
	 un,ok := sessions[c.Value]
	 if !ok{
		 http.Redirect(w,req,"/",http.StatusSeeOther)
	 }
	 u := users[un]	
	 tpl.ExecuteTemplate(w,"info.html",u)
}