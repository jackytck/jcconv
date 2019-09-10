package web

import (
	"fmt"
	"net/http"
)

// Index renders the index page with source represented as page.
func Index(page string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, page)
	}
}
