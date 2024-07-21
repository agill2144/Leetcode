func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
    rowOrder := topSort(rowConditions, k)
    colOrder := topSort(colConditions, k)
    if len(rowOrder) == 0 || len(colOrder) == 0 {return nil}

    valToRowIdx := map[int]int{}
    valToColIdx := map[int]int{}
    for i := 0; i < len(rowOrder); i++ {valToRowIdx[rowOrder[i]] = i}
    for i := 0; i < len(colOrder); i++ {valToColIdx[colOrder[i]] = i}
    fmt.Println(rowOrder, colOrder)

    matrix := make([][]int, k)
    for i := 0; i < len(matrix); i++ {matrix[i] = make([]int,k)}

    for i := 1; i <= k; i++ {
        r := valToRowIdx[i]
        c := valToColIdx[i]
        matrix[r][c] = i
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