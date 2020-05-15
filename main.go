package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

// main runs some tests to exercise the environment, including
// loading a timezone and verifying a well-known TLS certificate.
func main() {
	fmt.Println("\nTests from within scratch container:")
	errors := testTZ()
	errors += testTLS()
	os.Exit(errors)
}

func testTZ() (errors int) {
	berlin, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		fmt.Printf("Unable to load timezones: %s\n", err)
		errors++
	} else {
		fmt.Printf("Successfully loaded %q\n", berlin)
	}
	return
}

func testTLS() (errors int) {
	rsp, err := http.Get("https://google.de")
	if err != nil {
		fmt.Printf("Unable to establish https connection: %s\n", err)
		errors++
	} else {
		rsp.Body.Close()
		fmt.Println("Successfully established https connection")
	}
	return
}
