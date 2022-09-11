package main

import (
	"fmt"
	"math/rand"
	"time"
)

const numWrites = 5

func boringGenerator(msg string) <-chan string {
	ch := make(chan string)

	go func() {
		for i := 0; i < numWrites; i++ {
			ch <- fmt.Sprintf("%s - %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(5000)+1000) * time.Millisecond)
		}
		close(ch)
	}()

	return ch
}

func main() {
	fmt.Println("Hello boring")

	genCh := boringGenerator("Boring...")

	for msg := range genCh {
		fmt.Println(msg)
	}

	fmt.Println("Ended")
}
