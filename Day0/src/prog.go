package main

import (
	"fmt"
	"./statistics"
	"errors"
)

func VarsInput(err *error) map[int]bool {
	vars := make(map[int]bool)
	fmt.Println("What information would you like to see? Please enter numbers separated by spaces or press enter to see all infomation\n1: Mean\n2: Median\n3: Mode\n4: SD")
	for {
		n := 0
		var c rune
		_,er := fmt.Scanf("%d%c", &n, &c)
		if er != nil {
			break
		}
		if n <= 0 || n > 4 {
			*err = errors.New(fmt.Sprintf("%d: I don't know this option", n))
			break
		}
		vars[n] = true
		if c == '\n' { break }
		if c != ' ' {
			*err = errors.New(fmt.Sprintf("Incorrect input: %c", c))
			break
		}
	}
	if len(vars) == 0 {
		vars[1], vars[2], vars[3], vars[4] = true, true, true, true
	}
	return vars
}

func NumdersInput(err *error) (mas []int) {
	fmt.Println("Enter numbers. Enter any other char to finish")
	for k := 0; ; k++ {
		n := 0
		_,er := fmt.Scanf("%d", &n)
		if er != nil { break }
		mas = append(mas, n)
	}
	if len(mas) == 0 {
		*err = errors.New("Empty bunch")
	}
	return mas
}

func main() {
	var err error
	vars := VarsInput(&err)
	if err != nil {
		fmt.Println(err)
		return
	}
	mas := NumdersInput(&err)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _,str := range statistics.Statistics(mas, vars) {
		fmt.Printf("%s", str)
	}
}
