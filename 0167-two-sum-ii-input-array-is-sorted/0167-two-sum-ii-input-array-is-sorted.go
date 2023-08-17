// binary search the complement
func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        diff := target - nums[i]
        // fmt.Println("searching for: ", diff)
        val := binarySearch(diff, nums)
        if val != i && val != -1 {
            // fmt.Println("found at idx: ", val, i)
            if i < val {
                return []int{i+1, val+1}            
            } else {
                return []int{val+1, i+1}                            
            }
        }
    }
    return nil

}


func binarySearch(target int , nums []int) int {
    left := 0
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

