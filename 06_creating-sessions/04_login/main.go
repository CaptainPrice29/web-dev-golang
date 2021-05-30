package main

import (
	"net/http"
	"text/template"
	"golang.org/x/crypto/bcrypt"
	uuid "github.com/satori/go.uuid"
)
var tpl *template.Template
type user struct{
	UserName string
	Password []byte
	FirstName string
	LastName string
}
type ex struct{
	Alreadysignedup string
	Alreadyexist string
}
var users = make(map[string]user)
var sessions = make(map[string]string)
func init(){
	tpl = template.Must(template.ParseGlob("templates/*"))
//	bs, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
}
func main(){
	http.HandleFunc("/",foo)
	http.HandleFunc("/login",login)
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
func login(w http.ResponseWriter,req *http.Request){
	c, err := req.Cookie("session")
	if err != nil{
		http.Redirect(w,req,"/",http.StatusSeeOther)
		return
	}
    e := sessions[c.Value]
    _,ok := users[e]
    if  !ok{
		http.Redirect(w,req,"/",http.StatusSeeOther)
		return
    }

	if req.Method=="POST"{
		un := req.FormValue("uname")
		password := req.FormValue("password")
		u, ok := users[un]
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(password))
		if err != nil || !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		//start
    //	u, ok := users[un]
	//	if !ok {
	//		http.Error(w, "Username and/or password do not match", http.StatusForbidden)
	//		return
	//	}
		// does the entered password match the stored password?
	//	err := bcrypt.CompareHashAndPassword(u.Password, []byte(password))
	//	if err != nil {
	//		http.Error(w, "Username and/or password do not match", http.StatusForbidden)
	//		return
	//	}
		//end
		uid := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: uid.String(),
		}
		http.SetCookie(w, c)
		sessions[c.Value] = un
		http.Redirect(w, req, "/info", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w,"login.html",nil)
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
		
		if _, ok := users[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

	uid := uuid.NewV4()
	c = &http.Cookie{Name: "session",Value: uid.String()}
	http.SetCookie(w,c)
	// storing user in Users
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	u := user{un,bs,fname,lname}
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
	}else{
		tpl.ExecuteTemplate(w,"index.html",nil)
	}
}