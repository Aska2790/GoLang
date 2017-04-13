package speller_func

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var dictionary = make([]string, 0)

// read text from file
func ReadText(path string) []string {
	var res []string

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	buf := make([]string, 1)
	eof := false
	for i := 1; !eof; i++ {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			err = nil
			eof = true
		}
		buf = append(buf, string(line))
	}

	for i := range buf {
		dic := strings.Fields(buf[i])
		for j := range dic {
			res = append(res, dic[j])
		}
	}
	return res
}

func Speller(dic, text []string) int {
	var misspell_count int = 0
	dictionary = dic

	for i := 0; i < len(text); i++ {
		if IsCorrectWord(text[i]) {
			fmt.Println("misspell: ", text[i])
			misspell_count++
		}
	}
	return misspell_count
}

func IsCorrectWord(word string) bool {
	temp := strings.ToLower(word)
	temp = Trimmer(temp)
	if _, err := strconv.Atoi(temp); err != nil {
		for i := 0; i < len(dictionary); i++ {
			if strings.Compare(dictionary[i], temp) == 0 {
				return false
			}
		}
		return true
	}
	return false
}

func Report(misspell_count int, text, dictionary []string) {
	fmt.Println("Word count in Text : ", len(text))
	fmt.Println("Word Misspelled :", misspell_count)
	fmt.Println("Word count in Dictionary :", len(dictionary))
}

func Trimmer(word string) string {
	if strings.Index(word, ".") != -1 {
		s := strings.Split(word, ".")
		word = s[0]
	}
	if strings.Index(word, ",") != -1 {
		s := strings.Split(word, ",")
		word = s[0]
	}
	if strings.Index(word, "’") != -1 {
		s := strings.Split(word, "’")
		word = s[0]
	}
	return word
}
