package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	files:=[]string{"Test.conf","util.go","Makefile","misc.go","main.go"}
	fmt.Printf("Unsorted:    %q\n",files);

	//Sort
	sort.Strings(files)
	fmt.Printf("Underlying bytes: %q\n",files)

	SortFoldedStrings(files)
	fmt.Printf("Case insensitive: %q\n",files)

	sort.Strings(files)

	//Search sort.Search() 二分搜索算法
	target := "util.go"

	//普通的线性搜索
	for i,file := range files{
		if file == target{
			fmt.Printf("found \"%s\" at files[%d]\n",file,i)
			break
		}
	}

	//使用sort.search()方法
	i := sort.Search(len(files), func(i int) bool {
		return files[i] >= target
	})
	if i< len(files) && files[i] == target{
		fmt.Printf("found \"%s\" at files[%d]\n",files[i],i)
	}



}

func SortFoldedStrings(slice []string)  {
	sort.Sort(FoldedStrings(slice))
}

type FoldedStrings []string

func (slice FoldedStrings) Len() int {return len(slice)}

func (slice FoldedStrings) Less(i,j int) bool  {
	return strings.ToLower(slice[i]) < strings.ToLower(slice[i])
}

func (slice FoldedStrings) Swap(i,j int)  {
	slice[i],slice[j] = slice[j],slice[i]
}