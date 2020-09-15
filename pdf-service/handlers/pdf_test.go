package handlers

import (
	"testing"

	"github.com/go-playground/validator"
)

func TestValidation(t *testing.T) {

	validate := validator.New()

	// Check all fields empty.
	emptyData := &Details{
		Name:    "",
		Address: "",
	}

	err := validate.Struct(emptyData)
	if err == nil {
		t.Fatal(err)
	}

	// Check field Name empty.
	nameMissing := &Details{
		Name:    "",
		Address: "Go Street",
	}

	err = validate.Struct(nameMissing)
	if err == nil {
		t.Fatal(err)
	}

	// Check field Address empty.
	addressMissing := &Details{
		Name:    "Test",
		Address: "",
	}

	err = validate.Struct(addressMissing)
	if err == nil {
		t.Fatal(err)
	}

	// All fields correct.
	correct := &Details{
		Name:    "Test",
		Address: "Go Street",
	}

	err = validate.Struct(correct)
	if err != nil {
		t.Fatal(err)
	}

}
