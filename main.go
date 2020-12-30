package main
import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)
func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID: %d...", id)
}



func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hiBro", showSnippet)
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
