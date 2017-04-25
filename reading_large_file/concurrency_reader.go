// determine the size of file
// split between gorutines
// create buffers whose number equal to the number of gorutines
// create several gorutines whose number equal
// to the number of processor on the machine
// write all to the file with gorutines whose number equal
// to the number of processor on the machine
package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

var (
	in_path       = "D:\\1.mp4"   // read file that is on this path
	out_path      = "D:\\out.mp4" // write file that is on this path
	complete      = make(chan bool)
	index         = make(chan int)
	files         = make([]*os.File, kernels_count)
	outs          = make([]*os.File, kernels_count)
	kernels_count = runtime.NumCPU()
)

//======================================================================================================================
func main() {
	start := time.Now() // Now returns the current local time.
	// use all available kernels
	runtime.GOMAXPROCS(kernels_count) // GOMAXPROCS sets the maximum number of CPUs that can be executing
	initialization()

	fileSize := getFileSize() // get the file size

	// split the file into parts to set offset for read() gorutines
	part_slice := splitFileSize(fileSize, kernels_count)

	// Create buffers whose number is equal to the number of kernels
	buf := createBuffer(kernels_count, part_slice)

	// loop for reading
	for i := 0; i < kernels_count; i++ {
		go read(buf, part_slice, i, files[i])
	}

	//loop for writing
	for i := 0; i < kernels_count; i++ {
		go write(buf, part_slice, outs[i])
	}
	//waiting for completion
	for i := 0; i < kernels_count; i++ {
		<-complete
	}

	finish := time.Since(start)
	fmt.Println("time ", finish)
}

//======================================================================================================================
// init file handlers
func initialization() {
	for i := 0; i < kernels_count; i++ {
		files[i], _ = os.Open(in_path)   // Open opens the named file for reading.
		outs[i], _ = os.Create(out_path) // Create creates the named file with mode 0666
	}
}

//======================================================================================================================
// get file size
func getFileSize() int64 {
	// get info about the file
	fileInfo, err_getInfo := files[0].Stat()
	check_error("can`t get file size", err_getInfo)

	fileSize := fileInfo.Size()
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
	BufferSlice := make([][]byte, concurrency_count)
	for i := 0; i < concurrency_count; i++ {
		BufferSlice[i] = make([]byte, parts[i])
	}
	return BufferSlice
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
// read file
func read(buf [][]byte, offsets []int64, i int, file *os.File) {
	// ReadAt reads len(b) bytes from the File starting at byte offset off.
	file.ReadAt(buf[i], int64(i)*offsets[i])
	//fmt.Println("buf len : ", len(buf[i]), int64(i)*offsets[i], "num :", i )
	index <- i
}

//======================================================================================================================
// write file
func write(buf [][]byte, offsets []int64, out_file *os.File) {
	i := <-index
	// WriteAt writes len(b) bytes to the File starting at byte offset off.
	out_file.WriteAt(buf[i], int64(i)*offsets[i])
	//fmt.Println("buf len : ", len(buf[i]), int64(i)*offsets[i], "num :", i )
	complete <- true
}

//======================================================================================================================
