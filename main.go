package main

import (
	"fmt"
	"net/http"
	"time"
)

// main runs some tests to exercise the environment, including
// loading a timezone and verifying a well-known TLS certificate.
func main() {
	fmt.Println("\nTests from within scratch container:")
	testTZ()
	testTLS()
}

func testTZ() {
	berlin, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		fmt.Printf("Unable to load timezones: %s\n", err)
	} else {
		fmt.Printf("Successfully loaded %q\n", berlin)
	}
}

func testTLS() {
	rsp, err := http.Get("https://google.de")
	if err != nil {
		fmt.Printf("Unable to establish https connection: %s\n", err)
	} else {
		rsp.Body.Close()
		fmt.Println("Successfully established https connection")
	}
}
