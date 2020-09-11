package main

import (
	"net/http/httptest"
	"testing"
)

func TestWebEncrypt(t *testing.T) {
	testServer := httptest.NewServer(App())
	defer testServer.Close()

	testServer.URL

	print(rr.Body.String())

	if rr.Body.String() != "BCD" {
		t.Error("Web Encrypt failed")
	}

}
