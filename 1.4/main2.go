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

	tNow := time.Now()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	numStream := make(chan int)
	go func() {
		defer close(numStream)
		const numData = 20
		for i := range numData {
			select {
			case <-ctx.Done():
				return
			default:
			}

			select {
			case <-ctx.Done():
				return
			case numStream <- i:
			}
		}
	}()

	var wg sync.WaitGroup
	wg.Add(*numWorker)

	for range *numWorker {
		go func() {
			defer wg.Done()

			for n := range numStream {
				fmt.Println(n)
				time.Sleep(time.Second)
			}

			//for {
			//
			//	select {
			//	case <-ctx.Done():
			//		return
			//	default:
			//	}
			//
			//	select {
			//	case <-ctx.Done():
			//		return
			//	case n, ok := <-numStream:
			//		if !ok {
			//			return
			//		}
			//
			//		fmt.Println(n)
			//		time.Sleep(time.Second)
			//
			//	}

			//}
		}()
	}

	//	const numData = 20
	//myLoop:
	//	for i := range numData {
	//		select {
	//		case <-ctx.Done():
	//			break myLoop
	//		case numStream <- i:
	//
	//		}
	//
	//	}
	//	close(numStream)

	wg.Wait()
	fmt.Println("end")
	fmt.Println(time.Since(tNow))
	//time.Sleep(time.Second)

}
