func smallestDivisor(nums []int, threshold int) int {
    left := 1
    right := 0
    for i := 0; i < len(nums); i++ {if nums[i] > right {right = nums[i]}}

    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        
        sum := 0
        for k := 0; k < len(nums); k++{
            sum += int(math.Ceil(float64(nums[k])/float64(mid)))            
        }
        if sum <= threshold {
            ans = mid
            right = mid-1
        } else {
            left = mid+1
        }        
    }
    return ans
}