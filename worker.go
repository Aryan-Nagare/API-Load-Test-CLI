package main

import (
	"context"
	"sync"
)

func RunWorkers(
	ctx context.Context,
	cfg Config,
	executor *HTTPExecutor,
) <-chan RequestResult {
	results := make(chan RequestResult, cfg.Concurrency*2)

	var wg sync.WaitGroup

	for i := 0; i < cfg.Concurrency; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for {
				// Do not start another request after cancellation.
				if ctx.Err() != nil {
					return
				}

				result := executor.Do(ctx, cfg.URL)

				// Always publish the result of a request that started.
				// The main goroutine must drain this channel.
				results <- result

				if ctx.Err() != nil {
					return
				}
			}
		}()
	}

	// Close the results channel only after every worker exits.
	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}
