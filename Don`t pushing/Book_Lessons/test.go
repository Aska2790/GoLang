package main

import (
	"fmt"
	"strings"
)

func main() {

	s := Trimmer("Hello.,")
	fmt.Println(len(s)," = ", s)

}


func Trimmer(word string) string{
	if strings.Index(word, string('.')) != -1{
		s:= strings.Split(word, ".")
		word =s[0]
	}
	if strings.Index(word, string(',')) != -1{
		s:= strings.Split(word, ",")
		word = s[0]
	}
	return word
}
