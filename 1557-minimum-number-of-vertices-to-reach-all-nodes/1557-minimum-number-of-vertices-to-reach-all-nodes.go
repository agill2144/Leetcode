func findSmallestSetOfVertices(n int, edges [][]int) []int {
    indegrees := make([]int, n)
    for i := 0; i < len(edges); i++ {
        dest := edges[i][1]
        indegrees[dest]++
    }
    out := []int{}
    for i := 0; i < len(indegrees); i++ {
        if indegrees[i] == 0 {out = append(out, i)}
    }
    return out
}