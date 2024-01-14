
func splitArray(nums []int, k int) int {
    if k > len(nums) {return -1}
    // that the largest sum of any subarray is minimized
    // max ans is minized 
    // meaning there multiple placements of split to make k total splits
    // after each placement, we want the largest from each placment
    // then, from all the placements, we want the smallest such sum
    
    start := math.MinInt64
    end := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] > start {start = nums[i]}
        end += nums[i]
    }

    ans := 0
    left := start
    right := end
    for left <= right {
        mid := left + (right-left) / 2
        
        maxSumAllowed := mid
        countSubArr := 1
        rSum := 0
        for j := 0; j < len(nums); j++ {
            rSum += nums[j]
            if rSum > maxSumAllowed {
                countSubArr++
                rSum = nums[j]
            }
        }
        
        if countSubArr <= k {
            ans = mid
            right = mid-1
        } else {
            left = mid+1
        }

    }
    
    return ans
}