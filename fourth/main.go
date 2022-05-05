package main

import (
	"fmt"
	"html/template"
	"net/http"
	"text/template"
)
var tpl *template.Template
func init(){
	tpl=template.Must(template.ParseGlob("templates/*"))
}
func main(){
	http.HandleFunc("/",foo)
	http.HandleFunc("/bar",bar)
	http.HandleFunc("/barred",barred)
	http.Handle()

}
func foo(w http.ResponseWriter,req *http.Request){
  fmt.Print()
}

func bar(w http.ResponseWriter,req *http.Request){

}
func barred(){

}