package main

import (
	"./function"

	"sync"
	"time"
)

var (
	dictionary =  make ([]string, 1)
	text =  make ([]string, 1)
	misspell_count =0
)

const(
	dictionary_path = "D:\\dictionary.txt"
	text_path = "D:\\in.txt"

)

func main() {

	w := sync.WaitGroup{}
	w.Add(2)
	go func(w *sync.WaitGroup) {
		function.ReadText(dictionary_path, &dictionary)
		function.ReadText(text_path, &text)
		w.Done()
	}(&w)

	go func(w *sync.WaitGroup) {
		function.Speller(dictionary, text,&misspell_count)
		w.Done()
	}(&w)

	time.Sleep(2*time.Second)
	defer function.Report(&misspell_count,text, dictionary)
	w.Wait()
}
