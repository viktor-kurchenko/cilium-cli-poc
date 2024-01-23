package main

import (
	"time"
)

type TestSet struct {
	ID    int
	Tests []TestCase
}

type TestCase struct {
	ID   int
	Name string
}

func (t TestCase) Run() error {
	// imitate work
	time.Sleep(2 * time.Second)
	return nil
}

var testNames = [][]string{
	{"no-unexpected-packet-drops", "no-policies", "no-policies-extra", "allow-all-except-world", "client-ingress"},
	{"client-ingress-knp", "allow-all-with-metrics-check", "all-ingress-deny-knp", "all-egress-deny", "all-egress-deny-knp"},
	{"all-entities-deny"},
	{"cluster-entity", "host-entity", "echo-ingress", "echo-ingress-knp", "client-ingress-icmp", "client-egress"},
	{"client-egress-knp", "client-egress-expression", "client-egress-expression-knp", "client-with-service-account-egress-to-echo"},
	{"client-egress-to-echo-service-account", "to-entities-world", "to-cidr-external", "to-cidr-external-knp"},
}

func ProduceTestSets() []TestSet {
	setCounter := 1
	testCounter := 1
	sets := make([]TestSet, 0, len(testNames))
	for _, names := range testNames {
		tests := make([]TestCase, 0, len(names))
		for _, name := range names {
			tests = append(tests, TestCase{
				ID:   testCounter,
				Name: name,
			})
			testCounter++
		}
		sets = append(sets, TestSet{
			ID:    setCounter,
			Tests: tests,
		})
		setCounter++
	}
	return sets
}
