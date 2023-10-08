package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func max(x int64, y int64) int64 {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	var day = 1
	file, err := os.Open("input" + fmt.Sprint(day) + ".txt")
	check(err)

	var ans, sum int64
	var curr string

	scanner := bufio.NewScanner(file)

	ans = 0
	for scanner.Scan() {
		curr = scanner.Text()
		if curr == "" {
			ans = max(ans, sum)
			sum = 0
		} else {
			x, _ := strconv.ParseInt(curr, 10, 64)
			sum += x
		}
	}

	file.Close()
	fmt.Println(ans)
}
