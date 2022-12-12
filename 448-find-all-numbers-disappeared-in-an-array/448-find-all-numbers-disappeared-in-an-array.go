func findDisappearedNumbers(nums []int) []int {
    for i := 0; i < len(nums); i++ {
        num := abs(nums[i])
        idx := num-1
        if nums[idx] > 0 {
            nums[idx] *= -1
        }
    }
    result := []int{}
    for idx, num := range nums {
        if num > 0 {
            result = append(result, idx+1)
        }
    }
    return result
}

func abs(n int) int {
    if n < 0 {
        return n * -1
    }
    return n
}