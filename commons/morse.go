package commons

import (
	"io"
	"io/ioutil"
	"strings"
)

type BitToMorse interface {
	DecodeBits2Morse(io.Reader) ([]byte, error)
}

type ToMorse interface {
	Translate2Morse(io.Reader) ([]byte, error)
}

type ToHuman interface {
	Translate2Human(io.Reader) ([]byte, error)
}

type Translate interface {
	BitToMorse
	ToMorse
	ToHuman
}

type translate struct {
}

func (t *translate) DecodeBits2Morse(r io.Reader) ([]byte, error) {
	d, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	breaks := 0
	dots := 0
	lines := 0
	binary := string(d)
	var translateMorse string
	for i, b := range binary {
		char := string(b)
		if char == "0" {
			breaks = 1 + breaks
		}
		if char == "1" {
			dots = 1 + dots
			lines = 1 + lines
		}
		if breaks >= 4 && breaks <= 7 && string(binary[i + 1]) == "1" {
			translateMorse += " "
			breaks = 0
		}
		if dots > 0 && dots <= 3 && string(binary[i + 1]) == "0" {
			translateMorse += "."
			dots = 0
			lines = 0
			breaks = 0
		}
		if lines > 3 && string(binary[i + 1]) == "0" {
			translateMorse += "-"
			dots = 0
			lines = 0
			breaks = 0
		}
	}
	return []byte(translateMorse), nil
}

func (t *translate) Translate2Morse(r io.Reader) ([]byte, error) {
	d, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	words := string(d)
	words = strings.ToUpper(words)
	var translateWords string
	splitDataWords := strings.Split(words, " ")
	numOfWords := len(splitDataWords)
	for j, val := range splitDataWords {
		vLen := len(val)
		for i, c := range val {
			char := string(c)
			translateWords += alphaNumToMorse[char]
			if (i + 1) != vLen {
				translateWords += " "
			}
		}
		if numOfWords != (j + 1) {
			translateWords += "  "
		}
	}
	return []byte(translateWords), nil
}

func (t *translate) Translate2Human(r io.Reader) ([]byte, error) {
	d, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	data := string(d)
	splitDataWords := strings.Split(data, " ")
	var decodedValue string
	for _, val := range splitDataWords {
		if val == "" {
			decodedValue += " "
		}
		decodedValue += morseToAlphaNum[val]
	}

	return []byte(decodedValue), nil
}

func NewTranslate() Translate {
	return &translate{}
}