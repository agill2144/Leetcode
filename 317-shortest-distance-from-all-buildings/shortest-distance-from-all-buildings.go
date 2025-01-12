/*
    approach: bfs from each house
    - instead of bfs from each 0 cell
        - there could be more 0s than 1s
        - therefore this would take more time in worst case
    - run a bfs from each house ( because number of 1s < number of 0s )
        - looking for adj 0's
        - and marking dist from that house to this 0 cell
        - dist == level
        - create a dists matrix of same size
        - and set dists in here for each 1 cell
        - along with houses count seen so far
        - its possible that a house is surrounded with 2s and so it could not reach any 0 cell
        - therefore its important to know how many houses were able to reach a specific 0 cell
    - do this ^ for each house
    - once the dists matrix is filled up
        - each 0th cell pos in this matrix, now tell us
        - the dist it would take for all houses to reach this 0th cell
        - because we want smallest possible dist, find the smallest one 
            - ONLY if the number of houses that reached this cell == number of houses in input matrix
        - and return that smallest dist

    m = len(matrix)
    n = len(matrix[0])
    k = num of houses
    time = o(k * mn)
    sc = o(mn) + o(mn)
*/
func shortestDistance(matrix [][]int) int {
    m := len(matrix)
    n := len(matrix[0])
    dists := make([][][]int, m)
    for i := 0; i < m; i++ {
        dists[i] = make([][]int, n)
        for j := 0; j < n; j++ {
            dists[i][j] = make([]int, 2) // [dist,hCount]
        }
    }
    count := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 1 {
                count++
                bfs(matrix, i, j, dists)               
            }
        }
    }
    minDist := math.MaxInt64
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if dists[i][j][1] == count {
                minDist = min(minDist, dists[i][j][0])
            }
        }
    }
    if minDist == math.MaxInt64 {return -1}
    return minDist
}

func bfs(matrix [][]int, r, c int, dists [][][]int) {
	m, n := len(matrix), len(matrix[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	visited[r][c] = true

	dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	q := [][]int{{r, c}}
	distance := 0

	for len(q) > 0 {
		distance++
        qSize := len(q)
		for qSize != 0 {
			cell := q[0]
			q = q[1:]
			cr, cc := cell[0], cell[1]
            qSize--
			for _, dir := range dirs {
				nr, nc := cr+dir[0], cc+dir[1]
				if nr >= 0 && nr < m && nc >= 0 && nc < n && matrix[nr][nc] == 0 && !visited[nr][nc] {
					visited[nr][nc] = true
					dists[nr][nc][0] += distance
					dists[nr][nc][1]++
					q = append(q, []int{nr, nc})
				}
			}
		}
	}
}
