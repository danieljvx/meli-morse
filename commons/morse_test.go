package commons

import (
	"strings"
	"testing"
)

func TestNewTranslate(t *testing.T) {
	translate := NewTranslate()

	decodeBits, _ := translate.DecodeBits2Morse(strings.NewReader("000000001101101100111000001111110001111110011111100000001110111111110111011100000001100011111100000111111001111110000000110000110111111110111011100000011011100000000000"))
	decodeBitsReceived := ".... --- .-.. .- -- . .-.. .."
	if string(decodeBits) != decodeBitsReceived {
		t.Errorf("DecodeBits2Morse was incorrect, got: %s, want: %s.", string(decodeBits), decodeBitsReceived)
	}

	toHuman, _ := translate.Translate2Human(strings.NewReader(".... --- .-.. .-  -- . .-.. .."))
	toHumanReceived := "HOLA MELI"
	if string(toHuman) != toHumanReceived {
		t.Errorf("Translate2Human was incorrect, got: %s, want: %s.", string(toHuman), toHumanReceived)
	}

	toMorse, _ := translate.Translate2Morse(strings.NewReader("HOLA MELI"))
	toMorseReceived := ".... --- .-.. .-  -- . .-.. .."
	if string(toMorse) != toMorseReceived {
		t.Errorf("Translate2Morse was incorrect, got: %s, want: %s.", string(toMorse), toMorseReceived)
	}
}