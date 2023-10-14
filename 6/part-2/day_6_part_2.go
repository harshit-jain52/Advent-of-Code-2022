package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

type void struct{}

var blank void

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func gotMarker(idx int, str string) bool {
	m := make(map[byte]void)

	for i := idx - 13; i <= idx; i++ {
		m[str[i]] = blank
	}

	if len(m) == 14 {
		return true
	}

	return false
}

func main() {
	day := 6

	strbytes, err := os.ReadFile("../input" + fmt.Sprint(day) + ".txt")
	check(err)

	str := string(strbytes)

	l := utf8.RuneCountInString(str)

	for i := 13; i < l; i++ {
		if gotMarker(i, str) {
			fmt.Println(i + 1)
			break
		}
	}
}
