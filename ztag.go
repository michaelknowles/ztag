package main

import (
	"flag"
	"fmt"
	"os"
	"path"
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
	if !fileExists(fileName) {
		panic(fmt.Sprintf("File %s doesn't exist", fileName))
	}

	// Lowercase filetype
	fileType = strings.ToLower(fileType)
	// Validate type
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
		err := os.Symlink(fileName, symlink)
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