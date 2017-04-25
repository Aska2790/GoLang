package main

import (
	"fmt"
	"os"
	"runtime"

	"log"
	"time"
)

var (
	inPath  = "D:\\1.mp4" // read file that is on this path
	outPath = "D:\\2.mp4" // write file that is on this path

)

//======================================================================================================================
func main() {

	start := time.Now() // Now returns the current local time.

	kernels_count := runtime.NumCPU() // NumCPU returns the number of logical CPUs usable by the current process.
	runtime.GOMAXPROCS(kernels_count) // GOMAXPROCS sets the maximum number of CPUs that can be executing

	file, err_open := os.Open(inPath) // Open opens the named file for reading. If successful, methods on
	check_error("file couldn`t open ", err_open)
	defer file.Close() // Close closes the File, rendering it unusable for I/O.

	out_file, err_create := os.Create(outPath) // Create creates the named file with mode 0666
	check_error("file could`n create ", err_create)
	defer out_file.Close()

	fileSize := getFileSize(file) // get file size

	// split the file into parts to set offset for read() gorutines
	sliceWithSize := splitFileSize(fileSize, kernels_count)

	// Create buffers whose number is equal to the number of kernels
	buf := createBuffer(kernels_count, sliceWithSize)

	// loop for reading
	for i := 0; i < kernels_count; i++ {
		offset := int64(i) * sliceWithSize[i]
		file.ReadAt(buf[i], offset)
	}

	// loop for writing
	for i := 0; i < kernels_count; i++ {
		offset := int64(i) * sliceWithSize[i]
		out_file.WriteAt(buf[i], offset)
	}

	result := time.Since(start)
	fmt.Println("time to read and write", result)
}

//======================================================================================================================
// get file size
func getFileSize(file *os.File) int64 {
	// get info about the file
	fileInfo, err_getInfo := file.Stat() // Stat returns the FileInfo structure describing file.
	check_error("can`t get file size", err_getInfo)

	fileSize := fileInfo.Size() // length in bytes for regular files
	return fileSize
}

//======================================================================================================================
// check all errors
func check_error(msq string, err error) {
	if err != nil {
		log.Fatal(msq, " ", err)
	}
}

//======================================================================================================================
// Create buffers whose number is equal to the number of kernels
func createBuffer(concurrency_count int, parts []int64) [][]byte {
	slice := make([][]byte, concurrency_count)
	for i := 0; i < concurrency_count; i++ {
		slice[i] = make([]byte, parts[i])
	}
	return slice
}

//======================================================================================================================
// split the file into parts to set offset for read() gorutines
func splitFileSize(fileSize int64, concurrency int) []int64 {
	slice := make([]int64, concurrency)
	part := fileSize / int64(concurrency)

	for i := 0; i < concurrency; i++ {
		if fileSize > part {
			slice[i] = part
		} else {
			slice[i] = fileSize
		}
		fileSize -= part
	}
	return slice
}

//======================================================================================================================
