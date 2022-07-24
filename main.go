package main

import "net/http"

type Page struct {
	Title string
}

func HomePage(w http.ResponseWriter, r *http.Request) {
		p:= Page{Title: "Generate QR Code"}
}