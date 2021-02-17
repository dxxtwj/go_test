package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][] int {
	file, err := os.Open(filename)// 打开这个文件
	if err != nil {
		panic(err)
	}
	var row, col int// 定义行和列,
	// row 是行 竖下去的那种是行
	fmt.Fscanf(file, "%d %d", &row, &col)// 设置row,col，首先要知道几行几列， 值是 6 和 5，因为文件第一行写好了
	//fmt.Println(row, col)// 6, 5
	maze := make([][]int, row)
	//fmt.Println(maze)// // [[],[],[],[],[],[],[]]
	for i := range maze {// 循环第一个slice
		maze[i] = make([]int, col)
		for j := range maze[i] {// 循环第一个slice里面的小slice ，也就是说要做横向遍历


			// maze.in中的行和列弄到file这个变量中
			// 第一行（竖）第一列（横）是0 所以是0
			// 第二行是0，第二列是0  所以是0
			// 第三行是0，第三列是0  所以是0
			// 第四行是1，第四列是1  所以是1
			// ...
			fmt.Fscanf(file, "%d", &maze[i][j])// maze.in 中以列的形式放到变量中（横向遍历）
		}
	}
	return maze
}


type point struct {
	i, j int
}
// 执行到这些坐标，说明不可继续往下走
// maze at next is 0
// and steps at next is 0
// and next != start  走过了
var dirs = [4]point{
	{-1, 0},// 开始的上边
	{0, -1},// 开始的左边
	{1, 0},// 开始的第二步
	{0, 1},// 墙壁
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {// 行  p.i < 0 说明越界了 p.i >= len(grid) 说明撞墙了，遇到了1，要等于0才能走
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[p.i]) {// 列（横） p.j < 0 说明越界了 p.j >= len(grid[p.i]) 说明撞墙了，遇到了1，要等于0才能走
		return 0, false
	}
	//fmt.Println("grid", grid)
	return grid[p.i][p.j], true// 第一个值当前走的值，第二个是是否越界了
}
// maze 内容, start 开始;end 结束的行和列
func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	Q := []point{start}
	fmt.Println("这里是Q队列", Q)
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]
		if cur == end {// 开始==结束
			break
		}
		//fmt.Println(cur)
		//{0 0}
		//{1 0}
		//{2 0}
		//{1 1}
		//{1 2}
		//{0 2}
		//{2 2}
		//{0 3}
		//{0 4}
		//{1 4}
		//{2 4}
		//{3 4}
		//{3 3}
		//{4 3}
		//{4 2}
		//{5 3}
		//{5 2}
		for _, dir := range dirs {
			next := cur.add(dir)
			//fmt.Println(next)
			//{-1 0}
			//{0 -1}
			//{1 0}
			//{0 1}
			//{0 0}
			//{1 -1}
			//{2 0}
			//{1 1}
			//{1 0}
			//{2 -1}
			//{3 0}
			//{2 1}
			//{0 1}
			//{1 0}
			//{2 1}
			//{1 2}
			//{0 2}
			//{1 1}
			//{2 2}
			//{1 3}
			//{-1 2}
			//{0 1}
			//{1 2}
			//{0 3}
			//{1 2}
			//{2 1}
			//{3 2}
			//{2 3}
			//{-1 3}
			//{0 2}
			//{1 3}
			//{0 4}
			//{-1 4}
			//{0 3}
			//{1 4}
			//{0 5}
			//{0 4}
			//{1 3}
			//{2 4}
			//{1 5}
			//{1 4}
			//{2 3}
			//{3 4}
			//{2 5}
			//{2 4}
			//{3 3}
			//{4 4}
			//{3 5}
			//{2 3}
			//{3 2}
			//{4 3}
			//{3 4}
			//{3 3}
			//{4 2}
			//{5 3}
			//{4 4}
			//{3 2}
			//{4 1}
			//{5 2}
			//{4 3}
			//{4 3}
			//{5 2}
			//{6 3}
			//{5 4}
			//{4 2}
			//{5 1}
			//{6 2}
			//{5 3}
			// maze at next is 0
			// and steps at next is 0
			// and next != start  走过了
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}
			if next == start { // 回到原点了
				continue
			}
			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1// 因为我们需要把当前的步数加1,也就是已经探索了的区域，行和列如果有的话加1
			Q = append(Q, next)
		}
	}
	return steps
}

func main() {
	maze := readMaze("./maze.in")// 返回一个列（横） [[0 1 0 0 0] [0 0 0 1 0] [0 1 0 1 0] [1 1 1 0 0] [0 1 0 0 1] [0 1 0 0 0]]
fmt.Println("这个是终点，目标坐标", len(maze), len(maze[0]))
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}
