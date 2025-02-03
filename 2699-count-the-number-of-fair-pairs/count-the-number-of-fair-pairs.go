func countFairPairs(nums []int, lower int, upper int) int64 {
    n := len(nums)
    sort.Ints(nums)
    var count int64
    for i := 0; i < n; i++ {
        left := searchGreaterEqualTo(nums, i+1, lower-nums[i])
        right := searchLessEqualTo(nums, i+1, upper-nums[i])
        if left != -1 && right != -1 {
            count += int64(right-left+1)
        }
    }
    return count
}

func searchGreaterEqualTo(nums []int, left int, target int) int {
    n := len(nums)
    idx := -1
    right := n-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] >= target {
            idx = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return idx
}


func searchLessEqualTo(nums []int,left int,target int) int {
    n := len(nums)
    idx := -1
    right := n-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] <= target {
            idx= mid
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return idx
}