package ex03

import (
	"day5/present"
	"day5/present/exampls"
	"testing"
)

func search(g present.Present, expected []present.Present) bool {
	for _, i := range expected {
		if g == i {
			return true
		}
	}
	return false
}

func oneTest(t *testing.T, n int, testing, expected []present.Present) {
	got := GrabPresents(testing, n)
	if n < 0 {
		n = 0
	}
	g_c, g_p := 0, 0
	for _, g := range got {
		g_c += g.Size
		g_p += g.Value
	}
	if g_c > n {
		t.Errorf("overweight: expected %v, got %v", expected, got)
		return
	}
	e_c, e_p := 0, 0
	for _, e := range expected {
		e_c += e.Size
		e_p += e.Value
	}
	if g_p != e_p {
		t.Errorf("value: expected %v, got %v", expected, got)
		return
	}
}

func Test1_1(t *testing.T) {
	oneTest(t, 2, exampls.Array1,
		[]present.Present{{5, 1}, {3, 1}})
}

func Test1_2(t *testing.T) {
	oneTest(t, 0, exampls.Array1,
		[]present.Present{})
}

func Test1_3(t *testing.T) {
	oneTest(t, 3, exampls.Array1,
		[]present.Present{{5, 1}, {5, 2}})
}

func Test2(t *testing.T) {
	oneTest(t, 5, exampls.Array2,
		[]present.Present{})
}

func Test3_1(t *testing.T) {
	oneTest(t, 5, exampls.Array3,
		[]present.Present{})
}

func Test3_2(t *testing.T) {
	oneTest(t, 25, exampls.Array3,
		[]present.Present{{8, 7}, {8, 7}, {8, 7}})
}

func Test4_1(t *testing.T) {
	oneTest(t, 5, exampls.Array4,
		[]present.Present{{44, 1}, {7, 0}, {1, 4}})
}

func Test4_2(t *testing.T) {
	oneTest(t, 15, exampls.Array4,
		[]present.Present{{5, 6}, {111, 7}, {44, 1}, {7, 0}})
}

func Test4_3(t *testing.T) {
	oneTest(t, 25, exampls.Array4,
		[]present.Present{{12, 12}, {111, 7}, {44, 1}, {7, 0}, {1, 4}})
}

func Test4_4(t *testing.T) {
	oneTest(t, 7, exampls.Array4,
		[]present.Present{{111, 7}})
}

func Test4_5(t *testing.T) {
	oneTest(t, 9, exampls.Array4,
		[]present.Present{{111, 7}, {44, 1}, {7, 0}})
}

func Test4_6(t *testing.T) {
	oneTest(t, 1, exampls.Array4,
		[]present.Present{{44, 1}})
}

func Test4_7(t *testing.T) {
	oneTest(t, 48, exampls.Array4,
		[]present.Present{{5, 6}, {12, 12}, {111, 7}, {44, 1}, {18, 18}, {7, 0}, {1, 4}})
}

func Test5_1(t *testing.T) {
	oneTest(t, 5, exampls.Array5,
		[]present.Present{{1, 1}})
}

func Test5_2(t *testing.T) {
	oneTest(t, -5, exampls.Array5,
		[]present.Present{})
}
