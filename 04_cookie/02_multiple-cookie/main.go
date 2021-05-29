package main

import (
	"fmt"
	"log"
	"net/http"
)
func main(){
	http.HandleFunc("/",foo)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.HandleFunc("/read",read)
	http.HandleFunc("/multi",multi)
	http.ListenAndServe(":8080",nil)
}
func foo(w http.ResponseWriter,req *http.Request){
	http.SetCookie(w,&http.Cookie{Name:"first",Value:"cookie1"})
	fmt.Fprintln(w,"setting one cookie")
}
func read(w http.ResponseWriter,req *http.Request){
	c1,err:=req.Cookie("first")
	if err != nil{
		log.Println(err)
	} else{
	fmt.Fprintln(w,"Your Cookie",c1)
	}
	c2,err:=req.Cookie("second")
	if err != nil{
		log.Println(err)
	} else{
	fmt.Fprintln(w,"Your Cookie",c2)
	}
	c3,err:=req.Cookie("third")
	if err != nil{
		log.Println(err)
	} else{
	fmt.Fprintln(w,"Your Cookie",c3)
	}
    
}
func multi(w http.ResponseWriter,req *http.Request){
	http.SetCookie(w,&http.Cookie{Name:"second",Value:"cookie2"})
	http.SetCookie(w,&http.Cookie{Name:"third",Value:"cookie3"})
}