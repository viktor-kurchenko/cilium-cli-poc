package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Verifying Cilium installation ...")
	time.Sleep(2 * time.Second)

	c := make(chan setEvent)
	testSets := ProduceTestSets()
	monitor := NewMonitor(c, testSets)
	monitor.StartMonitor()

	wg := &sync.WaitGroup{}
	for _, set := range testSets {
		wg.Add(1)
		RunAsync(set, c, wg)
	}

	wg.Wait()
	close(c)
	fmt.Println("printing logs of failed tests ...")
}
