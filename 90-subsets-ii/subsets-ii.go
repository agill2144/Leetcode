// tc = o(2^n * n)
// sc = o(n) + o(n) + o(n)
func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)
    out := [][]int{}
    var dfs func(start int, path []int)
    dfs = func(start int, path []int) {
        // base
        newL := make([]int, len(path))
        copy(newL, path)
        out = append(out, newL)

        // logic
        for i := start; i < len(nums); i++ {
            if i > start && nums[i] == nums[i-1] {continue}
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