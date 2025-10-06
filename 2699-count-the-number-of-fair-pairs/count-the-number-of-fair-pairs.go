func countFairPairs(nums []int, lower int, upper int) int64 {
    sort.Ints(nums)
    count := 0
    for i := 0; i < len(nums); i++ {
        leftIdx := leftSearch(nums, i+1, lower-nums[i])
        rightIdx := rightSearch(nums, i+1, upper-nums[i])
        if leftIdx != -1 && rightIdx != -1 {
            count += (rightIdx-leftIdx+1)
        }
    }
    return int64(count)
}

func leftSearch(nums []int, left int, target int )int {
    idx := -1
    right := len(nums)-1
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

func rightSearch(nums []int, left int, target int )int {
    idx := -1
    right := len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] <= target {
            idx = mid
            left = mid+1
        }  else {
            right = mid-1
        }
    }
    return idx
}