package app

import (
	"bytes"
	"encoding/json"
	"github.com/danieljvx/meli-morse/requests"
	"github.com/danieljvx/meli-morse/responses"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestApp(t *testing.T) {
	app := App()

	testsDecodeBits := requests.TranslateRequest{
		Text:  "000000001101101100111000001111110001111110011111100000001110111111110111011100000001100011111100000111111001111110000000110000110111111110111011100000011011100000000000",
	}
	testsDecodeBitsResponse := ".... --- .-.. .- -- . .-.. .."

	tests2text := requests.TranslateRequest{
		Text:  ".... --- .-.. .-  -- . .-.. ..",
	}
	tests2textResponse := "HOLA MELI"

	tests2morse := requests.TranslateRequest{
		Text:  "HOLA MELI",
	}
	tests2morseResponse := ".... --- .-.. .-  -- . .-.. .."

	rbodyDecodeBits, _ := json.Marshal(testsDecodeBits)
	requestTestsDecodeBits := httptest.NewRequest("POST", "/translate/decodeBits", bytes.NewBuffer(rbodyDecodeBits))
	requestTestsDecodeBits.Header.Set("Content-Type", "application/json")
	responseDecodeBits, _ := app.Test(requestTestsDecodeBits)
	if responseDecodeBits.StatusCode != 201 {
		t.Errorf("translate/decodeBits StatusCode was incorrect, got: %d, want: %d.", responseDecodeBits.StatusCode, 201)
	}
	bodyDecodeBits, _ := ioutil.ReadAll(responseDecodeBits.Body)
	var decodeBitsResponse responses.TranslateResponse
	json.Unmarshal(bodyDecodeBits, &decodeBitsResponse)
	if decodeBitsResponse.Message != testsDecodeBitsResponse {
		t.Errorf("translate/decodeBits Body was incorrect, got: %v, want: %v", decodeBitsResponse.Message, testsDecodeBitsResponse)
	}

	rbody2text, _ := json.Marshal(tests2text)
	requestTests2text := httptest.NewRequest("POST", "/translate/2text", bytes.NewBuffer(rbody2text))
	requestTests2text.Header.Set("Content-Type", "application/json")
	response2text, _ := app.Test(requestTests2text)
	if response2text.StatusCode != 201 {
		t.Errorf("translate/2text StatusCode was incorrect, got: %d, want: %d.", response2text.StatusCode, 201)
	}
	body2text, _ := ioutil.ReadAll(response2text.Body)
	var actual2textResponse responses.TranslateResponse
	json.Unmarshal(body2text, &actual2textResponse)
	if actual2textResponse.Message != tests2textResponse {
		t.Errorf("translate/2text Body was incorrect, got: %v, want: %v", actual2textResponse.Message, tests2textResponse)
	}

	rbody2morse, _ := json.Marshal(tests2morse)
	requestTests2morse := httptest.NewRequest("POST", "/translate/2morse", bytes.NewBuffer(rbody2morse))
	requestTests2morse.Header.Set("Content-Type", "application/json")
	response2morse, _ := app.Test(requestTests2morse)
	if response2morse.StatusCode != 201 {
		t.Errorf("translate/2morse StatusCode was incorrect, got: %d, want: %d.", response2morse.StatusCode, 201)
	}
	body2morse, _ := ioutil.ReadAll(response2morse.Body)
	var actual2morseResponse responses.TranslateResponse
	json.Unmarshal(body2morse, &actual2morseResponse)
	if actual2morseResponse.Message != tests2morseResponse {
		t.Errorf("translate/2morse Body was incorrect, got: %v, want: %v", actual2morseResponse.Message, tests2morseResponse)
	}
}