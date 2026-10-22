package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	cfg, err := ParseConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(2)
	}

	executor := NewHTTPExecutor()

	result := executor.Do(
		context.Background(),
		cfg.URL,
	)

	fmt.Println("API Load Tester")
	fmt.Println("Status: ", result.StatusCode)
	fmt.Println("Latency:", result.Latency)
	fmt.Println("Bytes:  ", result.Bytes)

	if result.Err != nil {
		fmt.Fprintln(os.Stderr, "Request error:", result.Err)
		os.Exit(1)
	}
}
