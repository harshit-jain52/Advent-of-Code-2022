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

	file, _ := os.Open("../input" + fmt.Sprint(day) + ".txt")
	scanner := bufio.NewScanner(file)

	var rope [10][2]int

	for i := 0; i < 10; i++ {
		arr := [2]int{0, 0}
		rope[i] = arr
	}

	m := make(map[[2]int]void)

	m[rope[9]] = blank

	for scanner.Scan() {
		str := scanner.Text()
		instruct := strings.Split(str, " ")
		direction := instruct[0]
		steps, _ := strconv.Atoi(instruct[1])

		for i := 1; i <= steps; i++ {
			if direction == "U" {
				rope[0][1]++
			} else if direction == "D" {
				rope[0][1]--
			} else if direction == "R" {
				rope[0][0]++
			} else if direction == "L" {
				rope[0][0]--
			}

			for j := 1; j <= 9; j++ {
				hx, hy, tx, ty := rope[j-1][0], rope[j-1][1], rope[j][0], rope[j][1]
				if abs(hx-tx) <= 1 && abs(hy-ty) <= 1 {
					continue
				}

				if hx-tx == 0 {
					if hy > ty {
						rope[j][1]++
					} else {
						rope[j][1]--
					}
				} else if hy-ty == 0 {
					if hx > tx {
						rope[j][0]++
					} else {
						rope[j][0]--
					}
				} else {
					if hx > tx && hy > ty {
						rope[j][0]++
						rope[j][1]++
					} else if hx > tx && hy < ty {
						rope[j][0]++
						rope[j][1]--
					} else if hx < tx && hy > ty {
						rope[j][0]--
						rope[j][1]++
					} else if hx < tx && hy < ty {
						rope[j][0]--
						rope[j][1]--
					}
				}
			}

			m[rope[9]] = blank
		}
	}

	fmt.Println(len(m))
}
