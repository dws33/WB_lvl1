package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	numWorker := flag.Int("nw", 0, "")
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(*numWorker)

	numStream := make(chan int)
	for range *numWorker {
		go func() {
			defer wg.Done()
			for n := range numStream {
				fmt.Println(n)
				time.Sleep(time.Second)
			}
		}()
	}

	const numData = 20
MyLoop:
	for i := range numData {
		select {
		case <-ctx.Done():
			break MyLoop
		case numStream <- i:
		}
	}
	close(numStream)
	wg.Wait()
}
