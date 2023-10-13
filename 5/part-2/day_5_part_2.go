package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

// type stack []rune

func push(s []rune, c rune) []rune {
	return append(s, c)
}

func pop(s []rune) ([]rune, rune) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func top(s []rune) rune {
	l := len(s)
	return s[l-1]
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isAlphabet(c rune) bool {
	return (c >= 'A' && c <= 'Z')
}

func reverse(s []rune) []rune {
	l := len(s)
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

func main() {
	day := 5
	file, err := os.Open("../input" + fmt.Sprint(day) + ".txt")
	check(err)

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	firstline := scanner.Text()
	n := utf8.RuneCountInString(firstline)
	// fmt.Println(firstline)

	numStack := (n + 1) / 4

	allstacks := make([][]rune, 0, 1)
	for i := 0; i < numStack; i++ {
		new := make([]rune, 0, 1)
		allstacks = append(allstacks, new)
	}

	_, err = file.Seek(0, 0)
	check(err)
	scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}
		for i := 1; i < n; i += 4 {
			if !isAlphabet(rune(line[i])) {
				continue
			}

			allstacks[i/4] = append(allstacks[i/4], rune(line[i]))
		}
	}

	// fmt.Println(allstacks)

	for i := 0; i < numStack; i++ {
		allstacks[i] = reverse(allstacks[i])
	}

	// fmt.Println(allstacks)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")

		num, _ := strconv.ParseInt(words[1], 10, 64)
		from, _ := strconv.ParseInt(words[3], 10, 64)
		to, _ := strconv.ParseInt(words[5], 10, 64)

		new := make([]rune, 0, 0)
		for i := 0; i < int(num); i++ {
			var c rune
			allstacks[from-1], c = pop(allstacks[from-1])
			new = push(new, c)
		}

		for i := len(new) - 1; i >= 0; i-- {
			allstacks[to-1] = push(allstacks[to-1], new[i])
		}

	}
	// fmt.Println(allstacks)

	for i := 0; i < numStack; i++ {
		fmt.Printf("%s", strconv.QuoteRune(top(allstacks[i])))
	}
}
