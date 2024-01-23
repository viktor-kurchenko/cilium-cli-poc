package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Verifying Cilium installation ...")
	time.Sleep(2 * time.Second)

	c := make(chan setEvent)
	testSets := ProduceTestSets()
	monitor := NewMonitor(c, testSets)
	monitor.StartMonitor()

	for _, set := range testSets {
		RunAsync(set, c)
	}

	time.Sleep(13 * time.Second)
	close(c)
	fmt.Println("printing logs of failed tests ...")
}
