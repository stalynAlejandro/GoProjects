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
	"errors"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

// Validation
// 		Security Flaw: A user can supply an arbitrary path to be read/written on the server. To mitigate this, we can write a function
// 		to validate the title with regular expression.

// 		The function regexp.MustCompile will parse and compile the regular expression, and return a regexp.Regexp.MustCompile is
// 		distinct from Compile in that it will panic if the expression compilation fails, when Compile returns an error as a second
// 		parameter.
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// 		Function that uses validPath expression to validate path and extract the page title
// 		If the title is valid, it will be returned along with a nil error value. It the title is invalid, the function will write a "404"
// 		"Not found error" and return an error to the handler.
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}
	return m[2], nil //The title is the second subexpression
}

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

//		Create a global variable named templates, and initialize it with ParseFiles.
//		The function template.Must is a convenience wrapper that panics when passed a non nil-error value, and otherwise returns the
//		*Template unaltered. A panic is appropiate here: if the templates can't be loaded the only sensible thing to do is exit.
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

// 		We've used almost exactly the same templating code in both handlers. Let's remove this duplication by moving the templating code
// 		to its own function.

// 		The http.Error function sends a specified HTTP response code and error message.
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Handler viewHandler.
// 		Will allow users to view a wiki page. It will handle URLS prefixed with "/view/".
// 		Note the use of _ to ignore the error return value from loadPage. This is done here for simplicity and generally considered
// 		bad practice.

// 		First, this function extracts the page title from r.URL.Path, the path component of the request URL. The path is re-sliced with
// 		[len("/view/"):] to drop the leading "/view/" component of the request path. This is because the path will invariably begin with
// 		"/view/", which is not part of the page's title.

// 		The function the loads the page data, formats the page with a string of simple HTML, and writes it to w, the http.ResponseWriter
//		example: http://localhost:8080/view/test 	->		loadPage(text)

// 		Handling non-existing pages
// 		If the requested page doesn't exist, it should redirect the client to the editPage so the content may be created

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusNotFound)
		return
	}
	renderTemplate(w, "view", p)
}

// Handler editHandler.
// 		Display and edit page form.
// 		The html/template package is part of the Go standard library. We can use html/tmplate to keep the HTML in a separated file,
// 		allowing us to change the layout of our edit page without modifying the underlying Go code.

// 		The function template.ParseFiles will read the content of edit.html and return a *template.Template
// 		The method t.Execute executes the template, writing the generated HTML to the http.ResponseWrite. The .Title and .Body dotted
// 		identifiers refer to p.Title and p.Body
func editHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

// Handler saveHandler.
// 		Save the data entered via the form.
// 		Will handle the submission of forms located on the edit pages. After uncommenting the related in main.

// 		The page title (provided in the URL) and the form's only field, Body are stored in a new page. The save() method is
// 		the called to write the data to a file, and the client is redirected to the /view/ page.

// 		The value returned by FormValue is of type string. We must convert that value to []byte before it will fit into the
// 		Page struct. We use []byte(body) to perform the conversion.

//		Any errors that occur during p.save() will be reported to the user.
func saveHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}

	err = p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
