package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sync"
	"errors"
	"unicode/utf8"
)

var (
	linesFlag = flag.Bool("l", false, "count lines")
	charsFlag = flag.Bool("m", false, "count characters")
	wordsFlag = flag.Bool("w", false, "count words")
	foo func(*bufio.Scanner) (int, error)
)

func countLines(scanner *bufio.Scanner) (count int, err error) {
	for scanner.Scan() {
		count++
		if err = scanner.Err(); err != nil {
			return
		}
	}
	return
}

func countWords(scanner *bufio.Scanner) (count int, err error) {
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count++
		if err = scanner.Err(); err != nil {
			return
		}
	}
	return
}

func countChars(scanner *bufio.Scanner) (count int, err error) {
	for scanner.Scan() {
		count += utf8.RuneCountInString(scanner.Text()) + 1
		if err = scanner.Err(); err != nil {
			return
		}
	}
	return
}

func fooInit() error {
	countFlsgs := 0
	if *linesFlag == true {
		foo = countLines
		countFlsgs++
	}
	if *charsFlag == true {
		foo = countChars
		countFlsgs++
	}
	if *wordsFlag == true || countFlsgs == 0 {
		foo = countWords
		countFlsgs++
	}
	if countFlsgs != 1 {
		return errors.New("should be 1 or 0 flags")
	}
	return nil
}

func processFile(filename string, wg *sync.WaitGroup, result *sync.Map) {
	defer wg.Done()
	file, open_er := os.Open(filename)
	if open_er == nil {
		defer file.Close()
		count, foo_err := foo(bufio.NewScanner(file))
		(*result).Store(filename, fileInfo{count, foo_err})
	} else {
		(*result).Store(filename, fileInfo{0, open_er})
	}
}

type fileInfo struct {
	count int
	err error
}

func main() {
	flag.Parse()
	if er := fooInit(); er != nil {
		fmt.Println(er)
		return
	}
	if len(flag.Args()) == 0 {
		fmt.Println("error: no input files specified:(")
		return
	}
	var result sync.Map
	var wg sync.WaitGroup
	for _, filename := range flag.Args() {
		wg.Add(1)
		go processFile(filename, &wg, &result)
	}
	wg.Wait()
	for _, filename := range flag.Args() {
		var i interface{} = fileInfo{}
		i,_ = result.Load(filename)
		res := fileInfo(i.(fileInfo))
		if res.err == nil {
			fmt.Printf("%d\t%s\n", res.count, filename)
		} else {
			fmt.Printf("%s: %s:(\n", filename, res.err)
		}
	}
}
