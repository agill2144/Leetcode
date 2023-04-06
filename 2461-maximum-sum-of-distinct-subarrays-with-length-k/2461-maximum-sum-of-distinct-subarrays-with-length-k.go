func maximumSubarraySum(nums []int, k int) int64 {
    numIdx := map[int64]int{}
    var max int64 = 0
    var sum int64 = 0
    left := 0
    for i := 0; i < len(nums); i++ {
        num := int64(nums[i])
        sum += num

        // is this num unique in current subarray
        idxLastSeenAt, ok := numIdx[num]
        
        if ok || i-left+1 > k {
            for (i-left+1 > k) || (left <= idxLastSeenAt)  {
                sum -= int64(nums[left])
                left++
            }            
        }
        numIdx[num] = i
        if i-left+1 == k {
            // fmt.Println(nums[left:i+1], sum)
            if sum > max {max = sum}
        }
    }
    return max
}