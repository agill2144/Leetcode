func findDisappearedNumbers(nums []int) []int {
    if nums == nil || len(nums) == 0 {return nil}
    for i := 0; i < len(nums); i++ {
        num := abs(nums[i])
        idx := num-1
        if nums[idx] < 0 {continue}
        nums[idx] *= -1
    }
    out := []int{}
    for i := 0; i < len(nums); i++ {
        if nums[i] > 0 {
            out = append(out, i+1)
        }
    }
    return out
}

func abs(n int) int {
    if n < 0 {return n *-1}
    return n
}