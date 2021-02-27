package main

import (
	"log"
	"os"
	"text/template"
)
var fn = template.FuncMap{ 
	"sl1" : slice1,
	"sl2" : slice2,
	"sl3" : slice3,

}
func slice1(s []int)[]int{
	return s[1:]
}
func slice2(s []int)[]int{
	return s[1:]
}
func slice3(s []int)[]int{ 
	return s[1:]
}
var tpl *template.Template
var tpl1 *template.Template

func main() {
	s :=[]int{1,2,3,4,5,6,7,8}
	tpl = template.Must(template.New("").Funcs(fn).ParseGlob("template/*.gohtml"))
	err := tpl.ExecuteTemplate(os.Stdout,"tpl.gohtml",s)
	if err != nil{
		log.Fatalln(err)
	}
	
}
