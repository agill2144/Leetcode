func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)
    out := [][]int{}
    var dfs func(start int, path []int)
    dfs = func(start int, path []int) {
        // base

        // logic
        newL := make([]int, len(path))
        copy(newL, path)
        out = append(out, newL)
        for i := start; i < len(nums); i++ {
            // action
            path = append(path, nums[i])
            // recurse
            dfs(i+1, path)
            // backtrack
            path = path[:len(path)-1]
            skipElement := nums[i]
            for i+1 < len(nums) && nums[i+1] == skipElement {i++}
        }
    }
    dfs(0, nil)
    return out
}