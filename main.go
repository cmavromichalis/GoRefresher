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
	caesar := CaesarCipher{}
	cipher, err := caesar.encode(ps.ByName("word"), (ps.ByName("rotate")))

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, cipher)
}

func main() {

	router := httprouter.New()
	router.GET("/encrypt/:word/:rotate", Encrypt)
	router.GET("/decrypt/:word/:rotate", Decrypt)

	log.Fatal(http.ListenAndServe(":8080", router))
}
