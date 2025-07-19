package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {

	aJobs := make(chan int)
	go func() {
		defer close(aJobs)
		for i := range 5 {
			aJobs <- i
		}
	}()

	go func() {
		for range aJobs {
		}
		fmt.Println("exit by chan close")
	}()

	////////////////////////////////////////

	bJobs := make(chan int)
	go func() {
		defer close(bJobs)
		for i := range 5 {
			bJobs <- i
		}
	}()

	go func() {
		const stopVal = 3
		for j := range bJobs {
			if j == stopVal {
				fmt.Println("exit by stop val")
				return
			}
		}
	}()

	////////////////////////////////////////

	cancelChan := make(chan struct{})
	time.AfterFunc(2*time.Second, func() { close(cancelChan) })

	go func() {
	myLoop:
		for range 5 {
			select {
			case <-time.After(time.Second):
				continue
			case <-cancelChan:
				fmt.Println("exit by cancel chan")
				break myLoop
			}
		}
	}()

	////////////////////////////////////////

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func() {
	myLoop:
		for range 5 {
			select {
			case <-time.After(time.Second):
				continue
			case <-ctx.Done():
				fmt.Println("exit by cancel chan (ctx)")
				break myLoop
			}
		}
	}()

	////////////////////////////////////////

	go func() {
		go func() {
			time.Sleep(3 * time.Second)
			fmt.Println("exit by goexit") // will always print
		}()
		runtime.Goexit()
	}()

	time.Sleep(10 * time.Second)

}
