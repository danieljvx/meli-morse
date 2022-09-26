package responses

import (
"testing"
)

func TestTranslateResponse(t *testing.T) {
	response := TranslateResponse{
		Message: "hello meli",
		Status: 201,
	}
	if response.Message != "hello meli" {
		t.Errorf("TestTranslateResponse was incorrect, got: %s, want: %s.", response.Message, "hello meli")
	}
	if response.Status != 201 {
		t.Errorf("TestTranslateResponse was incorrect, got: %d, want: %d.", response.Status, 201)
	}
}
