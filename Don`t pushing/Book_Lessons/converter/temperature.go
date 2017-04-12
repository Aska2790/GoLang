package main

import "fmt"

var celsiuse  float64

func main() {
	fmt.Print("Температура в Цельсиях :")
	fmt.Scanf("%f", &celsiuse)
	result :=  (celsiuse*9)/5 +32
	fmt.Println("Температура в Фаренгейтах :", result)
}
