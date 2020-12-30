package main
import (
	"log"
	"net/http"
)
func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "This Method Not Allowed", 405)
		w.Header()["Date"] = nil
		return
	}
	w.Write([]byte("Create a new snippet, Bro..."))
}


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hiBro", createSnippet)
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

//$ curl -i -X PUT http://localhost:4000/hiBro
//HTTP/1.1 405 Method Not Allowed
//Allow: POST
//Date: Wed, 30 Dec 2020 07:48:07 GMT
//Content-Length: 18
//Content-Type: text/plain; charset=utf-8
//
//Method Not Allowed
