package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request){
	if (r.URL.Path != "/"){
		http.NotFound(w, r)
		return
	}

	files := []string {
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/home.html",
	}

	ts, err := template.ParseFiles(files...)

	if(err != nil){
		log.Println(err.Error())
		http.Error(w, "internal server error", 500)
		return
	}
	
	err = ts.ExecuteTemplate(w, "base",nil)

	if(err != nil){
		log.Println(err.Error())
		http.Error(w, "internal server error", 500)
		return
	}
}

func snippetView(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if (err != nil || id < 1) {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request){
	if (r.Method != "POST"){
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("create a new snippet"))
}

