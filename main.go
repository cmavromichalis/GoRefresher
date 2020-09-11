package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Decrypt(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	caesar := CaesarCipher{}
	cipher, err := caesar.decode(ps.ByName("word"), (ps.ByName("rotate")))

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, cipher)
}

func Encrypt(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var caesar Cipher

	if ps.ByName("cipher") == "caesar" {
		caesar = CaesarCipher{}
	} else {
		caesar = MockCaesarCipher{}
	}

	cipher, err := caesar.encode(ps.ByName("word"), (ps.ByName("rotate")))

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, cipher)
}

func App() http.Handler {
	router := httprouter.New()
	router.GET("/encrypt/:cipher/:word/:rotate", Encrypt)
	router.GET("/decrypt/:cipher/:word/:rotate", Decrypt)

	return router
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", App()))
}
