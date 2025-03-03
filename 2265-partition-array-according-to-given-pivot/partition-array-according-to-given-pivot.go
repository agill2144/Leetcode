func pivotArray(nums []int, pivot int) []int {
    out := []int{}
    n := len(nums)
    equal := 0
    for i := 0; i < n; i++ {
        if nums[i] < pivot {out = append(out, nums[i])}
        if nums[i] == pivot {equal++}
    }
    for equal > 0 {out = append(out, pivot);equal--}
    for i := 0; i < n; i++ {if nums[i] > pivot {out = append(out, nums[i])}}
    return out
}