package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type CaesarCipher struct{}

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var isStringAlphabetic = regexp.MustCompile(`^[a-zA-Z ]+$`).MatchString

// validate checks the word passed only has letters `^[a-zA-Z]+$`
func (cc CaesarCipher) validate(word string) bool {
	return isStringAlphabetic(word)
}

func (cc CaesarCipher) calculate(word string, rotateString string, increment bool) (string, error) {
	if !cc.validate(word) {
		return "", errors.New("The word given has non-alphabet characters")
	}

	rotate, err := strconv.Atoi(rotateString)
	if err != nil {
		return "", err
	}

	letters := strings.Split(strings.ToUpper(word), "")
	var cipherBuilder strings.Builder

	for _, value := range letters {
		if strings.TrimSpace(value) == "" {
			cipherBuilder.WriteString(value)
		} else {
			position := strings.Index(alphabet, value)

			if increment {
				cipherBuilder.WriteString(string(alphabet[(position+rotate)%26]))
			} else {
				if (position - rotate) == -1 {
					cipherBuilder.WriteString(string(alphabet[25]))
				} else {
					cipherBuilder.WriteString(string(alphabet[(position-rotate)%26]))
				}
			}
		}
	}

	return cipherBuilder.String(), nil
}

func (cc CaesarCipher) encode(plaintext string, rotate string) (string, error) {
	return cc.calculate(plaintext, rotate, true)
}

// decode transforms a CaesarCipher into plaintext by shifting
// the letters by the specified integer
func (cc CaesarCipher) decode(plaintext string, rotate string) (string, error) {
	return cc.calculate(plaintext, rotate, false)
}
