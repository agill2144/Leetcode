func subsets(nums []int) [][]int {
    out := [][]int{{}}
    for i := 0; i < len(nums); i++ {
        // range will take a snapshot of output array , so modifying it while 
        // looping over out will not change the initial snapshot of output array
        // that range started looping over
        for _, curr := range out {
            newP := make([]int, len(curr))
            copy(newP, curr)
            newP = append(newP, nums[i])
            out = append(out, newP)
        }
    }
    return out
}


// func subsets(nums []int) [][]int {
//     out := [][]int{}
//     var dfs func(start int, path []int)
//     dfs = func(start int, path []int) {
//         // base
//         newP := make([]int, len(path))
//         copy(newP, path)
//         out = append(out, newP)

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