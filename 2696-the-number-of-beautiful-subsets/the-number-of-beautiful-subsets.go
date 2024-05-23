// use for loop based recursion to form each subset
// use complement search to find out whether a incoming ith number
// can be part of our path so far
func beautifulSubsets(nums []int, k int) int {
    count := 0
    var dfs func(start int, path map[int]int)
    dfs = func(start int, path map[int]int) {
        // base
        if len(path) > 0 {
            count++
        }

        // logic
        for i := start; i < len(nums); i++ {
            c1 := nums[i] + k
            c2 := nums[i] - k
            _, ok := path[c1]
            _, ok2 := path[c2]
            if !ok && !ok2 {
                path[nums[i]]++
                dfs(i+1, path)
                path[nums[i]]--
                if path[nums[i]] == 0 {
                    delete(path, nums[i])
                }
            }
        }
    }
    dfs(0, map[int]int{})
    return count
}