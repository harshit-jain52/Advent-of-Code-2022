package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func min(x int64, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

var (
	ans         int64
	spaceToFree int64
)

func calc(idx *int, n int, lines []string) int64 {

	var data int64
	data = 0

	var str string
	*idx = *idx + 1 // ls

	for *idx < n {
		str = lines[*idx]
		*idx = *idx + 1

		if str[0] == '$' {
			cd := strings.Split(str, " ")
			if cd[2] == ".." {
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

	return data
}

func getAns(idx *int, n int, lines []string) int64 {

	var data int64
	data = 0

	var str string
	*idx = *idx + 1 // ls

	for *idx < n {
		str = lines[*idx]
		*idx = *idx + 1

		if str[0] == '$' {
			cd := strings.Split(str, " ")
			if cd[2] == ".." {
				if data >= spaceToFree {
					ans = min(ans, data)
				}
				return data
			}

			data += getAns(idx, n, lines)
			continue
		}
		if str[0] == 'd' {
			continue
		}

		num := strings.Split(str, " ")
		add, _ := strconv.Atoi(num[0])
		data += int64(add)
	}

	if data >= spaceToFree {
		ans = min(ans, data)
	}
	return data
}

func main() {
	day := 7
	file, _ := os.Open("../input" + fmt.Sprint(day) + ".txt")

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

	usedSpace := calc(p, n, filelines)
	spaceToFree = 30000000 - (70000000 - usedSpace)
	ans = 30000000
	idx = 1

	getAns(p, n, filelines)
	fmt.Println(ans)
}
