package handlers

import (
	"fmt"
	"net/http"
)

// CreateBarcode will create barcode and add it to pdf document.
func CreateBarcode(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	fmt.Println("Hello")
}
