package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var (
	in_path  = "D:\\1.mp4"
	out_path = "D:\\2.mp4"
)

//======================================================================================================================
func main() {
	meausure_ioutil()
	meausure_os_Read()
	meausure_os_ReadAt()
}

//======================================================================================================================
// function for check all errors
func check_error(msg string, err error) {
	if err != nil {
		log.Fatal(msg, " ", err)
	}
}

//======================================================================================================================
// measure file reading time with ioutil package
func meausure_ioutil() {
	start := time.Now() // Now returns the current local time.

	// read data from file
	data, err_read := ioutil.ReadFile(in_path) // ReadFile reads the file named by filename and returns the contents.
	check_error("couldn`t read ", err_read)

	// write data to file
	err_write := ioutil.WriteFile(out_path, data, 0644) // WriteFile writes data to a file named by filename.
	check_error("couldn`t write ", err_write)

	result := time.Since(start) // Since returns the time elapsed since t.
	fmt.Println("time  ioutil.ReadFile() ", result)
}

//======================================================================================================================
// measure file reading time with os package
func meausure_os_ReadAt() {
	start := time.Now() // Now returns the current local time.

	file, err_open := os.Open(in_path) // Open the named file for reading
	check_error("couldn`t open", err_open)
	defer file.Close()

	out_file, err_create := os.Create(out_path) // Create  the named file with mode 0666
	check_error("couldn`t create ", err_create)
	defer out_file.Close()

	fileInfo, _ := file.Stat()        // Stat returns the FileInfo structure describing file.
	file_size := fileInfo.Size()      // length in bytes for regular files
	buffer := make([]byte, file_size) // and create buffer with file size

	_, err_read := file.ReadAt(buffer, 0) // Read  up to len(b) bytes from the File.
	check_error("couldn`t read ", err_read)

	_, err_write := out_file.WriteAt(buffer, 0) // Write  len(b) bytes to the File.
	check_error("couldn`t write ", err_write)

	result := time.Since(start) // Since returns the time elapsed since t.
	fmt.Println("time   os.ReadAt() ", result)
}

//======================================================================================================================
// measure file reading time with os package
func meausure_os_Read() {
	start := time.Now()                // Now returns the current local time.
	file, err_open := os.Open(in_path) // Open  the named file for reading
	check_error("couldn`t open", err_open)
	defer file.Close()

	out_file, err_create := os.Create(out_path) // Create  the named file with mode 0666
	check_error("couldn`t create ", err_create)
	defer out_file.Close()

	fileInfo, _ := file.Stat()        // Stat returns the FileInfo structure describing file.
	file_size := fileInfo.Size()      // length in bytes for regular files
	buffer := make([]byte, file_size) // and create buffer with file size

	_, err_read := file.Read(buffer) // Read  up to len(b) bytes from the File.
	check_error("couldn`t read ", err_read)

	_, err_write := out_file.Write(buffer) // Write  len(b) bytes to the File.
	check_error("couldn`t write ", err_write)

	result := time.Since(start) // Since returns the time elapsed since t.
	fmt.Println("time  os.Read() ", result)
}

//======================================================================================================================
