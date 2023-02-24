package main

import (
	"day6/logo/logorand"
	"flag"
	"time"
)

func flags() (logorand.Logo, bool, int) {
	defer flag.Parse()
	t := flag.String("t", logorand.RandomString(true), "text in your logo")
	p := flag.String("p", "amazing_logo.png", "file name")
	x := flag.Int("x", 300, "width")
	y := flag.Int("y", 300, "length")
	f := flag.Bool("g", false, "permanent generator")
	T := flag.Int("T", 1000, "generator time duration (millisecond)")
	flag.Parse()
	return logorand.Init(*t, *p, *x, *y), *f, *T
}

func main() {
	logo, f, T := flags()
	logo.Create()
	logo.Save()
	for f == true {
		logo.Create()
		logo.Save()
		time.Sleep(time.Duration(T) * time.Millisecond)
	}
}
