package main
import (
	"log"
	"net/http"
)
func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}
	w.Write([]byte("Create a new snippet..."))
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
//Date: Wed, 30 Dec 2020 07:44:25 GMT
//Content-Length: 18
//Content-Type: text/plain; charset=utf-8
//
//Method Not Allowed
