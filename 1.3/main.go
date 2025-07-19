package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	numWorker := flag.Int("nw", 0, "")
	flag.Parse()

	var wg sync.WaitGroup
	wg.Add(*numWorker)

	numStream := make(chan int)
	for range *numWorker {
		go func() {
			defer wg.Done()
			for n := range numStream {
				fmt.Println(n)
			}
		}()
	}

	const numData = 20
	for i := range numData {
		numStream <- i
	}
	close(numStream)
	wg.Wait()
}
