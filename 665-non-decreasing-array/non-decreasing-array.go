func checkPossibility(nums []int) bool {
    n := len(nums)
    count := 0
    for i := 1; i < n; i++ {
        prev := nums[i-1]
        prevPrev := math.MinInt64
        if i-2 >= 0 {
            prevPrev = nums[i-2]
        }
        curr := nums[i]
        if prev > curr {
            if prevPrev > curr {
                nums[i] = prev
                count++
            } else {
                nums[i-1] = curr 
                count++
            }
        }
        if count == 2 {return false}
    }
    return true
}