package main

import (
	"./function"
	"sync"
	"time"
)

var (
	dictionary     = make([]string, 1)
	text           = make([]string, 1)
	misspell_count = 0
)

const (
	dictionary_path = "D:\\dictionary.txt"
	text_path       = "D:\\in.txt"
)

func main() {
	n := make(chan bool)
	w := sync.WaitGroup{}
	w.Add(3)

	go func(w *sync.WaitGroup) {
		function.ReadText(dictionary_path, &dictionary)
		w.Done()
		n <- true
	}(&w)

	go func(w *sync.WaitGroup) {
		function.ReadText(text_path, &text)
		w.Done()
	}(&w)

	time.Sleep(1 * time.Second)
	go func(w *sync.WaitGroup) {
		<-n
		function.Speller(dictionary, text, &misspell_count)
		w.Done()
	}(&w)

	defer function.Report(&misspell_count, text, dictionary)
	w.Wait()
}
