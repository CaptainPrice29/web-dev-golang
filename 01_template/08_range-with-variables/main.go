package main

import (
	"log"
	"os"
	"text/template"
)
type tpl *template.Template
func main(){
	b :=[]int{11,23,34,45,56,78}
	tpl := template.Must(template.ParseFiles("tpl.gohtml"))
	err := tpl.Execute(os.Stdout,b)
	if err != nil{
		log.Fatalln(err)
	}
}