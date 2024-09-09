func findMin(nums []int) int {
    left := 0
    right := len(nums)-1
    ans := math.MaxInt64
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == nums[right] && nums[right] == nums[left] {
            // all numbers are equal
            // save mid if mid is a better answer
            // move away from both since all values are same
            // it means we cannot decide which side to move away from
            // therefore shrink our space and re-evaluate
            ans = min(nums[mid], ans)
            left++
            right--
        } else if nums[left] <= nums[mid] {
            // left sorted
            // mid is bigger than left
            // save left if left is a better answer
            ans = min(ans, nums[left])
            left = mid+1
        } else  {
            // right sorted ( right is bigger than mid )
            // save if mid is a better answer
            ans = min(ans, nums[mid])
            right = mid-1
        }
    }
    return ans
}