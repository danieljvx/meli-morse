package routes

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestTranslateRoute(t *testing.T) {
	// Create a sample data string.
	dataString := `{"text-wrong": ".... --- .-.. .-  -- . .-.. .."}`

	// Define a structure for specifying input and output data of a single test case.
	tests := []struct {
		description   string
		route         string // input route
		method        string // input method
		tokenString   string // input token
		body          io.Reader
		expectedError bool
		expectedCode  int
	}{
		{
			description:   "/translate/decodeBits without body",
			route:         "/translate/decodeBits",
			method:        "POST",
			body:          nil,
			expectedError: false,
			expectedCode:  400,
		},
		{
			description:   "/translate/2text with body wrong",
			route:         "/translate/2text",
			method:        "POST",
			body:          strings.NewReader(dataString),
			expectedError: false,
			expectedCode:  400,
		},
		{
			description:   "/translate/2morse with method wrong",
			route:         "/translate/2morse",
			method:        "PUT",
			body:          strings.NewReader(dataString),
			expectedError: false,
			expectedCode:  405,
		},
	}

	app := fiber.New()

	TranslateRoute(app)

	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.route, test.body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, -1)
		errorType := err != nil
		if test.expectedError != errorType {
			t.Errorf(test.description + ", got: %s, want: %s.", test.expectedError, err)
		}

		if test.expectedError {
			continue
		}

		if test.expectedCode != resp.StatusCode {
			t.Errorf(test.description + ", got: %s, want: %s.", test.expectedCode, resp.StatusCode)
		}

	}
}