package main

import (
	"fmt"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)
func main(){
	http.HandleFunc("/",foo)
	http.HandleFunc("/set",set)
	http.HandleFunc("/read",read)
	http.ListenAndServe(":8080",nil)
}
func foo(w http.ResponseWriter,req *http.Request){
	fmt.Fprintln(w,`<h1><a href="/set">set cookie</a></h1>`)
}
func set(w http.ResponseWriter,req *http.Request){
	c := uuid.NewV4()
	id :=c.String()
	http.SetCookie(w,&http.Cookie{Name:"yourcookie",Value:id})
	fmt.Fprintln(w,`<h1><a href="/read">Read Cookie</a></h1>`)

}
func read(w http.ResponseWriter,req *http.Request){
	c,err:=req.Cookie("yourcookie")
	if err != nil{
		log.Println(err)
	}
	fmt.Fprintln(w,c)
}