func maximumSwap(num int) int {
    if num <= 9 {return num}
    nums := []int{}
    numStr := fmt.Sprintf("%v", num)
    for i := 0; i < len(numStr); i++ {nums = append(nums, int(numStr[i]-'0'))}
    n := len(nums)
    suffix := make([]int, len(nums)) 
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
        res = (res * 10) + nums[i]
    }
    return res
}