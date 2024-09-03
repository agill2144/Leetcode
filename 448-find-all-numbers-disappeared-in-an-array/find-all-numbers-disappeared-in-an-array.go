func findDisappearedNumbers(nums []int) []int {
    for i := 0; i < len(nums); i++ {
        idx := abs(nums[i])-1
        if nums[idx] < 0 {continue}
        nums[idx] *= -1
    }
    out := []int{}
    for i := 0; i < len(nums); i++ {
        if nums[i] > 0 {out = append(out, i+1)}
    }
    return out
}

func abs(x int) int {
    if x < 0 {return x *-1}
    return x
}