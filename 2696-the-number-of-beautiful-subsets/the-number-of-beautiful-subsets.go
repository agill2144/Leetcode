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

func abs(x int) int {
    if x < 0 {return x *-1}
    return x
}