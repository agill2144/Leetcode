func minSwaps(nums []int) int {
    ones := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 1 {ones++}
    }
    if ones == 0 {return 0}
    left := 0
    swaps := math.MaxInt64
    oneCount := 0
    i := 0
    for left < len(nums) {
        ii := i % len(nums)
        if nums[ii] == 1 {oneCount++}
        if i-left+1 == ones {
            zeros := (i-left+1)-oneCount
            swaps = min(swaps, zeros)
            if nums[left] == 1 {oneCount--}
            left++
        }
        i++
    }
    return swaps
}