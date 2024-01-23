package main

import (
	"fmt"

	"github.com/gosuri/uiprogress"
)

type TestMonitor struct {
	c      chan setEvent
	states []*setState
}

type setState struct {
	id             int
	testsTotal     int
	testsCompleted int
	testsFailed    int
	currentState   string
	bar            *uiprogress.Bar
}

type setEvent struct {
	id            int
	msg           string
	testCompleted int
	testFailed    int
}

func NewMonitor(c chan setEvent, sets []TestSet) TestMonitor {
	states := make([]*setState, 0, len(sets))
	for _, set := range sets {
		bar := uiprogress.AddBar(len(set.Tests)).AppendCompleted().PrependElapsed()
		bar.Width = 25
		state := &setState{
			id:         set.ID,
			testsTotal: len(set.Tests),
			bar:        bar,
		}
		id := set.ID
		bar.PrependFunc(func(b *uiprogress.Bar) string {
			return fmt.Sprintf("TestSet-%d [%s] [pass:%d / fail:%d / total:%d]", id, state.currentState, state.testsCompleted, state.testsFailed, state.testsTotal)
		})
		states = append(states, state)
	}
	return TestMonitor{
		c:      c,
		states: states,
	}
}

func (m *TestMonitor) StartMonitor() {
	go func() {
		uiprogress.Start()
		defer uiprogress.Stop()
		for msg := range m.c {
			id := msg.id - 1
			m.states[id].currentState = msg.msg
			m.states[id].testsCompleted += msg.testCompleted
			m.states[id].testsFailed += msg.testFailed
			if msg.testCompleted > 0 || msg.testFailed > 0 {
				m.states[id].bar.Incr()
			}
		}
	}()
}
