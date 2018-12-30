package main

import (
	"fmt"
	"os"
)

// 读取迷宫数据文件
func readMazeData(filepath string) [][]int {
	// fmt.Println(filepath)
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	// fmt.Println(file)
	var row, col int
	// 读取文件数据并保存到row, col变量中
	fmt.Fscanf(file, "%d %d", &row, &col)
	// fmt.Println(row, col)

	maze := make([][]int, row)

	for i := range maze {
		maze[i] = make([]int, col)

		for j := range maze[i] {
			// 读文件数据并保存到切片
			fmt.Fscanf(file, "%d", &maze[i][j])
		} 
	}
	return maze
}

// 二维切片中元素的位置i,j（不是坐标）
type point struct {
	i, j int
}

// 点的下标相加方法，值传递，返回新的point
func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool){
	// 行越界
	if p.i < 0 || p.i >= len(grid){
		return 0, false
	}
	// 列越界
	if p.j < 0 || p.j >= len(grid[p.i]){
		return 0, false
	}

    return grid[p.i][p.j], true
}

// 当前点的探索方向，推荐上左下右的顺序（注意与坐标的区别）
var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

// 走迷宫函数
func walk(maze [][]int, start, end point) [][]int {
	// 维护的二维切片，表示从起始点到达当前位置所经过的步数
	// 0表示未经过的点（起始点除外）
	steps := make([][]int, len(maze))
    for i := range steps{
		steps[i] = make([]int, len(maze[i]))
	}
   
	// 起始点加入队列 
	Q := []point{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]
		
		if cur == end{
			break
		}

		for _, dir := range dirs{
			next := cur.add(dir)  // 新发现目标点，为实现点位置的加法运算，需要为point结构体定义相应的方法
		    // 探索点需要满足的条件
			// maze在next点的值为0，表示可走
			// next不能回到start起始点
			// 当前点到next点的步数为0，否则表示已走过
			
			// 表示不能走
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}
			// 已走过
			val, ok = next.at(steps)
			if !ok || val == 0 {
				continue
			}
			// 回到原点 
			if next == start {
				continue
			}

			// 当前步数
			curSteps, _ := cur.at(steps) 
			// next点步数
			steps[next.i][next.j] = curSteps + 1
			// next点加入探索队列
			Q = append(Q, next)
		}
	}
	return steps
}

func main() {
	mazeDataPath := "go_pratice_code\\learn_ccmouse_code\\maze\\maze.in"
	maze := readMazeData(mazeDataPath)

	// for _, row := range maze {
	// 	for _, val := range row {
	// 		fmt.Printf("%d ", val)
	// 	}
	// 	fmt.Println()
	// }

	// 走迷宫函数，传入maze数据,起始点以及终点
	steps := walk(maze, point{0, 0}, point{len(maze)-1, len(maze[0])-1})
	fmt.Println(steps)

	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}