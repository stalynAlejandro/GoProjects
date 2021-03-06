/*
Writing Web Applications
- Creating a data structure with load and save methods
- Using the net/http package to build web applications
- Using html/template package to process HTML templates
- Using the regexp package to validate user input
- Using closures

Lets create a Wiki. A Wiki consists of a series of interconnected pages, each of which
has a title and a body (the page content).
*/

package main

import (
	"fmt"
	"io/ioutil"
)

// Data Structure
// 		The type []byte means "a byte slice". The body element is a []byte rather than string because that
// 		is the expected by the io libraries we will use, as you'll see below.
type Page struct {
	Title string
	Body  []byte
}

// Save Method.
// 		The Page struct describes how page data will be stored in memory. But what about persisent storage? We can
// 		address that by creating a save method on Page.

// 		This method's signature reads: 'This is a method named save that takes as its receiver p, a pointer to Page. It
// 		takes no parameters, and returns a value of type error.'

// 		This method will save the Page's body to a text file. For simplicity we will use the Title as the file name.
// 		The save method returns an error value because that is the return type of WriteFile. The save method returns the error
// 		value, to let the application handle it should anything go wrong while writing the file. If all goes well, Page.save()
// 		will return nil.

// 		The octal integer literal 0600, indicates that the file should be created with read-write permissions for the current
// 		user only.
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// Load Method.
// 		The function loadPage constructs the file name from the tilte parameter, reads the file's content into a new variable body,
// 		and returns a pointer to a Page literal constructed with the proper title and body values.

// 		Function can return multiple values. The standard library function io.ReadFile returns []byte and error.
// 		If ioutil.ReadFile encounters an error we have to handle it.

// 		Callers of this function can now new check the second parameter, if its nil then it has a successfully loaded a Page. If not
// 		it will be an error that can be handled by the caller.
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a simple Page.")}
	p1.save()

	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
