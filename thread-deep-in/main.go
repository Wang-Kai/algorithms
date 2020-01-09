package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 1; i++ {
		go fibonacci(fmt.Sprintf("===> %d\n", i))
	}

	time.Sleep(time.Second * 100)
}

func fibonacci(n string) {
	for {
		println("====> I love %s !", n)
	}
}
