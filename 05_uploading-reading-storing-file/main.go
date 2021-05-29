package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)
var tpl *template.Template
func init(){
	tpl=template.Must(template.ParseFiles("index.html"))
}
func main(){
	http.HandleFunc("/",foo)
	http.Handle("/favicon.ico", http.NotFoundHandler()) 
	http.ListenAndServe(":8080",nil)
}
func foo(w http.ResponseWriter,req *http.Request){
	var s string
	fmt.Println(req.Method)
	if req.Method == "POST"{
		f,k,err:=req.FormFile("upload")
		if err != nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}
		defer f.Close()
		fmt.Printf("file name- %v / size = %v\n",k.Filename,k.Size)
		fmt.Println("next info test,printing f")
		fmt.Println(f)
		fmt.Println("printing k")
		fmt.Println(k)
		r,err:=ioutil.ReadAll(f)
		if err != nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}
		s=string(r)
		//storing file on server
		n,err := os.Create(filepath.Join("./user/",k.Filename))
		if err != nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}
		defer n.Close()
		_,err=n.Write(r)
		if err !=nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}
	}
	err:=tpl.ExecuteTemplate(w,"index.html",s)
	if err != nil{
		http.Error(w,err.Error(),404)
	}
}