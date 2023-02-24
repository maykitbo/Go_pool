package main

import (
	"fmt"
	"os/exec"
)

func main() {
	runOneTest("day5/ex00")
	runOneTest("day5/ex01")
	runOneTest("day5/ex02")
	runOneTest("day5/ex03")
}

func runOneTest(name string) {
	cmd := exec.Command("go", "test", name)
	out, err := cmd.Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			fmt.Println(string(ee.Stderr))
		} else {
			fmt.Println(err)
		}
		return
	}
	fmt.Println(string(out))
}
