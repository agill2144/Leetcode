func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)
    out := [][]int{}
    var dfs func(start int, path []int)
    dfs = func(start int, path []int) {
        // base
        newP := make([]int, len(path))
        copy(newP, path)
        out = append(out, newP)
        if start == len(nums) {return}

        // logic
        for i := start; i < len(nums); i++ {
            path = append(path, nums[i])
            dfs(i+1, path)
            path = path[:len(path)-1]

            for i+1 < len(nums) && nums[i] == nums[i+1] {i++}
        }
    }
    dfs(0,nil)
    return out
}