func subsets(nums []int) [][]int {
    out := [][]int{{}}
    for i := 0; i < len(nums); i++ {
        for _, list := range out {
            newL := make([]int, len(list))
            copy(newL, list)
            newL = append(newL, nums[i])
            out = append(out, newL)
        }
    }
    return out
}

// tc = o( (2^n) * n)
// func subsets(nums []int) [][]int {
//     out := [][]int{}
//     var dfs func(start int, path []int)
//     dfs = func(start int, path []int) {
//         // base
        
//         newL := make([]int, len(path))
//         copy(newL, path)
//         out = append(out, newL)

//         // logic
//         for i := start; i < len(nums); i++ {
//             path = append(path, nums[i])
//             dfs(i+1, path)
//             path = path[:len(path)-1]
//         }
//     }
//     dfs(0, nil)
//     return out
// }