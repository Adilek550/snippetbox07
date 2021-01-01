	package main

	import (
		"fmt"
		"html/template"
		"log"
		"net/http"
		"strconv"
	)
	func downloadHandler(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./ui/static/file.zip")
	}
	func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
	http.NotFound(w, r)
	return
	}
	files := []string{
			"C:\\Users\\User\\go\\src\\snippetbox\\ui\\html\\html.page.tmpl",
			"C:\\Users\\User\\go\\src\\snippetbox\\ui\\html\\base.layout.tmpl",
			"C:\\Users\\User\\go\\src\\snippetbox\\ui\\html\\footer.partial.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
	log.Println(err.Error())
	http.Error(w, "Internal Server Error", 500)
	return
	}
	err = ts.Execute(w, nil)
	if err != nil {
	log.Println(err.Error())
	http.Error(w, "Internal Server Error", 500)
	}
	}
	func showSnippet(w http.ResponseWriter, r *http.Request) {
			id, err := strconv.Atoi(r.URL.Query().Get("id"))
			if err != nil || id < 1 {
				http.NotFound(w, r)
				return
			}
			fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
		}
		func createSnippet(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				w.Header().Set("Allow", http.MethodPost)
				http.Error(w, "Method Not Allowed", 405)
				return
			}
			w.Write([]byte("Create a new snippet..."))
		}
