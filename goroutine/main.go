package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("Hello World")

	// wait group synchronize
	wg := sync.WaitGroup{}
	var mu sync.Mutex

	numberWithoutMu := 0
	numberWithMu := 0

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			timems := rand.Intn(3000)
			time.Sleep(time.Duration(time.Duration(timems) * time.Millisecond))
			fmt.Println("Hello World from wg, sleep for", timems, "ms", i)
		}()
	}

	// channel communication
	wordCh := make(chan string, 10)

	//mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			time.Sleep(time.Duration(time.Duration(100) * time.Millisecond))
			mu.Lock()
			numberWithMu++
			fmt.Println("=======Hello from mutex", numberWithMu)
			mu.Unlock()

			numberWithoutMu++
			fmt.Println("=======Hello from no mutex", numberWithoutMu)

			if i < 10 {
				doSay(fmt.Sprintf("Hello world from channel %d", i), wordCh)
			}
		}()
	}

	// for i := 0; i < 1000; i++ {
	// 	go doSay(fmt.Sprintf("Hello world from channel %d", i), wordCh)
	// }

	for i := 0; i < 10; i++ {
		fmt.Println(<-wordCh)
	}

	wg.Wait()

	fmt.Println("mutex result: ", numberWithMu)
	fmt.Println("no mutex result: ", numberWithoutMu)

	fmt.Println("Done")

}

func doSay(word string, ch chan string) {
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	ch <- word
}
