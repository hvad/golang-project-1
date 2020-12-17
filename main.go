package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
)

func checkUrl(url string) {
	// Url check
	req, err := http.Get(url)
	if err != nil {
		fmt.Printf("Url not available - error %v.\n", err)
	}
	defer req.Body.Close()

	if req.StatusCode != 200 {
		fmt.Printf("Check Critical - status code error: %d %s.\n", req.StatusCode, req.Status)
	}
	fmt.Printf("Check OK - Url %s.\n", url)
}

func main() {

	url := flag.String("url", "http://example.com", "Url to check")
	flag.Parse()

	// Schedule check URL every 5 minutes
	s1 := gocron.NewScheduler(time.UTC)
	_, err := s1.Every(5).Minutes().Do(checkUrl, *url)
	if err != nil {
		log.Fatalf("error creating job: %v", err)
	}
	s1.StartBlocking()
}
