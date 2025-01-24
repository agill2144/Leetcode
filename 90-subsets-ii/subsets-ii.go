// tc = o(nlogn) + o(2^n * n)
// sc = o(2n) ; at the deepest layer of recursion, our recursion stack is n and path is n
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
            path = append(path, nums[i])
            dfs(i+1, path)
            path = path[:len(path)-1]
            for i+1 < len(nums) && nums[i+1] == nums[i] {i++}
        }
    }
    dfs(0,nil)
    return out
}