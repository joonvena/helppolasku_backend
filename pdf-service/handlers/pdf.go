package handlers

import (
	"fmt"
	"net/http"
)

// CreateDocument will create the pdf output from the bill
func CreateDocument(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	fmt.Println("Hello")
}
