func subsets(nums []int) [][]int {
    out := [][]int{}
    var dfs func(start int, path []int)
    dfs = func(start int, path []int) {
        // base
        
        newL := make([]int, len(path))
        copy(newL, path)
        out = append(out, newL)

        // logic
        for i := start ; i < len(nums); i++ {
            // action
            path = append(path, nums[i])
            // recurse
            dfs(i+1, path)
            // backtrack
            path = path[:len(path)-1]
        }
    }
    dfs(0, []int{})
    return out
}