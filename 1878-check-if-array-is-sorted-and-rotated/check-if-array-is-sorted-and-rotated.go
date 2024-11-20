func check(nums []int) bool {
    n := len(nums)
    if n <= 1 {return true}
    idx := -1
    for i := 0; i < n; i++ {
        if i == 0 {continue}
        if nums[i] < nums[i-1] {idx = i; break}
    }
    if idx == -1 {return true}
    start := idx+1
    end := idx
    for (start%n) != end {
        if nums[start%n] >=  nums[(start-1)%n] {start++} else {return false}
    }
    return true
}