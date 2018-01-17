package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", sign)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func sign(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("https://www.goodreads.com/api/auth_user")
	if err != nil {
		fmt.Fprintf(w, "request: %s", err)
	}

	fmt.Fprint(w, resp)

	// 	c := appengine.NewContext(r)

	// 	http := urlfetch.Client(c)

	// 	resp, err := http.Get("https://www.goodreads.com/api/auth_user")
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}

	// 	http.Redirect(w, r, "/", http.StatusFound)
}
