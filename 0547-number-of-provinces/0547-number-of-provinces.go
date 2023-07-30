func findCircleNum(isConnected [][]int) int {
    n := len(isConnected)
   
    visited := make([]bool, n)
    var dfs func(row int)
    dfs = func(row int) {
        // base
        if visited[row] {return}
        
        // logic
        visited[row] = true
        for j := 0; j < n; j++ {
            if isConnected[row][j] == 1 && row != j && !visited[j] {
                dfs(j)
            }
        }
    }
        count := 0

    
    for i := 0; i < n; i++ {
        if !visited[i] {
            dfs(i)
            count++
        }
    }
   
    return count
}
