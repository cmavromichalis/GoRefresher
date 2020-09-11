package main

import (
	"errors"
	"log"
	"regexp"
	"strings"
)

/*
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
*/
func main() {
	/*router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":8080", router))*/

	caesar := CaesarCipher{}

	log.Print("ABC")
	log.Print(caesar.encode("ABC", 1))
	log.Print("XYZ")
	log.Print(caesar.encode("XYZ", 1))
	log.Print("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	log.Print(caesar.encode("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 26))
	log.Print(caesar.encode("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 104))
	log.Print(caesar.encode("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 27))
	log.Print(caesar.encode("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 108))
}

type CaesarCipher struct{}

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var isStringAlphabetic = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

// validate checks the word passed only has letters `^[a-zA-Z]+$`
func (cc CaesarCipher) validate(word string) bool {
	return isStringAlphabetic(word)
}

func (cc CaesarCipher) calculate(word string, rotate int, increment bool) (string, error) {
	if !cc.validate(word) {
		return "", errors.New("The word given has non-alphabet characters")
	}

	letters := strings.Split(word, "")
	var cipherBuilder strings.Builder

	for _, value := range letters {
		if strings.TrimSpace(value) == "" {
			cipherBuilder.WriteString(value)
		} else {
			position := strings.Index(alphabet, value)

			if increment {
				cipherBuilder.WriteString(string(alphabet[(position+rotate)%26]))
			} else {
				cipherBuilder.WriteString(string(alphabet[(position-rotate)%26]))
			}
		}
	}

	return cipherBuilder.String(), nil
}

func (cc CaesarCipher) encode(plaintext string, rotate int) (string, error) {
	return cc.calculate(plaintext, rotate, true)
}

// decode transforms a CaesarCipher into plaintext by shifting
// the letters by the specified integer
func (cc CaesarCipher) decode(plaintext string, rotate int) (string, error) {
	return cc.calculate(plaintext, rotate, false)
}
