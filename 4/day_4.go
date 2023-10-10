package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getRange(s string) (int64, int64) {
	l, r, _ := strings.Cut(s, "-")
	lnum, _ := strconv.ParseInt(l, 10, 64)
	rnum, _ := strconv.ParseInt(r, 10, 64)

	return lnum, rnum
}

func main() {
	day := 4
	file, err := os.Open("input" + fmt.Sprint(day) + ".txt")
	check(err)

	scanner := bufio.NewScanner(file)
	var ct int

	for scanner.Scan() {
		str := scanner.Text()
		ranges := strings.Split(str, ",")
		l1, r1 := getRange(ranges[0])
		l2, r2 := getRange(ranges[1])

		if (r2-r1)*(l2-l1) <= 0 {
			ct++
		}
	}

	fmt.Println(ct)
}
