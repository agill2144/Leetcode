func countFairPairs(nums []int, lower int, upper int) int64 {
    sort.Ints(nums)
    var count int64 = 0
    for i := 0; i < len(nums); i++ {
        idx1 := closestOnRight(nums,i+1,lower-nums[i])
        idx2 := closestOnLeft(nums, i+1,upper-nums[i])
        if idx1 != -1 && idx2 != -1 {
            count += int64(idx2-idx1+1)
        }
    }
    return count
}

func closestOnRight(nums []int, left int, target int) int{
    right := len(nums)-1
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] >= target {
            ans = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
}

func closestOnLeft(nums []int, left int, target int) int{
    right := len(nums)-1
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] <= target {
            ans = mid
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return ans
}