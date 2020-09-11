package main

import (
	"testing"
)

func TestKnownGoodEncrypt(t *testing.T) {
	cc := CaesarCipher{}

	input := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	expected_output := "BCDEFGHIJKLMNOPQRSTUVWXYZA"
	rotate := "1"

	output, err := cc.encode(input, rotate)

	if err != nil {
		t.Error(err)
	}

	if output != expected_output {
		t.Error("Known good cipher has failed")
	}
}

func TestKnownGoodDecrypt(t *testing.T) {
	cc := CaesarCipher{}

	input := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	expected_output := "ZABCDEFGHIJKLMNOPQRSTUVWXY"
	rotate := "1"

	output, err := cc.decode(input, rotate)

	if err != nil {
		t.Error(err)
	}

	if output != expected_output {
		t.Error("Known good cipher has failed")
	}
}

func TestCipherWithSpaces(t *testing.T) {
	cc := CaesarCipher{}

	input := "HELLO THERE"
	expected_output := "IFMMP UIFSF"
	rotate := "1"

	output, err := cc.encode(input, rotate)

	if err != nil {
		t.Error(err)
	}

	if output != expected_output {
		t.Error("Known good cipher with spaces has failed")
	}

}

func TestBadCipherInput(t *testing.T) {
	cc := CaesarCipher{}

	input := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0983409890"
	rotate := "1"

	output, err := cc.encode(input, rotate)

	if output != "" {
		t.Error("Expected output to be blank")
	}

	if err == nil {
		t.Error("Expected an error with bad input")
	}
}

func TestBadRotateInput(t *testing.T) {
	cc := CaesarCipher{}

	input := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rotate := "one"

	output, err := cc.encode(input, rotate)

	if output != "" {
		t.Error("Expected output to be blank")
	}

	if err == nil {
		t.Error("Expected an error with bad input")
	}
}
