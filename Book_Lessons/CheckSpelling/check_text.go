
package main

import (
	"./functions"
)


const(
	dictionary_path = "D:\\dictionary.txt"
	text_path = "D:\\in.txt"

)
func main() {
	dictionary := functions.ReadText(dictionary_path)
	text := functions.ReadText(text_path)
	count := functions.Speller(dictionary, text)
	functions.Report(count, text, dictionary)
}





