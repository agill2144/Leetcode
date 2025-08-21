func reversePairs(nums []int) int {
    count := 0
    var dfs func(left, right int)
    dfs = func(left, right int) {
        // base
        if left >= right {return}        

        // logic
        mid := left + (right-left)/2
        dfs(left, mid)
        dfs(mid+1, right)

        i, j := left, mid+1
        for i <= mid && j <= right {
            if nums[i] > 2 * nums[j] {
                count += (mid-i+1)
                j++
            } else {
                i++
            }
        }

        i, j = left, mid+1
        merged := []int{}
        for i <= mid && j <= right {
            if nums[i] < nums[j] {
                merged = append(merged, nums[i])
                i++
            } else {
                merged = append(merged, nums[j])
                j++
            }
        }
        for i <= mid {merged = append(merged, nums[i]); i++}
        for j <= right {merged = append(merged, nums[j]); j++}
        ptr := 0
        for i := left; i <= right; i++ {
            nums[i] = merged[ptr]
            ptr++
        }
    }
    dfs(0, len(nums)-1)
    return count
}