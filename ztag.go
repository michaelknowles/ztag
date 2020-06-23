package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func help() {
	fmt.Println("Program to tag files")
	fmt.Println("Usage: ztag -file='filename' -type='type' tag1 tag2")
	fmt.Println("Available types: doujin, pic, story, vid")
	os.Exit(1)
}

func main() {
	// Parse command line args
	filePtr := flag.String("file", "", "The file to tag")
	typePtr := flag.String("type", "", "Type of file")

	flag.Parse()

	// Check required args
	fileName := *filePtr
	if len(fileName) == 0 {
		help()
	}
	fileType := *typePtr
	if len(fileType) == 0 {
		help()
	}
	tags := flag.Args()
	if len(tags) == 0 {
		help()
	}

	// Validate file exists
	// TODO Move into function and create unit tests
	if !fileExists(fileName) {
		panic(fmt.Sprintf("File %s doesn't exist", fileName))
	}
	// Get absolute path of fileName
	fileNameAbs, err := filepath.Abs(fileName)
	if err != nil {
		panic(fmt.Sprintf("Couldn't get absolute path of %s", fileNameAbs))
	}
	// Set fileName to its Base (remove any path info to just get file name alone)
	fileName = filepath.Base(fileName)

	// Validate type
	// TODO Move into function and create unit tests
	fileType = strings.ToLower(fileType)
	allowedTypes := []string{
		"doujin",
		"pic",
		"story",
		"vid",
	}
	validType := false
	for _, t := range allowedTypes {
		if t == fileType {
			validType = true
		}
	}
	if validType == false {
		panic(fmt.Sprintf("%s is not a valid type", fileType))
	}

	// Validate ZDIR env var
	// TODO Move into function and create unit tests
	zDir := os.Getenv("ZDIR")
	if len(zDir) == 0 {
		panic("ZDIR is not set")
	}
	if !dirExists(zDir) {
		panic(fmt.Sprintf("ZDIR %s doesn't exist", zDir))
	}

	// Create symlinks for tags
	for _, tag := range tags {
		// Lowercase tag
		tag = strings.ToLower(tag)
		// Create the tag if it doesn't already exist
		tagPath := path.Join(zDir, fileType, tag)
		if !dirExists(tagPath) {
			err := os.Mkdir(tagPath, 0755)
			if err != nil {
				panic(err)
			}
		}
		// Create the symlink under the tag
		symlink := path.Join(tagPath, fileName)
		err := os.Symlink(fileNameAbs, symlink)
		if err != nil {
			panic(err)
		}
	}
}

// fileExists checks if a given file exists
func fileExists(fileName string) bool {
	filePath := path.Clean(fileName)
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// dirExists checks if a given directory exists
func dirExists(dirName string) bool {
	dirPath := path.Clean(dirName)
	info, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}
