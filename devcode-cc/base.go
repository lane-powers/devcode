package base

import (
    "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/test", testHandler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, from the devcode base!")
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "tester tester, break 10")
}
