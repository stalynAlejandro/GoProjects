/*
Go by Example: Context

In the previous example we looked at setting up a simple HTTP server. HTTP servers are useful
for demostrating of context.Context for controlling cancellation.

A Context carries deadlines, cancellation signals, and other request-scoped values across API
boundaries and goroutines.

A context.Context is created for each request by the net/http machinery, and is available with
the Context() method.

Wait for a few seconds before sending a reply to the client. This could simulate some work the
server is doing. While working, keep an eye on the context's Done() channel for a signal that we
sould cancel the work and return as soon as possible.

The conext's Err() method returns an error that explains why the Done() channel was closed.


As before, we register our handler on the '/hello' route, and start serving.
Run the server in the background.
Simulate a client request to /hello, hitting Ctr+C shortly after start to signal cancellation.
*/

package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("Server: hello handler started")
	defer fmt.Println("server: helllo handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello/n")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server: ", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
