package main

import (
	"fmt"
	"os"
)

func main() {
	cfg, err := ParseConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(2)
	}

	fmt.Println("API Load Tester")
	fmt.Println("Target URL: ", cfg.URL)
	fmt.Println("Concurrency:", cfg.Concurrency)
	fmt.Println("Duration:   ", cfg.Duration)
	fmt.Println("Rate:       ", cfg.Rate, "requests/second")
}
