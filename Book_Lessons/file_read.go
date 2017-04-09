package main

import (
	"bufio"
	"os"
	"io"
	"log"
	"sync"
	"fmt"
	"time"
	"strings"
)

var(
	dictionary     = make([]string,0)
	text 	       = make([]string,0)
	misspell_count = 0
	misspell       = make(chan bool)
	check 	       = make(chan string)
	syncr       	= make(chan bool)
)

const(
	dictionary_path = "D:\\dictionary.txt"
	text_path = "D:\\in.txt"
	concurrency = 100
)
func main() {
	var w sync.WaitGroup




	go func(w *sync.WaitGroup){
		w.Add(1)
		dictionary =  ReadText(dictionary_path)
		w.Done()
	}(&w)
	go func (w *sync.WaitGroup){
		w.Add(1)
		text = ReadText(text_path)
		w.Done()
		syncr <- true

	}(&w)


	go func (w *sync.WaitGroup){
		w.Add(1)
		Speller()
		w.Done()
	}(&w)





	time.Sleep(3*time.Second)
	w.Wait()
	Report()

}


func ReadText(path string)(string){
	var str string
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
		buf = append(buf,line )

	}
	for i:= range buf{
		dic:= strings.Fields(buf[i])
		for j:= range dic{
			str += dic[i]
		}
	}
	return str
}



func Speller(){
	if <-syncr{
		for i := 0; i < len(text); i++ {
			go IsCorrectWord(line[j])
			go Counter()
		}
	}
}


func IsCorrectWord(word string){

	for i := range dictionary{
		fmt.Println(word)
		if word != dictionary[i]{
			misspell <- true
		}
	}
	misspell<-false
}


func Counter(){
	if <-misspell{
		fmt.Println("Yes")
		misspell_count++
	}
	fmt.Println("No")
}

func Report(){
	fmt.Println("Word count in Text : ",len(text) )
	fmt.Println("Word Misspelled :", misspell_count  )
	fmt.Println("Word count in Dictionary :", len(dictionary)  )
}