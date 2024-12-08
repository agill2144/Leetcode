func maximumSwap(num int) int {
    idxs := make([]int, 10)
    for i := 0; i < len(idxs); i++ {idxs[i] = -1}
    nums := strings.Split(fmt.Sprintf("%v", num), "")
    for i := 0; i < len(nums); i++ {
        numInt, _  := strconv.Atoi(nums[i])
        idxs[numInt] = i
    }
    swapped := false
    for i := 0; i < len(nums) && !swapped; i++ {
        curr, _ := strconv.Atoi(nums[i])
        for j := len(idxs)-1; j > curr; j-- {
            if idxs[j] != -1 && idxs[j] > i && j > curr {
                nums[i], nums[idxs[j]] = nums[idxs[j]], nums[i]
                swapped = true
                break
            }
        }
    }
    out := 0
    for i := 0; i < len(nums); i++ {
        n, _ := strconv.Atoi(nums[i])
        out = out * 10 +n
    }
    return out
}