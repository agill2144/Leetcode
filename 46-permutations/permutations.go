func permute(nums []int) [][]int {
    out := [][]int{}
    var dfs func(start int)
    dfs = func(start int) {
        // base
        if start == len(nums) {
            path := make([]int, len(nums))
            copy(path, nums)
            out = append(out, path)
            return
        }

        // logic
        for i := start; i < len(nums); i++ {
            nums[i], nums[start] = nums[start], nums[i]
            dfs(start+1)
            nums[i], nums[start] = nums[start], nums[i]
        }
    }
    dfs(0)
    return out
}
// func permute(nums []int) [][]int {
//     out := [][]int{}
//     var dfs func(path []int)
//     dfs = func(path []int) {
//         // base
//         if len(path) == len(nums) {
//             newP := make([]int, len(path))
//             copy(newP, path)
//             out = append(out, newP)
//             return
//         }

//         // logic
//         for i := 0; i < len(nums); i++ {
//             if !contains(path, nums[i]) {
//                 path = append(path, nums[i])
//                 dfs(path)
//                 path = path[:len(path)-1]
//             }
//         }
//     }
//     dfs(nil)
//     return out
// }

// func contains(list []int, target int) bool {
//     for _, ele := range list {
//         if ele == target {return true}
//     }
//     return false
// }