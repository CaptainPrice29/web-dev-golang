package main

import (
	"log"
	"os"
	"text/template"
)
type bike struct{
	Grpst string
	Prpse string
	Ctgry string
	Wght int
}
func(b bike) Groupset()string{
	return b.Grpst
}
func (b bike) Purpose()string{
	return b.Prpse
}
func(b bike) Category()string{
	return b.Ctgry
}
func(b bike) Weight()int{
	return b.Wght
}
func(b bike) Heavy(w int)string{
	if w < 10{
		return "light bike"
	}
		return "too heavy"
}
type tpl *template.Template
func main(){
	b := bike{
		Grpst: "Shimano SORA",
		Prpse: "Recreational",
		Ctgry: "Road/Gravel",
		Wght: 11,
	}
	tpl := template.Must(template.ParseFiles("tp1.gohtml"))
	err := tpl.Execute(os.Stdout,b)
	if err != nil{
		log.Fatalln(err)
	}
}