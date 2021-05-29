package main

import (
	"fmt"
	"net/http"
	"strconv"
)
func main(){
	http.HandleFunc("/",foo)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}
func foo(w http.ResponseWriter,req *http.Request){
	c,err:=req.Cookie("testcookie")
	if err == http.ErrNoCookie{
		c = &http.Cookie{Name: "testcookie",Value:"0"}
    }
	count,_:=strconv.Atoi(c.Value)
	count++
	c.Value = strconv.Itoa(count)
	http.SetCookie(w,c)
	fmt.Fprintln(w,c.Value)
}