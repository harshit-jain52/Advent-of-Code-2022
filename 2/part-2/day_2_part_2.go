package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var day = 2
	file, err := os.Open("../input" + fmt.Sprint(day) + ".txt")
	check(err)

	scanner := bufio.NewScanner(file)

	var score [4][4]int64
	var i, j int64
	for i = 1; i <= 3; i++ {
		for j = 1; j <= 3; j++ {
			switch i {
			case 1:
				switch j {
				case 1:
					score[i][j] = 3
				case 3:
					score[i][j] = 2 + 6
				default:
					score[i][j] = i + 3
				}
			case 2:
				switch j {
				case 1:
					score[i][j] = 1
				case 3:
					score[i][j] = 3 + 6
				default:
					score[i][j] = i + 3
				}
			case 3:
				switch j {
				case 1:
					score[i][j] = 2
				case 3:
					score[i][j] = 1 + 6
				default:
					score[i][j] = i + 3
				}
			}

		}
	}

	var (
		total int64
		p1    int64
		p2    int64
		round string
	)

	for scanner.Scan() {
		round = scanner.Text()
		p1 = 1 + int64(round[0]) - int64('A')
		p2 = 1 + int64(round[2]) - int64('X')

		total += score[p1][p2]
	}

	fmt.Println(total)
}
