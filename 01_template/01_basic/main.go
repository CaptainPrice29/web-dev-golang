package main

import (
	"fmt"
	"os"
	"text/template"
)
var temp *template.Template
type info struct{
	Name string
	Age   int
	Weight float64
}
func main(){
	//template.Must takes *Template and error then checks for error and return *Template
	temp = 	template.Must(template.ParseFiles("tpl.gohtml"))
	inf := info{
		Name: "Captain",
		Age: 55,
		Weight: 70.65,
	}
	err := temp.ExecuteTemplate(os.Stdout,"tpl.gohtml",inf)
	fmt.Println(err)

}