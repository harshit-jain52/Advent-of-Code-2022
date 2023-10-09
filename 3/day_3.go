package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getPriority(c rune) int64 {
	if c >= 'a' {
		return (int64(c) - int64('a') + 1)
	} else {
		return (int64(c) - int64('A') + 27)
	}
}

func main() {
	var day = 3
	var file, err = os.Open("input" + fmt.Sprint(day) + ".txt")
	check(err)

	scanner := bufio.NewScanner(file)
	var total int64

	for scanner.Scan() {
		str := scanner.Text()
		n := utf8.RuneCountInString(str)

	nested:
		for i := 0; i < n/2; i++ {
			for j := n / 2; j < n; j++ {
				if str[i] == str[j] {
					total += getPriority(rune(str[i]))
					break nested
				}
			}
		}
	}

	fmt.Println(total)
}
