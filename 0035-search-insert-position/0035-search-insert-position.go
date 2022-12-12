/*
    approach: binary search
    - look for the next but smallest number after target
    - look for the ceiling number of target
    
    time : o(logn)
    space: o(1)
*/
func searchInsert(nums []int, target int) int {
    if nums[len(nums)-1] < target {return len(nums)}
    if nums[0] > target {return 0}
    left := 0
    right := len(nums)-1
    insert := -1
    
    for left <= right {
        mid := left+(right-left)/2
        if nums[mid] >= target {
            if nums[mid] == target {return mid}
            insert = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return insert
}