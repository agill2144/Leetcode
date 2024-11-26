func findChampion(n int, edges [][]int) int {
    indegrees := make([]int, n)
    for i := 0; i < len(edges); i++ {
        v := edges[i][1]
        indegrees[v]++
    }
    ans := -1
    for i := 0; i < len(indegrees); i++ {
        if indegrees[i] == 0 {
            if ans == -1 {
                ans = i
            } else {
                return -1
            }
        }
    }
    return ans
}