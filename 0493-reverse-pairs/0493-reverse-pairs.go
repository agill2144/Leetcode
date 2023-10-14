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
        
        i := left
        j := mid+1
        for i <= mid {
            for j <= right {
                if nums[i] > nums[j]*2 {
                    count += mid-i+1
                    j++
                } else {
                    break
                }
            }
            i++
        }
        
        merged := []int{}
        l := left
        r := mid+1
        for l <= mid && r <= right {
            if nums[l] <= nums[r] {
                merged = append(merged, nums[l])
                l++
            }  else {
                merged = append(merged, nums[r])
                r++
            }
        }
        for l <= mid {merged = append(merged, nums[l]); l++}
        for r <= right {merged = append(merged, nums[r]); r++}
        ptr := 0
        for i := left; i <= right; i++ {
            nums[i] = merged[ptr]
            ptr++
        }
    }
    dfs(0, len(nums)-1)
    return count
    
}