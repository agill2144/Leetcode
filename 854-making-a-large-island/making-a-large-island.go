func largestIsland(grid [][]int) int {
    n := len(grid)
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    var dfs func(r, c, id int) int
    dfs = func(r, c, id int) int {
        // base
        if r < 0 || r == n || c < 0 || c == n || grid[r][c] != 1 {return 0}

        // logic
        grid[r][c] = id
        count := 1
        for _, dir := range dirs {
            count += dfs(r+dir[0], c+dir[1], id)
        }
        return count
    }

    //1. pre-process connected components
    // find connected lands and group them with an id
    // each dfs on a land will return back its area
    // save this area in hashmap {id: area}
    // ** also mutate / update each cell with its id number **
    id := 2
    idToArea := map[int]int{}
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                idToArea[id] = dfs(i, j, id)
                id++
            }
        }
    }
    maxArea := 0    
    //2. for each 0 cell, find neighboring islands
    // each cell value is updated to be an id instead of 1 and 0
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 0 {
                uniqIDs := map[int]bool{}
                for _, dir := range dirs {
                    nr := i+dir[0]
                    nc := j+dir[1]
                    if nr < 0 || nr == n || nc < 0 || nc == n  {continue}
                    if _, ok := idToArea[grid[nr][nc]]; ok {
                        uniqIDs[grid[nr][nc]] = true
                    }
                    total := 1
                    for uid, _ := range uniqIDs {
                        total += idToArea[uid]
                    }
                    maxArea = max(maxArea, total)
                }
            }
        }
    }

    if maxArea == 0 {maxArea = n*n}
    return maxArea
}

// brute force
// for every 0 cell, explore all islands around it
// tc = o(n^4)
// sc = o(n^2) + o(n^2)
// func largestIsland(grid [][]int) int {
//     n := len(grid)
//     dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
//     maxArea := 0
//     var dfs func(r, c int, visited [][]bool) int
//     dfs = func(r, c int, visited [][]bool) int{
//         // base
//         if r < 0 || r == n || c < 0 || c == n || visited[r][c] || grid[r][c] != 1 {return 0}

//         // logic
//         visited[r][c] = true
//         count := 1
//         for _, dir := range dirs {
//             count += dfs(r+dir[0], c+dir[1], visited)
//         }
//         return count
//     }
//     for i := 0; i < n; i++ {
//         for j := 0; j < n; j++ {
//             if grid[i][j] == 0 {
                
//                 total := 1
//                 visited := make([][]bool, n)
//                 for ii := 0; ii < n; ii++ {visited[ii] = make([]bool, n)}
//                 for _, dir := range dirs {
//                     nr := i+dir[0]
//                     nc := j+dir[1]
//                     total += dfs(nr,nc,visited)
//                 }
                
//                 maxArea = max(maxArea, total)
//             }
//         }
//     }
//     if maxArea == 0 {maxArea = n*n}
//     return maxArea
// }