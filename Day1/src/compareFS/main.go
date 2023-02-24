package main

import (
	"fmt"
	"flag"
	"os"
	"bufio"
)

var (
	old *string
	new *string
)

func init() {
    old = flag.String("old", "", "<old_base path>")
	new = flag.String("new", "", "<new_base path>")
}

func main() {
	flag.Parse()
	old_file,err := os.Open(*old)
	if err != nil {
		fmt.Printf("Error opening file %s: %s\n", *old, err)
		return
	}
	smap := make(map[string]int)
	scanner := bufio.NewScanner(old_file)
	for scanner.Scan() {
		if smap[scanner.Text()] == 0 {
			smap[scanner.Text()] = 1
		} else {
			smap[scanner.Text()]++
		}
	}
	old_file.Close()
	new_file,err := os.Open(*new)
	if err != nil {
		fmt.Printf("Error opening file %s: %s\n", *new, err)
		return
	}
	defer new_file.Close()
	scanner = bufio.NewScanner(new_file)
	for scanner.Scan() {
		if smap[scanner.Text()] == 0 {
			fmt.Printf("ADDED %s\n", scanner.Text())
		} else {
			smap[scanner.Text()]--
		}
	}
	for str,i := range smap {
		if i != 0 {
			fmt.Printf("REMOVED %s\n", str)
		}
	}
}
