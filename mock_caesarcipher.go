package main

import "strconv"

type MockCaesarCipher struct{}

func (mcc MockCaesarCipher) validate(word string) bool {
	r, _ := strconv.ParseBool(word)

	return r
}

func (mcc MockCaesarCipher) calculate(word string, rotateString string, increment bool) (string, error) {
	return word, nil
}

func (mcc MockCaesarCipher) encode(plaintext string, rotate string) (string, error) {
	return plaintext, nil
}

func (mcc MockCaesarCipher) decode(plaintext string, rotate string) (string, error) {
	return plaintext, nil
}
