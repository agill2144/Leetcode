func findDuplicates(nums []int) []int {
    out := []int{}
    for i := 0; i < len(nums); i++ {
        val := abs(nums[i])
        idx := val-1
        if nums[idx] < 0 {out = append(out, val); continue}
        nums[idx] *= -1
    }
    return out
}

func abs(x int) int {
    if x < 0 {return x * -1}
    return x
}