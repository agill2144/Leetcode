/*
    approach: brute force
    - try all possible pais
    time = o(n^2)
    space = o(1)
*/
func countFairPairs(nums []int, lower int, upper int) int64 {
    sort.Ints(nums)
    var count int64
    for i := 0; i < len(nums); i++ {
        // identify the range
        leftIdx := leftSearch(nums, i+1, lower-nums[i])
        rightIdx := rightSearch(nums, i+1, upper-nums[i])
        if rightIdx != -1 && leftIdx != -1 {
            count += (int64(rightIdx) - int64(leftIdx)+1)
        }

    }
    return count
}

func leftSearch(nums []int, left, target int) int {
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


func rightSearch(nums []int, left, target int) int {
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
