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

func main() {
	var day = 1
	file, err := os.Open("../input" + fmt.Sprint(day) + ".txt")
	check(err)

	var max1, max2, max3, sum int64
	var curr string

	scanner := bufio.NewScanner(file)

	max1 = 0
	for scanner.Scan() {
		curr = scanner.Text()
		if curr == "" {
			switch {
			case sum > max1:
				max3 = max2
				max2 = max1
				max1 = sum
			case sum > max2:
				max3 = max2
				max2 = sum
			case sum > max3:
				max3 = sum
			}
			sum = 0
		} else {
			x, _ := strconv.ParseInt(curr, 10, 64)
			sum += x
		}
	}

	file.Close()
	fmt.Println(max1 + max2 + max3)
}
