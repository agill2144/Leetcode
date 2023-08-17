// binary search the complement
// time = o(n) * o(logN)
// space = o(1)
func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        diff := target - nums[i]
        val := binarySearch(diff, i+1, nums)
        if val != i && val != -1 {
            if i < val {
                return []int{i+1, val+1}            
            } else {
                return []int{val+1, i+1}                            
            }
        }
    }
    return nil

}


func binarySearch(target int, left int, nums []int) int {
    right := len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            return mid
        } else if target > nums[mid] {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return -1
}

