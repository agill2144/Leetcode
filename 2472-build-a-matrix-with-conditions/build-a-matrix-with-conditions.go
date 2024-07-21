func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
    rowOrder := topSort(rowConditions, k)
    valToRowIdx := map[int]int{}
    colOrder := topSort(colConditions, k)
    if len(rowOrder) == 0 || len(colOrder) == 0 {return nil}
    for i := 0; i < len(rowOrder); i++ {
        valToRowIdx[rowOrder[i]] = i
    }

    matrix := make([][]int, k)
    // put val on each row first (col 0)
    rPtr := 0
    for i := 0; i < len(matrix); i++ {
        matrix[i] = make([]int,k)
        matrix[i][0] = rowOrder[rPtr]
        rPtr++
    }

    // now put vals for cols from the back
    cPtr := len(colOrder)-1
    for j := len(matrix[0])-1; j >= 0; j-- {
        val := colOrder[cPtr]
        rowIdx := valToRowIdx[val]
        colIdx := j
        matrix[rowIdx][0], matrix[rowIdx][colIdx] = matrix[rowIdx][colIdx], matrix[rowIdx][0]
        cPtr--
    }
    return matrix

}


func topSort(edges [][]int, n int) []int {
    indegrees := make([]int, n+1)
    adjList := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        a, b := edges[i][0], edges[i][1]
        adjList[a] = append(adjList[a], b)
        indegrees[b]++
    }
    q := []int{}
    count := 0
    for i := 1; i < len(indegrees); i++ {
        if indegrees[i] == 0 {
            q = append(q, i)
        }
    }
    order := []int{}
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        order = append(order, dq)
        count++
        for _, nei := range adjList[dq] {
            indegrees[nei]--
            if indegrees[nei] == 0 {
                q = append(q, nei)
            }
        }
    }  
    if count != n {return nil}
    return order
}