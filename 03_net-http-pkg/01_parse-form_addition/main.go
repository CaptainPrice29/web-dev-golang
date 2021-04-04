package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	
)
var tpl *template.Template
func init(){
	tpl=template.Must(template.ParseFiles("index.html","k.html"))
}
func sub(w http.ResponseWriter,req *http.Request){
	err:=tpl.ExecuteTemplate(w,"index.html",nil)
	if err != nil{
		log.Fatal(err)
	}
}
func res(w http.ResponseWriter,req *http.Request){
	err := req.ParseForm()
	if err !=nil{
		log.Fatalln(err)
	}
	f,_ := strconv.ParseFloat(req.FormValue("fnum"),64)
	s,_ := strconv.ParseFloat(req.FormValue("snum"),64)
	sum := f+s
	fmt.Println(req.Form)
	err=tpl.ExecuteTemplate(w,"k.html",sum)
	if err != nil{
		log.Fatal(err)
	}
}
func main(){
	http.HandleFunc("/",sub)
	http.HandleFunc("/k.html",res)
	log.Fatal(http.ListenAndServe(":8080",nil))
}