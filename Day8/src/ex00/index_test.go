package ex00

import (
	"fmt"
	"testing"
	"math/rand"
	"errors"
)

var (
	test_func func([]int, int) (int, error) = GetElement
	test_error error = errors.New("test error")
)


func oneTest(t *testing.T, arr []int, index int, er error) {
	n, get_er := test_func(arr, index)
	if (er != nil && get_er == nil) || (er == nil && get_er != nil) {
		fmt.Println("error: input:", arr, "index:", index, "error:", get_er, "expected error:", er)
		t.Fail()
		return
	}
	if er != nil {
		return
	}
	if index >= len(arr) {
		fmt.Println("wrong test")
		return
	}
	if n != arr[index] {
		fmt.Println("result: input:", arr, "index:", index, "got:", n, "expected:", arr[index])
		t.Fail()
	}
}

func Test1(t *testing.T) {
	oneTest(t, []int{1, 2, 3, 4}, 0, nil)
}

func Test2(t *testing.T) {
	oneTest(t, []int{1, 2, 3, 4}, 3, nil)
}

func Test3(t *testing.T) {
	oneTest(t, rand.Perm(1234000), 123456, nil)
}

func Test4(t *testing.T) {
	oneTest(t, rand.Perm(9999), 10000, test_error)
}

func Test5(t *testing.T) {
	oneTest(t, rand.Perm(5), -1, test_error)
}
