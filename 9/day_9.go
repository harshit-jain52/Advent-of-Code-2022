package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type void struct{}

var blank void

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func main() {
	day := 9

	file, _ := os.Open("input" + fmt.Sprint(day) + ".txt")
	scanner := bufio.NewScanner(file)

	var hx, hy, tx, ty = 0, 0, 0, 0
	m := make(map[[2]int]void)

	arr := [2]int{tx, ty}
	m[arr] = blank

	for scanner.Scan() {
		str := scanner.Text()
		instruct := strings.Split(str, " ")
		direction := instruct[0]
		steps, _ := strconv.Atoi(instruct[1])

		for i := 1; i <= steps; i++ {
			if direction == "U" {
				hy++
			} else if direction == "D" {
				hy--
			} else if direction == "R" {
				hx++
			} else if direction == "L" {
				hx--
			}

			if abs(hx-tx) <= 1 && abs(hy-ty) <= 1 {
				continue
			}

			if hx-tx == 0 {
				if hy > ty {
					ty++
				} else {
					ty--
				}
			} else if hy-ty == 0 {
				if hx > tx {
					tx++
				} else {
					tx--
				}
			} else {
				if hx > tx && hy > ty {
					tx++
					ty++
				} else if hx > tx && hy < ty {
					tx++
					ty--
				} else if hx < tx && hy > ty {
					tx--
					ty++
				} else if hx < tx && hy < ty {
					tx--
					ty--
				}
			}

			arr := [2]int{tx, ty}
			m[arr] = blank
		}
	}

	fmt.Println(len(m))
}
