package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func inRange(n int64) bool {
	nums := []int64{20, 60, 100, 140, 180, 220}

	for _, num := range nums {
		if n == num {
			return true
		}
	}

	return false
}

func main() {
	day := 10
	file, _ := os.Open("input" + fmt.Sprint(day) + ".txt")
	scanner := bufio.NewScanner(file)

	var (
		cycle    int64
		x        int64
		strength int64
	)

	cycle, x, strength = 1, 1, 0

	for scanner.Scan() {
		str := scanner.Text()
		if str == "noop" {
			if inRange(cycle) {
				strength += cycle * x
			}
			cycle++
		} else {
			add := strings.Split(str, " ")
			toadd, _ := strconv.Atoi(add[1])

			if inRange(cycle) {
				strength += cycle * x
			}
			cycle++

			if inRange(cycle) {
				strength += cycle * x
			}
			cycle++

			x += int64(toadd)
		}
	}

	fmt.Println(strength)
}
