func canJump(nums []int) bool {
    n := len(nums)
    memo := map[int]*bool{}
    var dfs func(ptr int)bool
    dfs = func(ptr int)bool {
        if ptr >= n-1 {return true}

        // already solved idx
        if memo[ptr] != nil {return *memo[ptr]}
        jumpSize := nums[ptr]

        // when jump size is 0, not possible to proceed,
        // save false and return false
        if jumpSize == 0 {
            memo[ptr] = toBoolPtr(false)
            return *memo[ptr]
        }

        var res bool
        for i := jumpSize; i >= 1; i-- {
            res = dfs(ptr+i)
            // as soon as we get a true, break
            if res {break}
        }
        // save the res
        memo[ptr] = toBoolPtr(res)
        // return the res
        return *memo[ptr]
    }
    return dfs(0)
}

func toBoolPtr(val bool) *bool {
    return &val
}