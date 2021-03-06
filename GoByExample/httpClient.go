/*
The Go standard library comes wiht excellent support for HTTP clients and servers in the
net/http package. In this example we'll use it to issue simple HTTP requests.

Issue an HTTP GET request to a server.  http.Get is a convenient shortcut around
creating an http.Client and calling its Get method; it uses the http.DefaultClient object
which has useful default settings.

Print the HTTP resposne status.
Print the first 5 lines of the response body.
*/

package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://gobyexample.com")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Print("Response status: ", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
