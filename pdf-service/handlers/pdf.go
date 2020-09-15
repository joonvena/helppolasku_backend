package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/jung-kurt/gofpdf"
)

// Details defines...
type Details struct {
	Name    string `validate:"required"`
	Address string `validate:"required"`
}

func CreateBarcode(rw http.ResponseWriter) (picture []byte, err error) {
	code, err := http.Get("http://localhost:9001")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	picture, err = ioutil.ReadAll(code.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	return
}

// CreateDocument will create the pdf output from the bill
func CreateDocument(rw http.ResponseWriter, r *http.Request) {
	var d Details

	if r.Body != nil {
		err := json.NewDecoder(r.Body).Decode(&d)
		if err != nil {
			http.Error(rw, "Can't decode to JSON"+err.Error(), http.StatusBadRequest)
			return
		}
	}

	validate := validator.New()
	err := validate.Struct(d)
	if err != nil {
		http.Error(rw, "Validation failed"+err.Error(), http.StatusBadRequest)
		return
	}

	code, err := CreateBarcode(rw)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}

	rw.Header().Set("Content-Type", "application/pdf")
	te := []string{"Selite", "Määrä", "Hinta"}
	info := []string{"Konsultointi", "2", "3000 €"}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	enc := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.CellFormat(25, 10, enc("Matti Meikäläinen"), "0", 1, "", false, 0, "")
	pdf.CellFormat(25, 10, "Testikatu 6 D29", "0", 2, "", false, 0, "")
	pdf.CellFormat(25, 10, "70200, Kuopio", "0", 3, "", false, 0, "")
	pdf.Ln(-1)
	for _, h := range te {
		pdf.CellFormat(63, 7, h, "B", 0, "", false, 0, "")
	}
	pdf.Ln(-1)
	for _, i := range info {
		pdf.CellFormat(63, 7, i, "0", 0, "", false, 0, "")
	}
	pdf.RegisterImageOptionsReader("barcode", gofpdf.ImageOptions{ImageType: "jpeg", ReadDpi: true}, bytes.NewReader(code))
	pdf.ImageOptions("barcode", 10, 150, 95, 12.7, false, gofpdf.ImageOptions{ImageType: "jpeg", ReadDpi: true}, 0, "")

	err = pdf.Output(rw)
	if err != nil {
		http.Error(rw, "Could not create pdf file "+err.Error(), http.StatusBadRequest)
		return
	}

}
