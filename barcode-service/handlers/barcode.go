package handlers

import (
	"image/jpeg"
	"net/http"
	"os"

	"github.com/boombuler/barcode/code128"
)

// CreateBarcode will create barcode and add it to pdf document.
func CreateBarcode(rw http.ResponseWriter, r *http.Request) {

	code, err := code128.Encode("458598029510141230001550000000123456789056433635201212")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-Type", "image/jpeg")

	f, _ := os.Create("code.png")
	defer f.Close()

	jpeg.Encode(rw, code, nil)

}
