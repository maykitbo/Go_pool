package coins

import (
	"fmt"
	"testing"
	"math/rand"
)

var test_func func(int, []int) []int = MinCoins2

func inArr(i int, arr []int) bool {
	for _,j := range arr {
		if i == j {
			return true
		}
	}
	return false
}

func error1(b *testing.B, arr, denomination []int, amount, price int) {
	fmt.Println(b.Name() + ":\ngot:", arr, "denominations: ", denomination, "expected: ", amount, "price: ", price)
	b.Fail()
}

func fullTest(b *testing.B, denomination []int, price, amount int) {
	arr := test_func(price, denomination)
	sum := 0
	if len(arr) == 0 && amount == 0 {
		return
	}
	if len(arr) != amount {
		error1(b, arr, denomination, amount, price)
		return
	}
	for _,i := range arr {
		if !inArr(i, denomination) {
			error1(b, arr, denomination, amount, price)
			return
		}
		sum += i
	}
	if sum != price {
		error1(b, arr, denomination, amount, price)
	}
}

func BenchmarkAndTest1(b *testing.B) {
	fullTest(b, []int{5, 10, 50, 100}, 30,
		3)
}

func BenchmarkAndTest2(b *testing.B) {
	fullTest(b, []int{5, 10, 50, 100}, 35,
		4)
}

func BenchmarkAndTest3(b *testing.B) {
	fullTest(b, []int{5, 7, 8, 10}, 14,
		2)
}

func BenchmarkAndTest4(b *testing.B) {
	fullTest(b, []int{1, 10, 20, 100}, 13,
		4)
}

func BenchmarkAndTest5(b *testing.B) {
	fullTest(b, []int{1, 10, 90, 100}, 180,
		2)
}

func BenchmarkAndTest6(b *testing.B) {
	fullTest(b, []int{1, 10, 90, 100, 1000, 3050}, 12345,
		13)
}

func BenchmarkAndTest6_1(b *testing.B) {
	fullTest(b, []int{100, 3050, 90, 1, 1000, 10}, 12345,
		13)
}

func BenchmarkAndTest7(b *testing.B) {
	fullTest(b, []int{4, 10, 90, 100, 1000, 3050}, 1,
		0)
}

func BenchmarkAndTest8(b *testing.B) {
	fullTest(b, []int{4, -90, 100, -1000, 3050}, 4,
		0)
}

func BenchmarkAndTest9(b *testing.B) {
	fullTest(b, []int{4, 4, 8, 8, 2, 2, 64, 64, 32, 32, 256, 128, 256, 128}, 1024,
		4)
}

func BenchmarkAndTest10(b *testing.B) {
	fullTest(b, []int{}, 1,
		0)
}

func BenchmarkAndTest11(b *testing.B) {
	fullTest(b, []int{7, 99, 77, 11, 456, 99, 3333, 2323, 958, 33, 121, 1, 19, 17, 999, 534, 327, 88, 56, 790, 12, 345, 6666, 141, 76}, 48977,
		11)
}

func Benchmark1(b *testing.B) {
	for i := 0; i < 1000; i++ {
		test_func(i, []int{1, 10, 90, 100})
	}
}

func Benchmark2(b *testing.B) {
	for i := 0; i < 1000; i++ {
		test_func(i * 2, []int{1, 10, 90, 100})
	}
}

func Benchmark3(b *testing.B) {
	for i := 0; i < 1000; i++ {
		test_func(i + 3, []int{99, 1, 7, 783, 467, 331, 90, 2, 11, 1237, 1999, 22, 177, 95, 44})
	}
}

func Benchmark4(b *testing.B) {
	for i := 0; i < 1000; i++ {
		test_func(i + 77, []int{88, 99, 11, 22, 44, 55, 33, 77, 66})
	}
}

func Benchmark5(b *testing.B) {
	arr := rand.Perm(5000)
	test_func(777, arr)
}

func Benchmark6(b *testing.B) {
	arr := rand.Perm(6666)
	test_func(666, arr)
}

func Benchmark7(b *testing.B) {
	arr := rand.Perm(1234)
	test_func(4321, arr)
}

func Benchmark8(b *testing.B) {
	arr := rand.Perm(1234000)
	test_func(4321, arr)
}

func Benchmark9(b *testing.B) {
	for i := 0; i < 1000; i++ {
		arr := rand.Perm(i + 1)
		test_func((i + 77) * 2, arr)
	}
}

