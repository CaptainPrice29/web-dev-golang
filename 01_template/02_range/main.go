package main

import (
	"os"
	"text/template"
)
var temp *template.Template
type info struct{
	Name []string
	Age   int
	Weight float64
}
func main(){
	x,err := os.Create("test.gohtml")
	if err != nil{
		panic(err)
	}
	temp = 	template.Must(template.ParseFiles("tpl.gohtml"))
	person1 := info{
		Name: []string{"CaptainPrice"},
		Age: 65,
		Weight: 70.65,
	}
	person2 := info{
		Name: []string{"Rust"},
		Age: 55,
		Weight: 80.21,
	}
	person3 := info{
		Name: []string{"CaptainPrice"},
		Age: 58,
		Weight: 69.75,
	}
	inf := []info{person1,person2,person3}
	err = temp.ExecuteTemplate(x,"tpl.gohtml",inf)
	if err != nil{
		panic(err)
	}

}