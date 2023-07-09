package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func GetFileSize(filename string) int64 {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		panic(err)
	}
	return fileInfo.Size()
}

func AppendFile(sourceFilename string, appendDestinationFilename string) {
	content, err := ioutil.ReadFile(sourceFilename)
	if err != nil {
		panic(err)
	}

	destinationFD, err := os.OpenFile(appendDestinationFilename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	_, err = destinationFD.Write(content)
	if err != nil {
		panic(err)
	}
	destinationFD.Close()
}

func SubstractFile(appendSourceFilename string, fileStartPosition int64, fileEndPosition int64, newFilename string) {
	// Open the destination file for writing
	destinationFD, err := os.Create(newFilename)
	if err != nil {
		panic(err)
	}

	// Open the source file for reading
	sourceFD, err := os.Open(appendSourceFilename)
	if err != nil {
		panic(err)
	}

	// Seek to the start byte position
	_, err = sourceFD.Seek(int64(fileStartPosition), io.SeekStart)
	if err != nil {
		panic(err)
	}

	// Read and write the specified range of bytes
	buffer := make([]byte, fileEndPosition-fileStartPosition+1)
	_, err = io.ReadFull(sourceFD, buffer)
	if err != nil {
		panic(err)
	}
	_, err = destinationFD.Write(buffer)
	if err != nil {
		panic(err)
	}
	sourceFD.Close()
	destinationFD.Close()
}

func main() {
	//filename := "/Users/ulises/workspace/container-experiments/object-storage-app/src/examples/audiofile.m4a"
	appendFilename := "/Users/ulises/workspace/container-experiments/object-storage-app/src/examples/appendDestination.obj"
	newFilename := "/Users/ulises/workspace/container-experiments/object-storage-app/src/examples/new-audiofile.m4a"
	//fileStartPosition := GetFileSize(appendFilename)
	var fileStartPosition int64
	var fileEndPosition int64
	fileStartPosition = 1507696
	//AppendFile(filename, appendFilename)
	//fileEndPosition := GetFileSize(appendFilename)
	fileEndPosition = 3015392
	//fmt.Printf("Filename: '%s' appended in position: %d\n", filename, fileStartPosition, fileEndPosition)
	SubstractFile(appendFilename, fileStartPosition, fileEndPosition, newFilename)
	fmt.Printf("New File Created: %s\n", newFilename)
}
