package handlers

import (
	"net/http"

	"github.com/jung-kurt/gofpdf"
)

// CreateDocument will create the pdf output from the bill
func CreateDocument(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "application/pdf")

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.CellFormat(50, 10, "Hello", "1", 0, "L", false, 0, "")
	pdf.Ln(-1)
	err := pdf.Output(rw)

	if err != nil {
		http.Error(rw, "Could not create pdf file"+err.Error(), http.StatusBadRequest)
		return
	}

}
