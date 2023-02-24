package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func rotateLog(logFile, archiveDir string) {
	info, err := os.Stat(logFile)
	if err != nil {
		log.Fatalf("error reading log file: %v", err)
	}
	now := time.Now()
	if now.Format("2006-01-02") == info.ModTime().Format("2006-01-02") {
		log.Printf("log file %s was modified today, skipping rotation", logFile)
		return
	}
	timestamp := info.ModTime().Unix()
	newLogFile := strings.TrimSuffix(logFile, filepath.Ext(logFile))
	newLogFile = fmt.Sprintf("%s_%d.tag.gz", filepath.Join(archiveDir, filepath.Base(newLogFile)), timestamp)
	log.Printf("rotating log file %s to %s", logFile, newLogFile)
	file, err := os.Open(logFile)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}
	defer file.Close()
	gzFile, err := os.Create(newLogFile)
	if err != nil {
		log.Fatalf("error creating gzipped archive: %v", err)
	}
	defer gzFile.Close()
	gzw := gzip.NewWriter(gzFile)
	if _, err := file.Seek(0, 0); err != nil {
		log.Fatalf("error seeking log file: %v", err)
	}
	if _, err := io.Copy(gzw, file); err != nil {
		log.Fatalf("error copying log data to gzipped archive: %v", err)
	}
	if err := gzw.Close(); err != nil {
		log.Fatalf("error closing gzipped archive: %v", err)
	}
	if err := os.Truncate(logFile, 0); err != nil {
		log.Fatalf("error truncating log file: %v", err)
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("usage: %s -a <archive directory> <log file 1> <log file 2> ... <log file n>", os.Args[0])
	}

	var archiveDir string
	var logFiles []string

	if os.Args[1] == "-a" {
		if len(os.Args) < 4 {
			log.Fatalf("usage: %s -a <archive directory> <log file 1> <log file 2> ... <log file n>", os.Args[0])
		}
		archiveDir = os.Args[2]
		logFiles = os.Args[3:]
	} else {
		archiveDir = filepath.Dir(os.Args[1])
		logFiles = os.Args[1:]
	}

	var wg sync.WaitGroup
	for _, logFile := range logFiles {
		wg.Add(1)
		go func(logFile string) {
			defer wg.Done()
			rotateLog(logFile, archiveDir)
		}(logFile)
	}
	wg.Wait()
}
