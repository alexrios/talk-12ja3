package main

import "testing"

func Test(t *testing.T) {
	times := 0
	for {
		times++
		counter := Datarace(0)
		if counter != 4 {
			t.Fatalf("it should be 4 but found %d on execution %d", counter, times)
		}
	}
}

