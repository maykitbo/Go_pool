package ex01

import (
	"day5/tree/exampls"
	"testing"
)

func oneTest(t *testing.T, got []bool, expected []bool) {
	if len(got) != len(expected) {
		t.Errorf("expected %t, got %t", expected, got)
	}
	for k, g := range expected {
		if got[k] != g {
			t.Errorf("expected %t, got %t", expected, got)
		}
	}
}

func Test1(t *testing.T) {
	oneTest(t, UnrollGarland(exampls.Root1),
		[]bool{true, true, true, false, true, true, false})
}

func Test2(t *testing.T) {
	oneTest(t, UnrollGarland(exampls.Root2),
		[]bool{true, true, true, true, true, false})
}

func Test3(t *testing.T) {
	oneTest(t, UnrollGarland(exampls.Root3),
		[]bool{false, true, false, false, true})
}

func Test4(t *testing.T) {
	oneTest(t, UnrollGarland(exampls.Root4),
		[]bool{true, true, true, false, true, true, false})
}

func Test5(t *testing.T) {
	oneTest(t, UnrollGarland(exampls.Root5),
		[]bool{true, true, true, true, true, false})
}

func Test6(t *testing.T) {
	oneTest(t, UnrollGarland(exampls.Root6),
		[]bool{true, true, false, false})
}

func Test7(t *testing.T) {
	oneTest(t, UnrollGarland(exampls.Root7), []bool{true})
}

func Test8(t *testing.T) {
	oneTest(t, UnrollGarland(exampls.Root8),
		[]bool{true, true, false, true, true, false, true})
}

func Test9(t *testing.T) {
	oneTest(t, UnrollGarland(exampls.Root9),
		[]bool{true, true, false})
}

func Test10(t *testing.T) {
	oneTest(t, UnrollGarland(exampls.Root10),
		[]bool{false, true, false, true, true})
}

func Test11(t *testing.T) {
	oneTest(t, UnrollGarland(exampls.Root11),
		[]bool{true, true, true, false, true, true, false, true, true, true, false, true, false, true, true, true, true})
}

func Test12(t *testing.T) {
	oneTest(t, UnrollGarland(exampls.Root12),
		[]bool{true, true, true, false, true, true, false, true, true, true, false, true, true, true, false, true})
}
