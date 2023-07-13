package main

import (
	"fmt"
	"time"
)

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}
func main() {
	start := time.Now()
	var c <-chan int
	select {
	case <-c:
	default:
		fmt.Printf("In default after %v\n\n", time.Since(start))
	}
}
