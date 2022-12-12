func findNumbers(nums []int) int {
    if nums == nil || len(nums)==0 {
        return 0
    }
    count := 0
    for i := 0; i < len(nums); i++ {
        str := strconv.Itoa(nums[i])
        if len(str) % 2 == 0 {
            count++
        }
    }
    return count
}