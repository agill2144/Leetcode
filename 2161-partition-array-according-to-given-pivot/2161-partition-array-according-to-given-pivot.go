func pivotArray(nums []int, pivot int) []int {
    left := []int{}
    right := []int{}
    count := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] < pivot {
            left = append(left, nums[i])
        } else if nums[i] > pivot {
            right = append(right, nums[i])
        } else {
            count++
        }
    }
    for count != 0 {
        left = append(left, pivot)
        count--
    }
    left = append(left, right...)
    return left
}