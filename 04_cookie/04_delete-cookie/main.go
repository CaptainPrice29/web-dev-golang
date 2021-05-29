package main

import (
	"fmt"
	"net/http"
)
func main(){
	http.HandleFunc("/",foo)
	http.HandleFunc("/set",set)
	http.HandleFunc("/read",read)
	http.HandleFunc("/delete",delete)
	http.ListenAndServe(":8080",nil)
}
func foo(w http.ResponseWriter,req *http.Request){
	fmt.Fprint(w,`<h1><a href="/set">SET</a></h1>`)
}
func set(w http.ResponseWriter,req *http.Request){
	http.SetCookie(w,&http.Cookie{Name:"captnCookie",Value:"some unique value"})
	fmt.Fprintln(w,`<h1><a href="/read">READ</a></h1>`)
	// fmt.Fprint(w,"cookie written")
}
func read(w http.ResponseWriter,req *http.Request){
	c,err:=req.Cookie("captnCookie")
	if err != nil{
		http.Error(w,err.Error(),http.StatusNotFound)
	}
	fmt.Fprintln(w,`<h1><a href="/delete">DELETE</a></h1>`)
	fmt.Fprintln(w,"cookie=>",c)
}
func delete(w http.ResponseWriter,req *http.Request){
	c,err := req.Cookie("captnCookie")
	if err != nil{
		http.Error(w,err.Error(),http.StatusNotFound)
	}
	c.MaxAge=-1
	http.SetCookie(w,c)
	fmt.Fprint(w,"Cookie Deleted")

}