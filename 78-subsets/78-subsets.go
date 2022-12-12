func subsets(nums []int) [][]int {
    result := [][]int{}
    var dfs func(i int, path []int)
    dfs = func(i int, path []int) {
        // base
        
        
        // logic
        newL := make([]int, len(path))
        copy(newL, path)
        result = append(result,newL)
        
        for x := i; x < len(nums); x++ {
            // action
            path = append(path, nums[x])
            // recurse
            dfs(x+1,path)
            // backtrack
            path = path[:len(path)-1]
        }
    }
    dfs(0,nil)
    return result
}