func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        tg := target-nums[i]
        if idx := search(i+1,tg, nums); idx != -1 {
            return []int{i+1, idx+1}
        }
    }
    return nil   
}

func search(left,target int, nums []int) int {
    right := len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {return mid}
        if target > nums[mid] {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return -1
}