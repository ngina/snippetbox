package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if i, err := strconv.Atoi(id); err != nil || i < 0 {
		log.Printf("unacceptable query parameter: [id=%#v]", id)
		http.NotFound(w, r)
		return
	} else {
		fmt.Fprintf(w, "Display a specific snippet with ID %d...", i)
		return
	}
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {

	//update our application so that the /snippet/create route only responds to
	// HTTP requests which use the POST method, like so:
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

func main() {

	// Use the http.NewServeMux() function to initialize a new servemux
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
