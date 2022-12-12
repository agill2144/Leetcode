package main

import "fmt"

func main() {
	fmt.Println(ratInMaze([][]int{
		{1, 0, 0, 0},
		{1, 1, 0, 1},
		{1, 1, 0, 0},
		{0, 1, 1, 1},
	}))
	fmt.Println(ratInMaze([][]int{
		{1, 0},
		{1, 0},
	}))
}

/*
Rat in a Maze Problem - I
Medium Accuracy: 37.73% Submissions: 100k+ Points: 4
Consider a rat placed at (0, 0) in a square matrix of order N * N.
It has to reach the destination at (N - 1, N - 1).
Find all possible paths that the rat can take to reach from source to destination.
The directions in which the rat can move are 'U'(up), 'D'(down), 'L' (left), 'R' (right).
Value 0 at a cell in the matrix represents that it is blocked and
rat cannot move to it while value 1 at a cell in the matrix represents that rat can be travel through it.
Note: In a path, no cell can be visited more than one time. If the source cell is 0, the rat cannot move to any other cell.

Example 1:
Input:
N = 4
m[][] = {
	{1, 0, 0, 0},
	{1, 1, 0, 1},
	{1, 1, 0, 0},
	{0, 1, 1, 1}
}
Output:
DDRDRR DRDDRR
Explanation:
The rat can reach the destination at
(3, 3) from (0, 0) by two paths - DRDDRR
and DDRDRR, when printed in sorted order
we get DDRDRR DRDDRR.

Example 2:
Input:
N = 2
m[][] = {
{1, 0},
{1, 0}
}
Output:
-1
Explanation:
No path exists and destination cell is
blocked.
Your Task:
You don't need to read input or print anything. Complete the function printPath() which takes N and 2D array m[ ][ ] as input parameters and returns the list of paths in lexicographically increasing order.
Note: In case of no path, return an empty list. The driver will output "-1" automatically.
*/

type dir struct {
	row, col int
	p        string
}

func ratInMaze(maze [][]int) (int, []string) {
	m := len(maze)
	n := len(maze[0])
	dr := m - 1
	dc := n - 1
	count := 0
	paths := []string{}
	var dfs func(r, c int, path string)
	dirs := []dir{
		{-1, 0, "U"},
		{0, -1, "L"},
		{1, 0, "D"},
		{0, 1, "R"},
	}
	dfs = func(r, c int, path string) {
		// base
		if r == m || r < 0 || c == n || c < 0 || maze[r][c] == -1 || maze[r][c] == 0 {
			return
		}
		if r == dr && c == dc {
			count++
			paths = append(paths, path)
			return
		}

		// logic
		maze[r][c] = -1 // mark cell visited
		for _, direction := range dirs {
			dfs(r+direction.row, c+direction.col, path+direction.p)
		}
		maze[r][c] = 1 // mark it unvisited
	}
	dfs(0, 0, "")
	if count == 0 {
		return -1, nil
	}
	return count, paths
}
