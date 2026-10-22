package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	cfg, err := ParseConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(2)
	}

	executor := NewHTTPExecutor()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		cfg.Duration,
	)
	defer cancel()

	start := time.Now()
	results := RunWorkers(ctx, cfg, executor)

	var total int
	var transportErrors int
	var httpErrors int

	for result := range results {
		total++

		if result.Err != nil {
			transportErrors++
		} else if result.StatusCode < 200 ||
			result.StatusCode >= 400 {
			httpErrors++
		}
	}

	elapsed := time.Since(start)
	if elapsed <= 0 {
		elapsed = time.Nanosecond
	}

	fmt.Println("\nAPI Load Test Complete")
	fmt.Println("----------------------")
	fmt.Println("Target:          ", cfg.URL)
	fmt.Println("Workers:         ", cfg.Concurrency)
	fmt.Println("Elapsed:         ", elapsed.Round(time.Millisecond))
	fmt.Println("Completed:       ", total)
	fmt.Println("Transport errors:", transportErrors)
	fmt.Println("HTTP errors:     ", httpErrors)
	fmt.Printf("Throughput:       %.2f req/s\n",
		float64(total)/elapsed.Seconds())
}
