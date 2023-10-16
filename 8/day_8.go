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

func max(x int8, y int8) int8 {
	if x < y {
		return y
	}
	return x
}

func min(x int8, y int8) int8 {
	if x > y {
		return y
	}
	return x
}

func isVisible(forest [][]int8, row int, col int, n int, m int) bool {

	var (
		final int8
		tmp   int8
	)

	final = 10

	// left
	tmp = -1
	for j := 0; j < col; j++ {
		tmp = max(tmp, forest[row][j])
	}
	final = min(final, tmp)

	//right
	tmp = -1
	for j := col + 1; j < m; j++ {
		tmp = max(tmp, forest[row][j])
	}
	final = min(final, tmp)

	// up
	tmp = -1
	for i := 0; i < row; i++ {
		tmp = max(tmp, forest[i][col])
	}
	final = min(final, tmp)

	//down
	tmp = -1
	for i := row + 1; i < n; i++ {
		tmp = max(tmp, forest[i][col])
	}
	final = min(final, tmp)

	return final < forest[row][col]
}

func main() {
	day := 8
	file, err := os.Open("input" + fmt.Sprint(day) + ".txt")

	check(err)

	forest := make([][]int8, 0, 1)

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
		row := make([]int8, 0, 1)
		for i := 0; i < m; i++ {
			num, err := strconv.Atoi(string(str[i]))
			check(err)

			row = append(row, int8(num))
		}

		forest = append(forest, row)
	}

	//fmt.Println(forest)

	ct := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if isVisible(forest, i, j, n, m) {
				ct++
			}
		}
	}

	fmt.Println(ct)
}
