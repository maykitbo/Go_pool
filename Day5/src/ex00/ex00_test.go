package ex00

import (
	"day5/tree/exampls"
	"testing"
)

func oneTest(t *testing.T, got bool, expected bool) {
	if got != expected {
		t.Errorf("expected %t, got %t", expected, got)
	}
}

func Test1(t *testing.T) {
	oneTest(t, AreToysBalanced(exampls.Root1), true)
}

func Test2(t *testing.T) {
	oneTest(t, AreToysBalanced(exampls.Root2), false)
}

func Test3(t *testing.T) {
	oneTest(t, AreToysBalanced(exampls.Root3), true)
}

func Test4(t *testing.T) {
	oneTest(t, AreToysBalanced(exampls.Root4), true)
}

func Test5(t *testing.T) {
	oneTest(t, AreToysBalanced(exampls.Root5), false)
}

func Test6(t *testing.T) {
	oneTest(t, AreToysBalanced(exampls.Root6), false)
}

func Test7(t *testing.T) {
	oneTest(t, AreToysBalanced(exampls.Root7), true)
}

func Test8(t *testing.T) {
	oneTest(t, AreToysBalanced(exampls.Root8), true)
}

func Test9(t *testing.T) {
	oneTest(t, AreToysBalanced(exampls.Root9), false)
}

func Test10(t *testing.T) {
	oneTest(t, AreToysBalanced(exampls.Root10), false)
}

func Test11(t *testing.T) {
	oneTest(t, AreToysBalanced(exampls.Root11), true)
}

func Test12(t *testing.T) {
	oneTest(t, AreToysBalanced(exampls.Root12), false)
}
