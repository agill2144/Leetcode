func countSubarrays(nums []int, k int64) int64 {
    if nums == nil || len(nums) == 0 || k <= 0 {return 0}
    var count int64 = 0    
    sum := 0
    prod := 1
    left := 0
    for right := 0; right < len(nums); right++ {
        numElements := right-left+1
        sum += nums[right]
        prod = sum * numElements
        for int64(prod) >= k {
            leftVal := nums[left]
            sum -= leftVal
            left++
            prod = sum * (right-left+1)
        }
        count += int64(right-left+1)
    }
    return count
}