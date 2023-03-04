func canJump(nums []int) bool {
    memoIdx := map[int]bool{}
    var dfs func(start int) bool
    dfs = func(start int)bool {
        // base
        if start >= len(nums)-1 {
            if start == len(nums)-1 {return true}
            return false
        }
        
        // logic
        if _, ok := memoIdx[start]; !ok {
            numJumps := nums[start]
            for j := numJumps; j >= 1; j-- {
                if ok := dfs(start+j); ok {return true}
            }
            memoIdx[start] = false
        }
        return memoIdx[start]
    }
    return dfs(0)
}