package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOutputType(t *testing.T) {

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateBarcode)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code got %v expected %v", rr.Code, http.StatusOK)
	}

	if output := rr.Header().Get("Content-Type"); output != "image/jpeg" {
		t.Errorf("Expected Content-Type %v got %v", output, "image/jpeg")
	}

}
