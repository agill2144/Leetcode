func subsets(nums []int) [][]int {
    out := [][]int{}
    var dfs func(start int, path []int)
    dfs = func(start int, path []int) {
        // base
        // logic
        newP := make([]int, len(path))
        copy(newP, path)
        out = append(out, newP)
        
        for i := start; i < len(nums); i++ {
            // action
            path = append(path, nums[i])
            // recurse
            dfs(i+1, path)
            // backtrack
            path = path[:len(path)-1]
        }
    }
    dfs(0, nil)
    return out
}