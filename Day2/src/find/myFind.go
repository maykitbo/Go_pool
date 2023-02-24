package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"errors"
)

type findOptions struct {
	showSymlinks    bool
	showDirectories bool
	showFiles       bool
	fileExtension   string
}

func handleFlags() (findOptions, error) {
	var options findOptions
	flag.BoolVar(&options.showSymlinks, "sl", false, "Show only symbolic links")
	flag.BoolVar(&options.showDirectories, "d", false, "Show only directories")
	flag.BoolVar(&options.showFiles, "f", false, "Show only files")
	flag.StringVar(&options.fileExtension, "ext", "", "Show only files with the specified extension")
	flag.Parse()
	if options.fileExtension != "" && options.showFiles == false {
		return options, errors.New("-ext without -f")
	}
	if options.showSymlinks == false && options.showDirectories == false && options.showFiles == false {
		options.showSymlinks, options.showDirectories, options.showFiles = true, true, true
	}
	return options, nil
}

func myFind(root string, options findOptions) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			if os.IsPermission(err) {
				return nil
			}
			return err
		}
		if options.showSymlinks && info.Mode()&os.ModeSymlink != 0 {
			link,err := os.Readlink(path)
			if err != nil {
				return err
			}
			if _, err := os.Stat(link); os.IsNotExist(err) {
				fmt.Printf("%s -> [broken]\n", path)
			} else {
				if os.IsPermission(err) {
					return nil
				}
				fmt.Printf("%s -> %s\n", path, link)
			}
		} else if options.showDirectories && info.IsDir() {
			if path != "." { fmt.Println(path) }
		} else if options.showFiles && !info.IsDir() {
			if options.fileExtension == "" || strings.HasSuffix(path, "."+options.fileExtension) {
				fmt.Println(path)
			}
		}
		return nil
	})
}

func main() {
	options, er := handleFlags()
	if er != nil {
		fmt.Println(er)
		return
	}
	root := flag.Arg(0)
	err := myFind(root, options)
	if err != nil {
		fmt.Printf("%q: no such file or directory:(\n", root)
	}
}




