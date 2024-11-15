func check(nums []int) bool {
    idx := -1
    for i := 0; i < len(nums); i++ {
        if i == 0 {continue}
        if nums[i] < nums[i-1] {idx = i; break}
    }
    if idx == -1 {return true}
    n := len(nums)
    start := idx+1
    for start % n != idx {
        if nums[(start-1)%n] <= nums[start%n] {start++} else {break}
    }
    return (start%n) == idx
}