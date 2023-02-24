package libjx

import (
	"os"
	"bufio"
	"strings"
)

func stripJsonComments(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var newBytes []byte
	var commentStarted = false
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "//") {
			line = line[:strings.Index(line, "//")]
		}
		if strings.Contains(line, "/*") {
			commentStarted = true
			if !strings.Contains(line, "*/") {
				continue
			}
			line = line[strings.Index(line, "*/")+2:]
			commentStarted = false
		}
		if commentStarted {
			continue
		}
		newBytes = append(newBytes, []byte(line+"\n")...)
	}
	return newBytes, nil
}
