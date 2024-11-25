func countFairPairs(nums []int, lower int, upper int) int64 {
    sort.Ints(nums)
    var count int64 = 0
    /*
        0,1,7,4,4,5
        sum >= 3 and <= 6
    */
    for i := 0; i < len(nums); i++ {
        left := leftMostOnRightSideOfTarget(nums, i+1, lower-nums[i])
        right := rightMostOnLeftSideOfTarget(nums,i+1, upper-nums[i])
        if left != -1 && right != -1 {
            count += int64(right-left+1)
        }
    }
    return count    
}


func leftMostOnRightSideOfTarget(nums []int, left int, target int) int {
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

func rightMostOnLeftSideOfTarget(nums []int, left int, target int) int {
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