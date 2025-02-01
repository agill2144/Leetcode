func findMaxLength(nums []int) int {
    idxs := map[int]int{0:-1}
    res := 0
    count := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {count--}
        if nums[i] == 1 {count++}
        left, ok := idxs[count]
        if ok {
            res = max(res, i-left)
        }
        _, ok2 := idxs[count]
        if !ok2 {
            idxs[count] = i
        }
    }
    return res
}