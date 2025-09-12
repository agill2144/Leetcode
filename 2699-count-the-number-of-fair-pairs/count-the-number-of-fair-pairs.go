func countFairPairs(nums []int, lower int, upper int) int64 {
    n := len(nums)
    sort.Ints(nums)
    var count int64
    for i := 0; i < n; i++ {
        left := searchGreaterOrEqualTo(nums, i+1, lower-nums[i])
        right := searchLessThanOrEqualTo(nums, i+1, upper-nums[i])
        if left != -1 && right != -1 {
            count += int64(right-left+1)
        }
    }
    return count
}


func searchGreaterOrEqualTo(nums []int, left int, target int) int {
    ans := -1
    right := len(nums)-1
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

func searchLessThanOrEqualTo(nums []int, left int, target int) int {
    ans := -1
    right := len(nums)-1
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