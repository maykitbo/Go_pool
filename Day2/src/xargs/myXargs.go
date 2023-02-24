package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func isPiped() bool {
	fileInfo, _ := os.Stdin.Stat()
	return (fileInfo.Mode() & os.ModeCharDevice) == 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if len(os.Args) < 2 {
		fmt.Println("Error: Executable file not found after myXargs")
		return
	}
	if !isPiped() {
		fmt.Println("Error: This program should be used with input from a pipe.")
		return
	}
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	args := append(os.Args[2:], input...)
	cmd := exec.Command(os.Args[1], args...)
	output, _ := cmd.Output()
	for _,out := range output {
		fmt.Printf("%s", string(out))
	}
}
