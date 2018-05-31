/*

This is about router. why gorilla/mux ?? 


*/

/* without gorilla mux*/
package main 

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func handlerFunc(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/"{
		fmt.Fprintf(w, "<h1>Welcome to home page</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprintf(w, "<h2>contact page </h2>")
	} else {
		w.WriterHeader(http.StatusNotFound)
		fmt.Fprintf(w, "<h1>Not Found</h1>")
	}
}

func main(){
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000",nil)
}

/*----------------------------------------------*/

/* with gorilla mux */

/* import gorilla mux */

func home(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>welcome to home page</h1>")
}

func contact(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h2>contact page </h2>")
}

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)
}

/* Additional example on gorilla Mux */
/*
The Go's net/http package provides a lot of functionalitites
for the HTTP protocol. one thing it doesnt do very well is complex
request routing like segmenting a request url into single parameters.
*/
steps:
1. create a router of mux
2. call HandleFunc on router , to register the request handlers
//Biggest strength of gorilla/mux is to extract the params from router
suppose the URL is :
/books/go-programming-book/page/10
The above URL has two dynamic segments 
--Book title slug
--page

To have a request handler match the URL mentioned 
above , replace the dynamic segments of with placeholders
in your URL pattern 

r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request){
	//get the book 
	//navigate to page etc
	})


the last thing is to get the data from these segments.
The package comes with the function mux.Vars(r)
which takes the http.Request as parameter and return a map
of segments.

func(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	vars["title"] //the title of the book
	vars["page"] //the page 
}



Methods::

Restrict the request handler to specific HTTP methods.

r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

Hostnames & Subdomains
Restrict the request handler to specific hostnames or subdomains.

r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")


Schemes
Restrict the request handler to http/https.

r.HandleFunc("/secure", SecureHandler).Schemes("https")
r.HandleFunc("/insecure", InsecureHandler).Schemes("http")




Path Prefixes & Subrouters
Restrict the request handler to specific path prefixes.

bookrouter := r.PathPrefix("/books").Subrouter()
bookrouter.HandleFunc("/", AllBooks)
bookrouter.HandleFunc("/{title}", GetBook)


