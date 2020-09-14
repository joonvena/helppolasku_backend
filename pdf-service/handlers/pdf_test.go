package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestValidation(t *testing.T) {

}

func TestThatOutputIsCorrect(t *testing.T) {

	// var input Details

	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateDocument)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if output := rr.Header().Get("Content-Type"); output != "application/pdf" {
		t.Errorf("Doument type was %v instead of %v", output, "application/pdf")
	}
}
