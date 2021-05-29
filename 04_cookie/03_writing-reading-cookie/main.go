package main

import (
	"fmt"
	"net/http"
)
func main(){
	http.HandleFunc("/",set)
	http.HandleFunc("/read",read)
	http.ListenAndServe(":8080",nil)
}
func set(w http.ResponseWriter,req *http.Request){
	http.SetCookie(w,&http.Cookie{Name:"captnCookie",Value:"some unique value",MaxAge: 120})
	fmt.Fprint(w,"cookie written")
}
func read(w http.ResponseWriter,req *http.Request){
	c,err:=req.Cookie("captnCookie")
	if err != nil{
		http.Error(w,err.Error(),http.StatusNotFound)
	}
	fmt.Fprint(w,"cookie-",c)
}