func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)
    ans := [][]int{}
    var dfs func(start int, path []int)
    dfs = func(start int, path []int) {
        // base
        newP := make([]int, len(path))
        copy(newP, path)
        ans = append(ans, newP)
        // logic
        for i := start; i < len(nums); i++ {
            // action
            path = append(path, nums[i])
            // recurse
            dfs(i+1, path)
            // backtrack
            fmt.Println("back from: ", path)
            path = path[:len(path)-1]
            fmt.Println("to: ", path)
            
            // hidden: not choose case here
            skipVal := nums[i]
            for i+1 < len(nums) && nums[i+1] == skipVal {
                fmt.Println("skipCheck: ",skipVal, nums[i])
                i++
            }
        }
    }
    dfs(0, nil)
    return ans
}