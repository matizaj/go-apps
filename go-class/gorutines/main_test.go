package main

import (
	"strings"
	"testing"
)

func Test_printSomething(t *testing.T) {
	msg = "hi there"
	wg.Add(2)
	go updateMessage("test")
	go updateMessage("test2")
	wg.Wait()
	if !strings.Contains(msg, "test") {
		t.Errorf("something goes wrong")
	}
}
