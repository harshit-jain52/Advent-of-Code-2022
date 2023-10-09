package main

import (
	"bufio"
	"fmt"
	"os"
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
	var file, err = os.Open("../input" + fmt.Sprint(day) + ".txt")
	check(err)

	scanner := bufio.NewScanner(file)
	var total int64
	var group [3]string
	idx := 0

	for scanner.Scan() {
		str := scanner.Text()
		group[idx] = str
		idx++

		if idx == 3 {
			idx = 0

		nested:
			for i := range group[0] {
				for j := range group[1] {
					for k := range group[2] {
						if (group[0][i] == group[1][j]) && (group[0][i] == group[2][k]) {
							total += getPriority(rune(group[0][i]))
							break nested
						}
					}
				}
			}
		}
	}

	fmt.Println(total)
}
