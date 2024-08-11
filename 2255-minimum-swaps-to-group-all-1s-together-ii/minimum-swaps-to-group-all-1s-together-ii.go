func minSwaps(nums []int) int {
    n := len(nums)
    ones := 0
    for i := 0; i < n; i++ {if nums[i] == 1 {ones++}}
    if ones == 0 {return 0}
    left := 0
    zeros := 0
    swaps := n+1
    i := 0
    for left < n {
        if nums[i%n] == 0 {zeros++}
        size := i-left+1
        if size == ones {
            swaps = min(swaps, zeros)
            if nums[left] == 0 {zeros--}
            left++
        }
        i++
    }
    return swaps
}