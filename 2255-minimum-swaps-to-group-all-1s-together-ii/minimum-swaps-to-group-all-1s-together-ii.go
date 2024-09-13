func minSwaps(nums []int) int {
    totalOnes := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 1 {totalOnes++}
    }
    if totalOnes <= 1 {return 0}
    minSwaps := len(nums)+1
    left := 0
    count := 0
    n := len(nums)
    for i := 0; i < 2*n && left < n; i++ {
        if nums[i%n] == 1 {
            count++
        }
        if i-left+1 == totalOnes {
            minSwaps = min(minSwaps, totalOnes-count)
            if nums[left] == 1 {
                count--
            }
            left++
        }
    }
    return minSwaps
}