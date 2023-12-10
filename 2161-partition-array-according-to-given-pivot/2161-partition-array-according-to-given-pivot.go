func pivotArray(nums []int, pivot int) []int {
    out := []int{}
    countEq := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] < pivot {
            out = append(out, nums[i])
        } else if nums[i] == pivot {
            countEq++
        }
    }
    for countEq > 0 {
        out = append(out, pivot)
        countEq--
    }
    
    for i := 0; i < len(nums); i++ {
        if nums[i] > pivot {
            out = append(out, nums[i])
        }
    }
    return out
}