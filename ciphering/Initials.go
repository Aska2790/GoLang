/**
 школа Ш++
задание: Взятие инициалов полного имени
автор: Арслан Аннаев
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	name := GetName("Введите полное имя : ")
	init := Initials(name)
	fmt.Println(init)
}

// Ввод ФИО
func GetName(text string) string {
	in := bufio.NewScanner(os.Stdin)
	fmt.Print(text)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}

// обработка строки
func Initials(name string) string {
	var word string

	if strings.HasSuffix(name, "-") { // заменять если имя составное и между стоит тире
		strings.Replace(name, "-", " ", len(name))
	}

	str := strings.Fields(name) // разбивка на отдельные слова

	for i := 0; i < len(str); i++ { // пройтисьм по словам взять перввй символ
		k := string(str[i])
		word += strings.ToUpper(string(k[0])) + "." // перевести вверхний регистр и сконкатенировать
	}
	return word
}
