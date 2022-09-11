package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func boringGenerator(msg string) <-chan string {
	ch := make(chan string)
	rand.Seed(time.Now().UnixNano())
	go func() {
		for i := 0; i < 5; i++ {
			ch <- fmt.Sprintf("%s - %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(5000)+1000) * time.Millisecond)
		}
		close(ch)
	}()

	return ch
}

func fanIn(cs ...<-chan string) <-chan string {
	ch := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(len(cs))

	for _, c := range cs {
		go func(localC <-chan string) {
			defer wg.Done()
			for in := range localC {
				ch <- in
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}
func main() {
	fannedInCh := fanIn(boringGenerator("Joe"), boringGenerator("Elsa"), boringGenerator("Max"))

	for v := range fannedInCh {
		fmt.Println(v)
	}
	fmt.Println("Still boring")
}
