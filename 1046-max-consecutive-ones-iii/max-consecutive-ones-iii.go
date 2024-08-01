func longestOnes(nums []int, k int) int {
    zerosDq := []int{}
    maxWin := 0
    left := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {
            zerosDq = append(zerosDq, i)
        }
        if len(zerosDq) > k {
            if left <= zerosDq[0] {
                left = zerosDq[0] + 1
                zerosDq = zerosDq[1:]
            }
        }
        maxWin = max(maxWin, i-left+1)
    }
    return maxWin
}