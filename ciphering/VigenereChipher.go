/**
 школа Ш++
задание: Шифр Видженера
автор: Арслан Аннаев
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

const (
	UP_A      = 65
	LOW_a     = 97
	LET_COUNT = 26
)

var (
	alphabet = make(map[string]int)
	str      = "abcdefghijklmnopqrstuvwxyz"
)

func main() {
	key := GetStr("Enter the key")
	text := GetStr("Enter the Text ")
	fmt.Println(Encoder(text, key))
}

// инициализаци мапы для ключа
func InitMap() {
	for i := 0; i < LET_COUNT; i++ {
		alphabet[string(str[i])] = i + 1
	}
}

// шифр видженера
func Encoder(text, key string) string {
	InitMap()

	var result = []rune(text)
	iterator := 0

	for i := 0; i < len(text); i++ {
		if unicode.IsLetter(rune(text[i])) {
			if iterator >= len(key) {
				iterator = 0
			}
			result[i] = Coders(text[i], byte(key[iterator]))
			iterator++
		}
	}
	bs := []byte(string(result))
	return string(bs)
}

// кодек для преобразования символа со смещением
func Coders(UserLetter, KeysLetter byte) rune {
	k := strings.ToLower(string(KeysLetter))
	var key int = alphabet[string(k)]
	var s rune

	if unicode.IsUpper(rune(UserLetter)) { // если она верхнего регистра
		s = rune((((UserLetter - UP_A) + byte(key)) % LET_COUNT) + UP_A) // возвращаем букву со смещением

	} else if unicode.IsLower(rune(UserLetter)) {
		s = rune((((UserLetter - LOW_a) + byte(key)) % LET_COUNT) + LOW_a) // буква ниж рег со смещением
	}
	return s
}

// Get string to encrypt
func GetStr(text string) string {
	in := bufio.NewScanner(os.Stdin)
	fmt.Println(text)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}
