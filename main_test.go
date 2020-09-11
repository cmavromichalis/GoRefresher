package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestWebEncrypt(t *testing.T) {
	go func() {
		log.Fatal(http.ListenAndServe(":8080", App()))
	}()

	res, err := http.Get("http://localhost:8080/encrypt/mock/ABC/1/")
	if err != nil {
		t.Error(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(string(body), "ABC") {
		t.Error("Web Encrypt did not return as expected")
	}
}
