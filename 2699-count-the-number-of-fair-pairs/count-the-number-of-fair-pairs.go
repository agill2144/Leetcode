func countFairPairs(nums []int, lower int, upper int) int64 {
    sort.Ints(nums)
    var count int64 
    for i := 0; i < len(nums); i++ {
        left := leftMostOnRightSide(nums, lower-nums[i], i+1)
        right := rightMostOnLeftSide(nums, upper-nums[i],i+1)
        if left != -1 && right != -1 {
            count += int64(right-left+1)
        }
    }
    return count
}

func leftMostOnRightSide(nums []int, target, left int) int {
    idx := -1
    right := len(nums)-1
    for left <= right {
        mid := left+(right-left)/2
        if nums[mid] >= target {
            idx = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return idx
}

func rightMostOnLeftSide(nums []int, target, left int) int {
    idx := -1
    right := len(nums)-1
    for left <= right {
        mid := left+(right-left)/2
        if nums[mid] <= target {
            idx = mid
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return idx
}