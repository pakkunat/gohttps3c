package main

import (
	"crypto/tls"
	"fmt"
	"github.com/lucas-clemente/quic-go/http3"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	r := http3.RoundTripper{
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS13,
			MaxVersion: tls.VersionTLS13,
		},
	}

	req, _ := http.NewRequest("GET", "https://127.0.0.1/", nil)

	resp, err := r.RoundTrip(req)
	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Print(string(body))
}
