package main

import (
	"strings"
	"path/filepath"
	"fmt"
)

func main() {
	var files string ="input.txt"
	suffixes := []string{".txt", ".jpg", ".png"}
	channel_1 := source(files)
	channel_2 := filterSuffixes(suffixes, channel_1)
	sink(channel_2)
}


func source ( files[] string)<-chan string{
	out := make (chan string, 1000)
	go func(){
		for _, filename := range files{
			out <- filename
		}
		close(out)
	}()
	return out
}


func filterSuffixes(suffixes[] string, in<-chan string)<-chan string{
	out:=make(chan string, cap(in))
	go func(){
		for _, filename := range in{
			if len(suffixes)==0{
				out<-filename
				continue
			}

			ext:= strings.ToLower(filepath.Ext(filename))
			for _,suffix:= range suffixes{
				if ext == suffix{
					out<-filename
					break
				}
			}
		}
		close(out)
	}()
	return out
}

func sink(in <-chan string ){
	for filename:= range in{
		fmt.Println(filename)
	}
}



