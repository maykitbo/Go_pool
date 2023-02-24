package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var addLogs bool
	flag.BoolVar(&addLogs, "add", false, "path to log file to add logs to")
	flag.Parse()

	var f *os.File
	var err error
	if addLogs == false {
		f, err = os.Create(flag.Args()[0])
		if err != nil {
			fmt.Printf("error creating log file: %s\n", err)
			return
		}
	} else {
		f, err = os.OpenFile(flag.Args()[0], os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("error opening log file: %s\n", err)
			return
		}
	}
	defer f.Close()

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		randomDay := time.Now().AddDate(0, 0, rand.Intn(3)-3)
		logTime := time.Date(randomDay.Year(), randomDay.Month(), randomDay.Day(),
			rand.Intn(24), rand.Intn(60), rand.Intn(60), 0, randomDay.Location())
		logMessage := fmt.Sprintf("%s: This is random log %d\n", logTime.Format(time.RFC3339), i)
		if _, err = f.WriteString(logMessage); err != nil {
			fmt.Printf("error writing to log file: %s\n", err)
			return
		}
	}
}


