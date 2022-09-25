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
	binary := string(d)
	return []byte(binary), nil
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
			translateWords += " " + "/" + " "
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

	// strings.TrimSpace
	splitDataWords := strings.Split(data, " ")
	// numOfWords := len(splitDataWords)
	var decodedValue string
	for _, val := range splitDataWords {
		if val == "/" {
			decodedValue += " "
		}
		decodedValue += morseToAlphaNum[val]
		// if numOfWords == (i + 1) {
		// 	decodedValue += " " + "/" + " "
		// }
	}

	return []byte(decodedValue), nil
}

func NewTranslate() Translate {
	return &translate{}
}