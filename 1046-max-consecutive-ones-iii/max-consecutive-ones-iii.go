// we are allowed to have atleast k zeros in our window
func longestOnes(nums []int, k int) int {
    maxSize := 0
    n := len(nums)
    left := 0
    zeros := 0
    for i := 0; i < n; i++ {
        if nums[i] == 0 {zeros++}
        for zeros > k {
            if nums[left] == 0 {zeros--}
            left++
        }
        maxSize = max(maxSize, i-left+1)
    }
    return maxSize
}