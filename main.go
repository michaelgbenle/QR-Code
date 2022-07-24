package main

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title string
}

func HomePage(w http.ResponseWriter, r *http.Request) {
		p:= Page{Title: "Generate QR Code"}
		t,_:= template.ParseFiles("generator.html")
		t.Execute(w,p)
}
func CodePage(w http.ResponseWriter, r *http.Request) {

}