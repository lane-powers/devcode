package main
// package base

import (
  "fmt"
	"net/http"
	"io/ioutil"
)

type Page struct {
	Title string
	Body []byte
}

const lenPath = len("/view/")

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/test", testHandler)
	http.HandleFunc("/test/", testHandler)
	http.HandleFunc("/view/", restViewHandler)
}

// taken verbatim from golang.org/doc/articiles/wiki
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
        	return nil, err
	}
	fmt.Println("Title: ", title, "\nBody: ", body)
	return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, from the devcode base!")
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "tester tester, break 10")
}

func restViewHandler(w http.ResponseWriter, r *http.Request) {
	Title := r.URL.Path[lenPath:]
	p, _ := loadPage(Title)
	if p != nil {
		fmt.Println(p)
		fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	} else {
		fmt.Fprintf(w, "<h1>NO TEMPLATE FILE</h1>")
	}
}

// if standalone, uncomment main
func main() {
	http.ListenAndServe(":8080", nil)
}
