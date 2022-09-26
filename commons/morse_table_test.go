package commons

import (
"testing"
)

func TestAlphaNumToMorse(t *testing.T) {
	for k, v := range alphaNumToMorse {
		if k != morseToAlphaNum[v] {
			t.Errorf("alphaNumToMorse == morseToAlphaNum was incorrect, got: %s, want: %s.", k, morseToAlphaNum[v])
		}
	}
}
