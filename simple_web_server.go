package main

import(
	"fmt"
	"log"
	"net/http"
)
//The function handler is of the type http.HandlerFunc. 
//It takes an http.ResponseWriter and an http.Request as its arguments.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main(){
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
//ListenAndServe always returns an error, since it only returns when an unexpected error occurs. 
//In order to log that error we wrap the function call with log.Fatal

