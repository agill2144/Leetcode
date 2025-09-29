func combinationSum2(candidates []int, target int) [][]int {
    counts := map[int]int{}
    for i := 0; i < len(candidates); i++ {
        counts[candidates[i]]++
    }
    freq := [][]int{}
    for k, v := range counts {
        freq = append(freq, []int{k,v})
    }
    out := [][]int{}
    var dfs func(start, sum int, path []int)
    dfs = func(start, sum int, path []int) {
        // base
        if sum >= target {
            if sum == target {
                newP := make([]int, len(path))
                copy(newP, path)
                out =append(out, newP)
            }
            return
        }


        // logic
        for i := start; i < len(freq); i++ {
            if freq[i][1] == 0 {continue}
            sum += freq[i][0]
            freq[i][1]--
            path = append(path, freq[i][0])
            dfs(i, sum, path)
            path = path[:len(path)-1]
            freq[i][1]++
            sum -= freq[i][0]
        }
    }
    dfs(0,0,nil)
    return out
}
// func combinationSum2(candidates []int, target int) [][]int {
//     sort.Ints(candidates)
//     out := [][]int{}
//     var dfs func(start int, sum int, path []int)
//     dfs = func(start, sum int, path []int) {
//         // base
//         if sum >= target {
//             if sum == target {
//                 newP := make([]int, len(path))
//                 copy(newP, path)
//                 out = append(out, newP)
//             }
//             return
//         }

//         // logic
//         for i := start; i < len(candidates); i++ {
//             sum += candidates[i]
//             path = append(path, candidates[i])
//             dfs(i+1, sum, path)
//             path = path[:len(path)-1]
//             sum -= candidates[i]

//             for i+1 < len(candidates) && candidates[i] == candidates[i+1] {i++}
//         }
//     }
//     dfs(0,0,nil)
//     return out
// }