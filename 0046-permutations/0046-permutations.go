func permute(nums []int) [][]int {
    out := [][]int{}
    var dfs func(path []int)
    dfs = func(path []int) {
        // base
        if len(path) == len(nums) {
            newP := make([]int, len(path))
            copy(newP, path)
            out = append(out, newP)
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

func listContains(list []int, ele int) bool {
    for _, x := range list{ 
        if x == ele {return true}
    }
    return false
} 