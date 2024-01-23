package main

import (
	"fmt"
	"sync"
	"time"
)

func RunAsync(set TestSet, c chan setEvent, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		createNamespace(c, set.ID)
		deployTestPods(c, set.ID)
		success := true
		for _, t := range set.Tests {
			c <- setEvent{
				id:  set.ID,
				msg: fmt.Sprintf("running test: %s", t.Name),
			}
			pass, fail := 1, 0
			if err := t.Run(); err != nil {
				pass, fail = 0, 1
				success = false
			}
			c <- setEvent{
				id:            set.ID,
				msg:           t.Name,
				testCompleted: pass,
				testFailed:    fail,
			}
		}
		msg := "test set pass"
		if !success {
			msg = "test set failed"
		}
		c <- setEvent{
			id:  set.ID,
			msg: msg,
		}
	}()
}

func createNamespace(c chan setEvent, setID int) {
	c <- setEvent{
		id:  setID,
		msg: fmt.Sprintf("creating namespace: test-set-%d ...", setID),
	}
	time.Sleep(2 * time.Second)
}

func deployTestPods(c chan setEvent, setID int) {
	c <- setEvent{
		id:  setID,
		msg: fmt.Sprintf("deploying pods in namespace: test-set-%d ...", setID),
	}
	time.Sleep(5 * time.Second)
}
