func findJudge(n int, trust [][]int) int {
    indegrees := make([]int, n+1)
    for i := 0; i < len(trust); i++ {
        a := trust[i][0]
        b := trust[i][1]
        indegrees[b]++
        indegrees[a]--
    }
    for i := 1; i <= n; i++ {
        if indegrees[i] == n-1 {return i}
    }
    return -1
}