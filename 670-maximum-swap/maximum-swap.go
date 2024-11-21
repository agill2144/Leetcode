/*
    bring the largest value in smallest unit to the largest unit
    example
    2139
    9 is at the smallest unit (tens)
    and now we need a largest unit whose value is smaller than largest value (from left -> right)
    2 is smaller than 9 and is also in the largest unit ( thousands unit )
    therefore swap 2 and 9 = 9132
*/
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