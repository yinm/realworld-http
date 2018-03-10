package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	// create Client
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	// do communicate
	resp, err := client.Get("https://localhost:18443")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(dump))
}
