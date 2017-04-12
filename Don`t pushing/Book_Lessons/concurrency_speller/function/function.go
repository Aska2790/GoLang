package function

import (
	"os"
	"log"
	"bufio"
	"io"
	"strings"
	"fmt"
	"strconv"
)

var dictionary  = make([]string,0)

// read text from file
func ReadText(path string, buffer *[]string){


	file, err := os.Open(path)
	if err!=nil{
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	buf := make([]string,1)
	eof:=false
	for i:=1; !eof; i++{
		line, err:= reader.ReadString('\n')
		if err ==io.EOF{
			err = nil
			eof = true
		}
		buf = append(buf,string(line ))
	}

	for i:=range buf {
		dic := strings.Fields(buf[i])
		for j:=range dic{
			*buffer = append(*buffer, dic[j])
		}
	}
}


func Speller(dic, text []string, misspell_count *int) {
	dictionary = dic

	for i := 0; i < len(text); i++ {

		if IsCorrectWord(text[i]){
			*misspell_count++
		}

	}
}



func IsCorrectWord(word string) bool {

	temp := strings.ToLower(word)
	temp = Trimmer(temp)
	if _, err:= strconv.Atoi(temp); err== nil {
		return false
	}

	for i := 0; i < len(dictionary); i++ {
		if strings.Compare(dictionary[i], temp) == 0 {
			return false
		}
	}
	return true
}

func Report(misspell_count *int, text, dictionary []string){
	fmt.Println("Word count in Text : ",len(text) )
	fmt.Println("Word Misspelled :",*misspell_count  )
	fmt.Println("Word count in Dictionary :", len(dictionary)  )
}

func Trimmer(word string) string{
	if strings.Index(word, ".") != -1{
		s:= strings.Split(word, ".")
		word =s[0]
	}
	if strings.Index(word, ",") != -1{
		s:= strings.Split(word, ",")
		word = s[0]
	}
	if strings.Index(word, "’") != -1{
		s:= strings.Split(word, "’")
		word = s[0]
	}


	return word
}
