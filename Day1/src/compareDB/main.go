package main

import (
	"../libjx"
	"fmt"
	"flag"
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
	oldObject,err := libjx.DeCoder(*old)
	if err != nil {
		fmt.Println(err)
		return
	}
	newObject,err := libjx.DeCoder(*new)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _,str := range libjx.Compare(oldObject, newObject) {
		fmt.Println(str)
	}

}


