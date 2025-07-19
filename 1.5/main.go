package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int)
	go func() {
		defer close(jobs)
		const numWork = 10
		const nSec = 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), nSec)
		defer cancel()
	numGenLoop:
		for i := range numWork {
			select {
			case <-ctx.Done():
				break numGenLoop
			default:
			}

			select {
			case <-ctx.Done():
				break numGenLoop
			case jobs <- i:
			}

		}
	}()

	for j := range jobs {
		fmt.Println(j)
		time.Sleep(time.Second)
	}
}
