package main

import (
	"fmt"
	"../libjx"
	"os"
)

func main() {
	f := 0
	if len(os.Args) > 1 && os.Args[1] == "-f" {
		f = 1
	}
	if len(os.Args) != 2 + f {
		fmt.Println("Usage: ./readDB <file_name> or ./readDB -f <file_name>")
		return
	}
	object, err := libjx.DeCoder(os.Args[1 + f])
	if err != nil {
		fmt.Printf("\"%s\" read error: %s\n", os.Args[1 + f], err)
		return
	}
	if f == 1 {
		object.Print()
		return
	}
	err = libjx.EnCoder(object, os.Args[1])
	if err != nil {
		fmt.Printf("\"%s\" create error: %s\n", os.Args[1], err)
		return
	}
}


