package statistics

import (
	"fmt"
	"testing"
)

func oneTest(t *testing.T, expected, got []string) {
	if len(expected) != len(got) {
		t.Errorf("%q", "")
	}
	for k,i := range expected {
		if i != got[k] {
			t.Logf("expected %q, got %q", i, got[k])
		}
	}
}

func stringTest(t *testing.T, expected, got float64, name string) {
	strexpected := fmt.Sprintf("%.6f", expected)
	strgot := fmt.Sprintf("%.6f", got)
	if strexpected != strgot {
		fmt.Printf("%s: Expected %s, got %s\n", name, strexpected, strgot)
		t.Fail()
	}
}

func fullTest(t *testing.T, expected []float64, in []int) {
	if len(expected) < 4 { t.Fail() }
	stringTest(t, expected[0], meanFunc(in), "mean")
	stringTest(t, expected[1], median(in), "median")
	stringTest(t, expected[2], float64(mode(in)), "mode")
	stringTest(t, expected[3], sD(in, meanFunc(in)), "sd")
}

func TestStatistics(t *testing.T) {
	oneTest(t, []string{"Mean: 1.00\n"}, Statistics([]int{1}, map[int]bool{1: true}))
	oneTest(t, []string{"Mean: 11.17\n", "Median: 18.00\n"}, Statistics([]int{1, -7, 18, 18, 19, 18}, map[int]bool{1: true, 2: true}))
	oneTest(t, []string{"Mode: 0\n", "SD: 0.00\n"}, Statistics([]int{0, 0, 0}, map[int]bool{3: true, 4: true}))
	fullTest(t, []float64{22.142857142857, 23, 23, 12.298996142875}, []int{10, 2, 38, 23, 38, 23, 21})
	fullTest(t, []float64{6.5, 6.5, -7, 8.0777472107018}, []int{-7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
	mas := make([]int, 100000)
	for k := 0; k < 100000; k++ {
		mas[k] = k
	}
	fullTest(t, []float64{49999.5, 49999.5, 0, 28867.513458}, mas)
}
