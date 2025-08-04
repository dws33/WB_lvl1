package main

import (
	"fmt"
	"time"
)

func Sleep(duration time.Duration) {
	done := make(chan struct{})

	go func() {
		timer := time.NewTimer(duration)
		<-timer.C
		close(done)
	}()

	<-done
}

func main() {
	fmt.Println("Начало:", time.Now().Format("15:04:05"))
	Sleep(2 * time.Second)
	fmt.Println("Конец:  ", time.Now().Format("15:04:05"))
}
