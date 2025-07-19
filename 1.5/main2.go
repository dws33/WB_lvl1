package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func numGen(ctx context.Context, numWork int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range numWork {
			select {
			case <-ctx.Done():
				return
			default:
			}

			select {
			case <-ctx.Done():
				return
			case out <- i:
			}

		}
	}()
	return out

}

type cancelFunc func()

func numGen2(numWork int) (<-chan int, cancelFunc) {
	out := make(chan int)
	cancelChan := make(chan struct{})

	cancel := sync.OnceFunc(func() {
		close(cancelChan)
	})
	go func() {
		defer close(out)
		for i := range numWork {
			select {
			case <-cancelChan:
				return
			default:
			}

			select {
			case <-cancelChan:
				return
			case out <- i:
			}

		}
	}()
	return out, cancel

}

func main() {

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()

	//jobs := numGen(ctx, 20)
	jobs, cancelGen := numGen2(20)
	defer cancelGen()
	go func() {
		time.Sleep(5 * time.Second)
		cancelGen()
	}()

	for j := range jobs {
		fmt.Println(j)
		time.Sleep(time.Second)
	}
}
