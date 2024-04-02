package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://localhost:8080/?p=" + os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	body, _ := io.ReadAll(resp.Body)
	log.Println(string(body))
}
