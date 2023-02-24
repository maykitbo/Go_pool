package statistics

import (
	"fmt"
	"sort"
	"math"
)

func median(mas []int) (median float64) {
	sort.Ints(mas)
    l := len(mas)
    if l == 0 {
        median = 0
    } else if l%2 == 0 {
        median = float64(mas[l/2-1] + mas[l/2]) / 2.0
    } else {
        median = float64(mas[l/2])
    }
	return median
}

func meanFunc(mas []int) (sum float64) {
	for _,i := range mas {
		sum += float64(i) / float64(len(mas))
	}
	return sum
}

func mode(mas []int) (mode int) {
	mMap := make(map[int]int)
	for _,n := range mas {
		mMap[n] += 1
	}
	max := 0
	for _, key := range mas {
		count := mMap[key]
		if count > max {
			mode = key
			max = count
		}
	}
	return mode
}

func sD(mas []int, mean float64) (sum float64) {
	for _,i := range mas {
		sum += math.Pow((mean - float64(i)), 2) / float64(len(mas))
	}
	return math.Sqrt(sum)
}

func Statistics(mas []int, vars map[int]bool) (result []string) {
	var mean float64
	if vars[1] || vars[4] {
		mean = meanFunc(mas)
		if vars[1] { result = append(result, fmt.Sprintf("Mean: %.2f\n", mean)) }
	}
	if vars[2] { result = append(result, fmt.Sprintf("Median: %.2f\n", median(mas))) }
	if vars[3] { result = append(result, fmt.Sprintf("Mode: %d\n", mode(mas))) }
	if vars[4] { result = append(result, fmt.Sprintf("SD: %.2f\n", sD(mas, mean))) }
	return result
}