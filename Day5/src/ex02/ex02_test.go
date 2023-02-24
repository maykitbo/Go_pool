package ex02

import (
	"day5/present"
	"day5/present/exampls"
	"errors"
	"testing"
)

var test_er = errors.New("Error")

func oneTest(t *testing.T, n int, testing, expected []present.Present, expected_er error) {
	got, got_er := GetNCoolestPresents(testing, n)
	if (expected_er != nil && got_er == nil) || (expected_er == nil && got_er != nil) {
		t.Errorf("something is wrong with the errors: expected %v, got %v", expected_er, got_er)
		return
	}
	if got_er != nil {
		return
	}
	if len(got) != len(expected) {
		t.Errorf("expected %v, got %v", expected, got)
		return
	}
	for k, g := range expected {
		if got[k] != g {
			t.Errorf("expected %v, got %v", expected, got)
			return
		}
	}
}

func Test1_1(t *testing.T) {
	oneTest(t, 2, exampls.Array1,
		[]present.Present{{5, 1}, {5, 2}}, nil)
}

func Test1_2(t *testing.T) {
	oneTest(t, 3, exampls.Array1,
		[]present.Present{{5, 1}, {5, 2}, {4, 5}}, nil)
}

func TestError1(t *testing.T) {
	oneTest(t, 10, exampls.Array1, nil, test_er)
}

func Test2(t *testing.T) {
	oneTest(t, 0, exampls.Array2,
		[]present.Present{}, nil)
}

func TestError2(t *testing.T) {
	oneTest(t, -1, exampls.Array2, nil, test_er)
}

func Test3(t *testing.T) {
	oneTest(t, 3, exampls.Array3,
		[]present.Present{{8, 7}, {8, 7}, {8, 7}}, nil)
}

func Test4_1(t *testing.T) {
	oneTest(t, 1, exampls.Array4,
		[]present.Present{{111, 7}}, nil)
}

func Test4_2(t *testing.T) {
	oneTest(t, 4, exampls.Array4,
		[]present.Present{{111, 7}, {44, 1}, {18, 18}, {12, 12}}, nil)
}

func Test4_3(t *testing.T) {
	oneTest(t, 12, exampls.Array4,
		[]present.Present{{111, 7}, {44, 1}, {18, 18}, {12, 12}, {7, 0}, {5, 6}, {1, 3}, {1, 3}, {1, 3}, {1, 4}, {0, 0}, {0, 3}}, nil)
}

func Test5_1(t *testing.T) {
	oneTest(t, 0, exampls.Array5,
		[]present.Present{}, nil)
}

func Test5_2(t *testing.T) {
	oneTest(t, 1, exampls.Array5,
		[]present.Present{{1, 1}}, nil)
}
