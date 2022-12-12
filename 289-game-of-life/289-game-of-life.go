
/*

0 = dead
1 = live

Any live cell with fewer than two live neighbors dies as if caused by under-population.
Any live cell with two or three live neighbors lives on to the next generation.
Any live cell with more than three live neighbors dies, as if by over-population.
Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.

Given the current state of the m x n grid board, return the next state.


Approach 1:
- Simulate the above rules with its immediate neighbours in a new mxn matrix
- Return the new matrix
- Time: o(mn)
- Space: o(mn)


Approach 2:
- Simulate the above rules with its immediate neighbours but mutate the same matrix
- But we will lose a cells previous state by doing that, which leads to side effects for other cells ( wrong answers )
- Then we can come up with a new standard of marking a cell dead or alive in a way it can tell its previous history
- To mark a cell dead
	- we are supposed to use 0
	- but now we will use 2 ( which means it was previously alive )

- To mark a cell alive
	- we are supposed to use 1
	- but now we will use 3 ( which means it was previously dead )

- 3 = alive ( but was previously dead )
- 2 = dead ( but was previously alive )
- 1 = alive ( no previous history )
- 0 = dead ( no previous history )
*/
func gameOfLife(board [][]int) [][]int {
	if board == nil || len(board) == 0 {
		return nil
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			numNeighborsAlive := countAliveNeighbors(i, j, board)
			if board[i][j] == 1 && (numNeighborsAlive < 2 || numNeighborsAlive > 3) {
				board[i][j] = 2
			} else if board[i][j] == 0 && numNeighborsAlive == 3 {
				board[i][j] = 3
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 3 {
				board[i][j] = 1
			}
			if board[i][j] == 2 {
				board[i][j] = 0
			}
		}
	}
	return board
}

/*
- neighbors are located in left, right, top, bottom, diagonally too
- left: j-1
- right: j+1
- top: i-1
- bottom: i+1
- diagonal left top: i-1, j-1
- diagonal left down: i+1, j-1
- diagonal right top: i-1, j+1
- diagonal right down: i+1, j+1
*/
func countAliveNeighbors(i, j int, board [][]int) int {
	// we could manually write the above rules or we can use a directions array
	dirs := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}, {-1, -1}, {1, -1}, {-1, 1}, {1, 1}}
	m := len(board)
	n := len(board[0])
	aliveCount := 0
	// check all directions if possible
	for _, dir := range dirs {
		r := i + dir[0]
		c := j + dir[1]
		if (r <= m-1 && r >= 0 && c <= n-1 && c >= 0) && (board[r][c] == 1 || board[r][c] == 2) {
			aliveCount++
		}
	}
	return aliveCount
}