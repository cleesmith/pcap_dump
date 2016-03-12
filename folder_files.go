package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func listFilesInFolder(filenamePrefix string, path string) (filesList []string, err error) {
	// check if path is valid file or dir
	stat, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	filesList = make([]string, 0)
	if stat.IsDir() {
		files, err := filepath.Glob(path + filenamePrefix)
		if err != nil {
			return nil, err
		}
		filesList = append(filesList, files...)
	} else {
		return nil, errors.New("error: path must be a folder not a file")
	}
	return filesList, nil
}

func main() {
	files, err := listFilesInFolder("beacon_sub_*", "/Users/chrissmith/go/src/github.com/cleesmith/pcap_dump/pub.netsniff-ng.org/pcaps/802.11/info_elements/")
	if err != nil {
		fmt.Printf("err=%v\n", err)
	}
	if len(files) > 0 {
		for _, file := range files {
			fmt.Printf("file=%v\n", filepath.Base(file))
		}
		fmt.Printf("found %v files\n", len(files))
	} else {
		fmt.Printf("no files found (%v)\n", len(files))
	}
	tn := time.Now().UTC()
	fmt.Printf("now=%v\n", fmt.Sprintf("%d-%d-%d", tn.Year(), tn.Month(), tn.Day()))
}
