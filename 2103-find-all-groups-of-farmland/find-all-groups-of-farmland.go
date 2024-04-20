// time = 2 o(mn) ; o(mn)
// space = o(1)
// brute force; greedy
func findFarmland(land [][]int) [][]int {
    m := len(land)
    n := len(land[0])
    out := [][]int{}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if land[i][j] == 1 {
                sr, sc := i, j
                er, ec := -1,-1

                // go as far down as possible
                r1 := i
                for r1 < m && land[r1][j] == 1 { r1++ }
                r1--
                er = r1

                // go as far right as possible
                c1 := j
                for c1 < n && land[i][c1] == 1 { c1++ }
                c1--
                ec = c1
                
                // save the sr, sc, er, ec
                out = append(out, []int{sr, sc, er, ec})

                // mark the cells between sr, sc and er and ec visited
                for r := sr; r <= er ; r++ {
                    for c := sc; c <= ec; c++ {
                        land[r][c] = 0
                    }
                }
            }
        }
    }
    return out
}
// dfs; conntected component
// func findFarmland(land [][]int) [][]int {
//     m := len(land)
//     n := len(land[0])
//     dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
//     maxR, maxC := math.MinInt64, math.MinInt64
//     var dfs func(r, c int)
//     dfs = func(r, c int) {
//         // base
//         if r < 0 || r == m || c < 0 || c == n || land[r][c] != 1 {return}

//         // logic
//         maxR = max(r, maxR)
//         maxC = max(c, maxC)
//         land[r][c] = 0
//         for _, dir := range dirs {
//             dfs(r+dir[0], c+dir[1])
//         }
//     }
//     out := [][]int{}
//     for i := 0; i < m; i++ {
//         for j := 0; j < n; j++ {
//             if land[i][j] == 1 {
//                 sr, sc := i, j
//                 maxR, maxC = math.MinInt64, math.MinInt64
//                 dfs(i,j)
//                 out = append(out, []int{sr,sc,maxR,maxC})
//             }
//         }
//     }
//     return out
// }