package main

import (
	"fmt"
	"net/http"
	"text/template"
)


func renderTemp(w http.ResponseWriter, tmpl string) {
	parsedTemp, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemp.Execute(w, nil)
	if err != nil {
		fmt.Println("error occured", err)
		return
	}
}