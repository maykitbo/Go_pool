package ex02

import (
	"day5/present"
	"errors"
	"sort"
)

func GetNCoolestPresents(presents []present.Present, n int) ([]present.Present, error) {
	if n > len(presents) || n < 0 {
		return nil, errors.New("Invalid input")
	}
	sort.SliceStable(presents, func(i, j int) bool {
		if presents[i].Value == presents[j].Value {
			return presents[i].Size < presents[j].Size
		}
		return presents[i].Value > presents[j].Value
	})
	return presents[:n], nil
}
