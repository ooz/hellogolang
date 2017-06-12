package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.heise.de")
	if err != nil {
		log.Fatalf("fetch: %v", err)
	}
	b, err := ioutil.ReadAll(resp.Body) // io.Reader
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", b)
}