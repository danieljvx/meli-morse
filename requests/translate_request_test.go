package requests

import (
"testing"
)

func TestTranslateRequest(t *testing.T) {
	request := TranslateRequest{
		Text: "hello meli",
	}
	if request.Text != "hello meli" {
		t.Errorf("TestTranslateRequest was incorrect, got: %s, want: %s.", request.Text, "hello meli")
	}
}
