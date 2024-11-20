func maximumSwap(num int) int {
    if num <= 9 {return num}
    nums := strings.Split(fmt.Sprintf("%v", num), "")
    n := len(nums)
    suffix := make([]int, len(nums)) // idxs
    // 2736
    // ["2", "7", "3", "6"]
    // [ 1    3    3    3]
    suffix[n-1] = n-1
    for i := n-2; i >= 0; i-- {
        maxIdx := suffix[i+1]
        maxSoFar := nums[maxIdx]
        rightIdx := i+1
        rightVal := nums[rightIdx]
        if rightVal > maxSoFar {
            suffix[i] = rightIdx
        } else {
            suffix[i] = maxIdx
        }
    }
    for i := 0; i < n; i++ {
        maxOnRight := nums[suffix[i]]
        if maxOnRight > nums[i] {
            nums[i], nums[suffix[i]] = nums[suffix[i]], nums[i]
            break
        }
    }
    res := 0
    for i := 0; i < len(nums); i++ {
        nu, _ := strconv.Atoi(nums[i])
        res = (res * 10) + nu
    }
    return res
}