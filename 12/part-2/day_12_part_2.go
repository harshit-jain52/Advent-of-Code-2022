package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Position struct{
	X int
	Y int
}


func replaceAtIndex1(str string, replacement rune, index int) string {
    out := []rune(str)
    out[index] = replacement
    return string(out)
}

func enqueue(queue []Position, element Position) []Position { 
	queue = append(queue, element) 
	return queue 
} 
	  
func dequeue(queue []Position) (Position, []Position) { 
	element := queue[0]
	if len(queue) == 1 { 
		var tmp = []Position{} 
		return element, tmp 
		
	} 
	return element, queue[1:] 
}

func main(){
	day := 12
	file, _ := os.Open("../input" + fmt.Sprint(day) + ".txt")
	scanner := bufio.NewScanner(file)
	var grid[]string

	var start, end Position

	for scanner.Scan(){
		str:=scanner.Text()

		Spos := strings.IndexRune(str,'S')
		if Spos != -1{
			start.Y, start.X = len(grid),Spos
		}
		Epos := strings.IndexRune(str,'E')
		if Epos != -1{
			end.Y, end.X = len(grid),Epos
		}

		grid = append(grid, str)
	}

	file.Close()
	
	m,n:=len(grid), len(grid[0])
	grid[start.Y] = replaceAtIndex1(grid[start.Y],'a',start.X)
	grid[end.Y] = replaceAtIndex1(grid[end.Y],'z',end.X)

	var q[]Position
	var v Position
	
	steps := [4]Position{{0,1},{1,0},{0,-1},{-1,0}}
	dist:=make([][]int64, m)
	for i:=range dist{
		dist[i] = make([]int64,n)
		for j:=range dist[i]{
			dist[i][j] = -1
			if grid[i][j]=='a'{
				dist[i][j]=0
				q = enqueue(q,Position{j,i})
			}
		}
	}

	for{
		v, q = dequeue(q)
		if v == end{
			break
		}

		for st:=range steps{
			newV := Position{v.X+steps[st].X,v.Y+steps[st].Y}
			if newV.X<0 || newV.X>=n || newV.Y <0 || newV.Y>=m{
				continue
			}

			if dist[newV.Y][newV.X]!=-1 || (grid[newV.Y][newV.X]>grid[v.Y][v.X]+1) {
				continue
			}

			dist[newV.Y][newV.X] = 1 + dist[v.Y][v.X]
			q = enqueue(q,newV)
		}
	}

	println(dist[end.Y][end.X])
}