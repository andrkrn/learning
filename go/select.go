package main

import "fmt"

func main() {
	v := make(chan uint64)
	q := make(chan uint64)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-v)
		}
		q <- 0
	}()

	A(v, q)
}

func A(v chan uint64, q chan uint64) {
	var x, y uint64 = 0, 1
	for {
		select {
		case v <- x:
			x, y = y, x+y
		case <-q:
			return
		}
	}
}
