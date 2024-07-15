package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"ascii-web-complete/ascii"
)

func main() {
	http.HandleFunc("/", Handler)
	fmt.Println("Server open on http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "ERROR PARSING THE FORM", http.StatusBadRequest)
		return
	}

	type Data struct {
		Result string
		Color  string
	}

	var result strings.Builder
	word := r.Form.Get("Word")
	files := r.Form.Get("Files")
	color := r.Form.Get("Colors")

	if strings.Contains(word, "\n") {
		slice := strings.Split(word, "\r\n")
		for _, char := range slice {
			result.WriteString(Inputs(w, []string{char, files}))
		}
	} else {
		result.WriteString(Inputs(w, []string{word, files}))
	}

	data := Data{
		Result: result.String(),
		Color:  color,
	}

	temp, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "ERROR WHILE PARSING THE FILES", http.StatusInternalServerError)
		return
	}

	err = temp.Execute(w, data)
	if err != nil {
		http.Error(w, "ERROR EXECUTING THE PROGRAM", http.StatusInternalServerError)
	}
}

func Inputs(w http.ResponseWriter, input []string) string {
	filenames := input[1]
	return ascii.Ascii(w, input[0], filenames)
}
