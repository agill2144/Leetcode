func permute(nums []int) [][]int {
    
    out := [][]int{}
    var dfs func(path []int)
    dfs = func(path []int) {
        // base
        if len(path) == len(nums) {
            newPath := make([]int, len(path))
            copy(newPath, path)
            out = append(out, newPath)
            return
        }
        
        // logic
        for i := 0; i < len(nums); i++ {
            if !listContains(path, nums[i]) {
                path = append(path, nums[i])
                dfs(path)
                path = path[:len(path)-1]
            }
        }
    }
    dfs(nil)
    return out
}

func listContains(list []int, x int) bool {
    for _, ele := range list {
        if ele == x {return true}
    }
    return false
}