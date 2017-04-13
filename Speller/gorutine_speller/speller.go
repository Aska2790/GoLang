package main

import (
	"./speller_func"
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
		speller_func.ReadText(dictionary_path, &dictionary)
		w.Done()
		n <- true
	}(&w)

	go func(w *sync.WaitGroup) {
		speller_func.ReadText(text_path, &text)
		w.Done()
	}(&w)

	time.Sleep(1 * time.Second)
	go func(w *sync.WaitGroup) {
		<-n
		speller_func.Speller(dictionary, text, &misspell_count)
		w.Done()
	}(&w)

	defer speller_func.Report(&misspell_count, text, dictionary)
	w.Wait()
}
