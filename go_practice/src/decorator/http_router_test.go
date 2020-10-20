package main

import (
	"log"
	"net/http"
	"testing"
)

func TestHttpRouterDecorator(t *testing.T) {
	http.HandleFunc("/v1/hello", WithServerHeader(WithAuthCookie(hello)))
	http.HandleFunc("/v2/hello", WithServerHeader(WithBasicAuth(hello)))
	http.HandleFunc("/v3/hello", WithServerHeader(WithBasicAuth(WithDebugLog(hello))))
	http.HandleFunc("/v4/hello", Handler(hello, WithServerHeader, WithBasicAuth, WithDebugLog))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
