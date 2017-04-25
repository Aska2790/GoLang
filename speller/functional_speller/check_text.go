package main

import (
	"./speller_func"
)

const (
	dictionary_path = "D:\\dictionary.txt"
	text_path       = "D:\\in.txt"
)

func main() {
	dictionary := speller_func.ReadText(dictionary_path)
	text := speller_func.ReadText(text_path)
	count := speller_func.Speller(dictionary, text)
	speller_func.Report(count, text, dictionary)
}
