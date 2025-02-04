func minimumSize(nums []int, maxOps int) int {
    left := 1
    right := slices.Max(nums)
    if maxOps <= 0 {return right}
    res := -1
    for left <= right {
        mid := left + (right-left)/2
        atMax := mid
        ops := 0
        for i := 0; i < len(nums); i++ {
            ops += int(math.Ceil(float64(nums[i])/float64(atMax)))-1
            if ops > maxOps {break}
        }
        if ops <= maxOps {
            res = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return res
}