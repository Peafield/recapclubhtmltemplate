package main

import (
	"html/template"
	"net/http"
)

// 1# Create a var of type template
var tmpl *template.Template

// 2# Use func init() and the code below to intialise all html files in static (*.html means all html files)
func init() {
	tmpl = template.Must(template.ParseGlob("static/*.html"))
}

func main() {
	// 3# Use the code below to pass over all your files in static (html, css, js) to the server
	path := "static"
	fs := http.FileServer(http.Dir(path))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// 4# 'Handle' the GET and POST requests from the browser with HandleFunc which then executes a function.
	http.HandleFunc("/", index)
	http.HandleFunc("/range", ranging)
	// 5# Serve up the server using ListenAndServe chosing the port number you'd like it to appear on (':8080' seems to be standard)
	http.ListenAndServe(":8080", nil)
}

// 6# Create functions to Write (w) and Read (r) from the browser.
func index(w http.ResponseWriter, r *http.Request) {
	data := "Hello World!"
	// 7# You can pass over data into html using the code below (check out index.html for the code you need there)
	tmpl.ExecuteTemplate(w, "index.html", data)
}

// 8# You can range over slices too! (check out range.html for the code you need there)
func ranging(w http.ResponseWriter, r *http.Request) {
	data := []int{1, 2, 3, 4}
	tmpl.ExecuteTemplate(w, "range.html", data)
}

// 9# To start the server, do 'go run main.go' and go to http://localhost:8080/ on your browser. The first page you see will be index.html
// go to http://localhost:8080/range to see the range.html
