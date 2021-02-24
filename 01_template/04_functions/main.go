package main

import (
	"log"
	"os"
	"text/template"
)
var tpl *template.Template

var fn = template.FuncMap{
	"ad" : add,
	"mul" : multiply,
}
type numbers struct{
	Num1 int
	Num2 int
}
func add(n numbers)int{
	return n.Num1+n.Num2
}
func multiply(n numbers)int{
	 return n.Num1*n.Num2
}
func main(){
	n := numbers{5,3}
	tpl = template.Must(template.New("").Funcs(fn).ParseFiles("tpl.gohtml"))
	err := tpl.ExecuteTemplate(os.Stdout,"tpl.gohtml",n)
	if err != nil{
		log.Fatalln(err)
	}
}