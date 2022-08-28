package main

import (
	"bufio"
	"fmt"
	"go-reloaded/src"
	"os"
	"regexp"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Incorrect number of arguments")
		return
	}
	file1 := args[0]
	file2 := args[1]

	if file2[len(file2)-3:] == ".go" {
		fmt.Println("can not use .go files as name")
		return
	}

	f, err := os.Open(file1) // opening file
	if err != nil {
		fmt.Print(file1)
		fmt.Println(" file does not exist")
		return
	}
	defer f.Close()

	d := bufio.NewScanner(f)

	var tmp []byte
	var ans []byte
	arg := src.Reload{} // creating new struct
	for d.Scan() {
		arg.Data = d.Text()

		arg.Articles() // formatting data
		arg.Convert()

		spaces := regexp.MustCompile(`(\s*)((\.\.\.)|(!\?)|(\.)|(,)|(!)|(\?)|(:)|(;))(\s*)`)
		arg.Data = spaces.ReplaceAllString(arg.Data, "$2 ")

		arg.Apostrophe()

		if len(arg.Data) != 0 && arg.Data[len(arg.Data)-1] == ' ' {
			arg.Data = arg.Data[:len(arg.Data)-1]
		}

		tmp = []byte(arg.Data)
		tmp = append(tmp, '\n')
		ans = append(ans, tmp...)
	}

	err = os.WriteFile(file2, ans, 0644) // creating file with corrected data inside
	if err != nil {
		panic(err)
	}
}
