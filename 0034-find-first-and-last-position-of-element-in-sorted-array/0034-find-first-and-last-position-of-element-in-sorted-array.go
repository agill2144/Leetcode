func searchRange(nums []int, target int) []int {
    if nums == nil || len(nums) == 0 {
        return []int{-1,-1}
    }
    left := 0
    right := len(nums)-1
    start := -1
    
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] >= target {
            if nums[mid] == target {
                start = mid
            }
            right = mid-1
        } else if nums[mid] < target {
            left = mid+1
        }
    }
    if start == -1 {return []int{-1,-1}}
    
    l := start
    r := len(nums)-1
    end := -1    
    for l <= r {
        m := l+(r-l)/2
        if nums[m] <= target {
            if nums[m] == target {
                end = m
            }
            l= m+1
        } else {
            r = m - 1
        }
    }
    return []int{start, end}
}