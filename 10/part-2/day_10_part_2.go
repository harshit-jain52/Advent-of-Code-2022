package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(x int64) int64 {
	if x < 0 {
		x = -x
	}
	return x
}

func main() {
	day := 10
	file, _ := os.Open("../input" + fmt.Sprint(day) + ".txt")
	file2, _ := os.Create("output" + fmt.Sprint(day) + ".txt")
	defer file2.Close()

	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(file2)

	var (
		cycle int64
		x     int64
	)

	cycle, x = 0, 1

	for scanner.Scan() {
		str := scanner.Text()
		if str == "noop" {
			if abs(cycle%40-x) <= 1 {
				writer.WriteString("#")
			} else {
				writer.WriteString(".")
			}
			cycle++
			if cycle%40 == 0 {
				writer.WriteString("\n")
			}
		} else {
			add := strings.Split(str, " ")
			toadd, _ := strconv.Atoi(add[1])

			if abs(cycle%40-x) <= 1 {
				writer.WriteString("#")
			} else {
				writer.WriteString(".")
			}
			cycle++
			if cycle%40 == 0 {
				writer.WriteString("\n")
			}

			if abs(cycle%40-x) <= 1 {
				writer.WriteString("#")
			} else {
				writer.WriteString(".")
			}
			cycle++
			if cycle%40 == 0 {
				writer.WriteString("\n")
			}

			x += int64(toadd)
		}
	}

	writer.Flush()
}
