package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ans int64

func calc(idx *int, n int, lines []string) int64 {

	var data int64
	data = 0

	var str string
	*idx = *idx + 1 // ls

	for *idx < n {
		str = lines[*idx]
		*idx = *idx + 1
		//fmt.Println(str)
		if str[0] == '$' {
			cd := strings.Split(str, " ")
			if cd[2] == ".." {
				if data <= 100000 {
					ans += data
				}
				return data
			}

			data += calc(idx, n, lines)
			continue
		}
		if str[0] == 'd' {
			continue
		}

		num := strings.Split(str, " ")
		add, _ := strconv.Atoi(num[0])
		data += int64(add)
	}

	if data <= 100000 {
		ans += data
	}
	return data
}

func main() {
	day := 7
	file, _ := os.Open("input" + fmt.Sprint(day) + ".txt")

	_ = bufio.NewScanner(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var filelines []string

	for scanner.Scan() {
		filelines = append(filelines, scanner.Text())
	}

	idx := 1
	n := len(filelines)

	var p = &idx

	calc(p, n, filelines)

	fmt.Println(ans)
}
