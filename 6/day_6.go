package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func gotMarker(idx int, str string) bool {
	if str[idx] == str[idx-1] {
		return false
	}
	if str[idx] == str[idx-2] {
		return false
	}
	if str[idx] == str[idx-3] {
		return false
	}
	if str[idx-1] == str[idx-2] {
		return false
	}
	if str[idx-1] == str[idx-3] {
		return false
	}
	if str[idx-2] == str[idx-3] {
		return false
	}

	return true
}

func main() {
	day := 6

	strbytes, err := os.ReadFile("input" + fmt.Sprint(day) + ".txt")
	check(err)

	str := string(strbytes)

	l := utf8.RuneCountInString(str)

	for i := 3; i < l; i++ {
		if gotMarker(i, str) {
			fmt.Println(i + 1)
			break
		}
	}
}
