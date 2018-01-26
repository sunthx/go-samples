package main

import (
	"strings"
	"path/filepath"
	"fmt"
)

func source(files []string) <-chan string {
	out:=make(chan string,1000)
	go func() {
		for _,filename := range files{
			out <- filename
		}
		close(out)

	}()

	return out
}

func filterSuffixes(suffixes []string,in <-chan string) <-chan string{
	out := make(chan string,cap(in))
	go func() {
		for fileName :=range in{
			if len(suffixes) == 0{
				out <- fileName
				continue
			}

			ext := strings.ToLower(filepath.Ext(fileName))
			for _,suffix := range suffixes{
				if ext == suffix{
					out <- fileName
					break
				}
			}

		}

		close(out)
	}()

	return out
}

func sink(in <-chan string)  {
	for filename := range in{
		fmt.Println(filename)
	}
}

func main(){
	files := []string{"1.txt","2.txt","3.mp4","4.mp4", "5.flv"}
	suffixes:=[]string{".txt", ".flv"}

	sink(filterSuffixes(suffixes,source(files)))

}