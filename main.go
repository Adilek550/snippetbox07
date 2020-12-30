package main
import (
	"log"
	"net/http"
)
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Bro"))
}
func showBro(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a Bro..."))
}
func createBro(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new Bro..."))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hiBro", home)
	mux.HandleFunc("/Bro", showBro)
	mux.HandleFunc("/Bro/Brat", createBro)
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
