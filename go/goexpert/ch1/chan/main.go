package main

import (
	"context"
	"fmt"
	"time"
)

func send(ctx context.Context, ch chan int64) {
	for {
		time.Sleep(time.Second * 5)
		select {
		case <-ctx.Done():
			return
		default:
			ch <- time.Now().Unix()
		}
	}
}

func recvx(ctx context.Context, ch chan int64) {
	for {
		time.Sleep(time.Second * 5)
		select {
		case <-ctx.Done():
			return
		case t := <-ch:
			fmt.Printf("recvx: %d\n", t)
		}
	}
}

func main() {
	ch := make(chan int, 5)

	for i := 0; i < 10; i++ {
		select {
		default:
			ch <- i
		case a := <-ch:
			fmt.Printf("recvx: %v\n", a)
		}
	}
}
