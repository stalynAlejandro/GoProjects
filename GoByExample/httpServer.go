/*
Go by Example: HTTP Servers

Writing a basic HTTP server is easy using the net/http package.

A fundamental concept in net/http server is handlers.

A handler is an object implementing the http.Handler interface. A common wat to write
a handler is by using the http.HandlerFunc adapter on functions with the appropiate signature.

Functions serving as handlers take a http.ResponseWritter and a http.Request as arguments.
The response writer is used to fill in the HTTP response. Here our multiple response is just 'hello\n'

This handler does something a little more sophisticated by reading all the HTTP request headers and
ecoing them into the response body.

We register our handlers on server routes using the http.HandleFunc convenience function. It sets up
the default router in the net/http package and takes a function as an argument.

Finally we call the ListenAndServe with the port and a handler.

nil tells it to use the default router we've just set up.

Run the server in the background.
Access the /hello route.
*/

package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello \n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v \n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}
