package main

import (
	"fmt"
	"log"
	"net/http"
)

// https://pkg.go.dev/net/http
// ENABLING DEBUG - From the above, utilise as environment variables:
/*
GODEBUG=http2client=0  # disable HTTP/2 client support
GODEBUG=http2server=0  # disable HTTP/2 server support
GODEBUG=http2debug=1   # enable verbose HTTP/2 debug logs
GODEBUG=http2debug=2   # ... even more verbose, with frame dumps
*/

func main() {
	fmt.Println("Access web server at https://localhost:1234")
	h := http.FileServer(http.Dir("assets"))
	// Handlers. Root handler is the localhost:portnum URL, other is accessible at the /hello path in front of the URL
	// root directory of the web page:
	http.Handle("/", h)
	// /hello directory of web page:
	http.HandleFunc("/hello", SayHello)
	// Create keys with:
	// openssl req -newkey rsa:2048 -nodes -keyout server.key -x509 -days 365 -out server.crt
	log.Fatal(http.ListenAndServeTLS(":1234", "server.crt", "server.key", nil))
}

// Pointers * and & - https://www.golang-book.com/books/intro/8 - Essentially we assign the value of 'http.request' to 'request_var' within the function,
// So we effectively print (http.Request).Proto:
func SayHello(w http.ResponseWriter, request_var *http.Request) {
	log.Println("Hello, my protocol is: ", request_var.Proto)
	fmt.Fprintf(w, "Hello, I'm the new protocol H2\n")
}
