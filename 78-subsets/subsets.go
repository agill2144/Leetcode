func subsets(nums []int) [][]int {
    out := [][]int{{}}
    for i := 0; i < len(nums); i++ {
        for _, p := range out {
            newP := make([]int, len(p))
            copy(newP, p)
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
//         if start == len(nums) {return}

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