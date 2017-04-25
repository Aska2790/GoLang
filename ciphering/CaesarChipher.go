/**
 школа Ш++
задание: Шифр Цезаря
автор: Арслан Аннаев
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

const (
	UPPER_A      = 65
	LOWER_a      = 97
	LETTER_COUNT = 26
)

func main() {
	temp := GetString()
	var key byte = GetKey()
	fmt.Println(Encoders(temp, key))
	fmt.Println(Decoder(Encoders(temp, key), key))

}

// преобразование в строку со смещением
func Coder(text string, key byte) string {
	var result = []rune(text)

	for s := range text {
		if unicode.IsLetter(rune(text[s])) { // если буква
			if unicode.IsUpper(rune(text[s])) { // если она верхнего регистра
				result[s] = rune((((text[s] - UPPER_A) + byte(key)) % LETTER_COUNT) + UPPER_A) // возвращаем букву со смещением

			} else if unicode.IsLower(rune(text[s])) {
				result[s] = rune((((text[s] - LOWER_a) + byte(key)) % LETTER_COUNT) + LOWER_a) // буква ниж рег со смещением
			}
		}
	}

	bs := []byte(string(result)) // приведение к типу байт
	return string(bs)
}

// шифровка текста
func Encoders(text string, key byte) string {
	return Coder(text, key)
}

// Расшифровка текста
func Decoder(text string, key byte) string {
	return Coder(text, -key)
}

// получение строки для шифрации
func GetString() string {
	in := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter the text")
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()

}

// Get key from user
func GetKey() byte {
	var temp int
	fmt.Println("Please enter the key for encod text")
	fmt.Scanf("%d", &temp)
	return byte(temp)
}
