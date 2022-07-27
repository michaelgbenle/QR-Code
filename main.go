package main

import (
	"html/template"
	"image/png"
	"log"
	"net/http"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

type Page struct {
	Title string
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	p := Page{Title: "Generate QR Code"}
	t, _ := template.ParseFiles("generator.html")
	err := t.Execute(w, p)
	if err != nil {
		log.Fatal(err)
	}
}
func CodePage(w http.ResponseWriter, r *http.Request) {
	message := r.FormValue("message")
	qrCode, err := qr.Encode(message, qr.L, qr.Auto)
	if err != nil {
		log.Fatal(err)
	}
	qrCode, err = barcode.Scale(qrCode, 600, 600)
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(w, qrCode)
	if err != nil {
		log.Fatal(err)
	}

}
func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/code", CodePage)
	err := http.ListenAndServe(":5050", nil)
	if err != nil {
		log.Fatal(err)
	}
}
