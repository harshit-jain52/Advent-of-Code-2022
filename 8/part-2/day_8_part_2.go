package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func max(x int64, y int64) int64 {
	if x < y {
		return y
	}
	return x
}

func scenicScore(forest [][]int64, row int, col int, n int, m int) int64 {
	var (
		prod int64
		ct   int64
	)
	prod = 1

	// left
	ct = 0
	for j := col - 1; j >= 0; j-- {
		ct++
		if forest[row][j] >= forest[row][col] {
			break
		}
	}
	prod *= ct

	// right
	ct = 0
	for j := col + 1; j < m; j++ {
		ct++
		if forest[row][j] >= forest[row][col] {
			break
		}
	}
	prod *= ct

	// up
	ct = 0
	for i := row - 1; i >= 0; i-- {
		ct++
		if forest[i][col] >= forest[row][col] {
			break
		}
	}
	prod *= ct

	// down
	ct = 0
	for i := row + 1; i < n; i++ {
		ct++
		if forest[i][col] >= forest[row][col] {
			break
		}
	}
	prod *= ct

	return prod
}

func main() {
	day := 8
	file, err := os.Open("../input" + fmt.Sprint(day) + ".txt")

	check(err)

	forest := make([][]int64, 0, 1)

	var (
		m = -1
		n = 0
	)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := scanner.Text()
		n++
		if m == -1 {
			len := utf8.RuneCountInString(str)
			m = len
		}
		row := make([]int64, 0, 1)
		for i := 0; i < m; i++ {
			num, err := strconv.Atoi(string(str[i]))
			check(err)

			row = append(row, int64(num))
		}

		forest = append(forest, row)
	}

	//fmt.Println(forest)

	var score int64
	score = 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			score = max(score, scenicScore(forest, i, j, n, m))
		}
	}

	fmt.Println(score)
}
