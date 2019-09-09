package web

import (
	"fmt"
	"net/http"
)

// HelloServer prints hello message.
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
