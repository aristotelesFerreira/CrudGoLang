package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func hardTask(name string) {
	// Finish WaitGroup
	defer wg.Done()

	for i := 0; i < 10; i++ {
		time.Sleep(1 + time.Second)
		fmt.Println("Hard Task %s...\n", name)
	}

	fmt.Println("Hard Task %s DONE\n", name)

}

func main() {
	for i := 0; i < 10; i++ {
		// Add WaitGRoup
		wg.Add(1)
		// Go Routines
		go hardTask(strconv.Itoa(i))
	}
	// Wait WaitGroup
	wg.Wait()
	fmt.Println("DONE")
}
