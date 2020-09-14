package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jung-kurt/gofpdf"
)

// Details defines...
type Details struct {
	Name    string `validate:"required"`
	Address string `validate:"required"`
}

// CreateDocument will create the pdf output from the bill
func CreateDocument(rw http.ResponseWriter, r *http.Request) {
	var d Details
	if r.Body != http.NoBody {
		err := json.NewDecoder(r.Body).Decode(&d)
		if err != nil {
			http.Error(rw, "Can't decode to JSON"+err.Error(), http.StatusBadRequest)
		}
	}

	//validate := validator.New()
	//err := validate.Struct(d)
	//if err != nil {
	//	http.Error(rw, "Validation failed"+err.Error(), http.StatusBadRequest)
	//	return
	//}

	code, err := http.Get("http://localhost:9001")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := ioutil.ReadAll(code.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	log.Print(code)

	rw.Header().Set("Content-Type", "application/pdf")

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	// pdf.CellFormat(50, 10, "Hello", "1", 0, "L", false, 0, "")
	pdf.RegisterImageOptionsReader("barcode", gofpdf.ImageOptions{ImageType: "jpeg", ReadDpi: true}, bytes.NewReader(b))
	pdf.ImageOptions("barcode", 10, 20, 95, 12.7, false, gofpdf.ImageOptions{ImageType: "jpeg", ReadDpi: true}, 0, "")
	pdf.Ln(-1)
	err = pdf.Output(rw)

	if err != nil {
		http.Error(rw, "Could not create pdf file "+err.Error(), http.StatusBadRequest)
		return
	}

}
