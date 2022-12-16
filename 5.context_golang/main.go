package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctxCancel, cancel := context.WithCancel(ctx)

	// defer cancel()
	cancel()

	ctcTimeout, cancelTimeout := context.WithTimeout(ctxCancel, time.Second*2)
	defer cancelTimeout()

	select {
	case <-ctcTimeout.Done():
		fmt.Println("timeout")

	case <-ctxCancel.Done():
		fmt.Println("cancel")
	}

}
