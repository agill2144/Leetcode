func pivotArray(nums []int, pivot int) []int {
    left := []int{}
    count := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] < pivot {
            left = append(left, nums[i])
        } else if nums[i] == pivot {
            count++
        }
    }
    for count != 0 {
        left = append(left, pivot)
        count--
    }
    for i := 0; i < len(nums); i++ {
        if nums[i] > pivot {
            left = append(left, nums[i])
        }
    }    
    return left
}