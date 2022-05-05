package main

import (
	// "fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init(){
	tpl=template.Must(template.ParseGlob("*"))
}

type Person struct{
	Fname string
	Lname string
	Subcribed bool

}




func main(){
	http.HandleFunc("/",abc)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8081",nil)
}

func abc(w http.ResponseWriter,req *http.Request){
	f:=req.FormValue("first")
	l:=req.FormValue("last")
	s:=req.FormValue("subscribe")=="on"

	err:=tpl.ExecuteTemplate(w,"index.html",Person{f,l,s})
	if err!=nil{
		http.Error(w,err.Error(),500)
		log.Fatal(err)
	}
}

