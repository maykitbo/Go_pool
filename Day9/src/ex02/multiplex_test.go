package ex02

import (
	"testing"
)

func TestMultiplex(t *testing.T) {
	c1 := make(chan interface{})
	c2 := make(chan interface{})
	c3 := make(chan interface{})
	go func() {
		c1 <- 1
		c2 <- 2
		c3 <- 3
		close(c1)
		close(c2)
		close(c3)
	}()
	out := multiplex(c1, c2, c3)
	expected := map[interface{}]bool{1: true, 2: true, 3: true}
	for v := range out {
		_, ok := expected[v]
		if !ok {
			t.Fatalf("Unexpected value received: %v", v)
		}
		delete(expected, v)
		if len(expected) == 0 {
			break
		}
	}
	if len(expected) > 0 {
		t.Fatalf("Not all expected values were received")
	}
}


