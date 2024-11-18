func findClosestElements(nums []int, k int, x int) []int {
    left := 0
    right := len(nums)-1
    for right-left+1 != k {
        leftVal := abs(nums[left]-x)
        rightVal := abs(nums[right]-x)
        if leftVal <= rightVal {
            right--
        } else {
            left++
        }
    }
    return nums[left:right+1]
}

func abs(x int) int {
    if x < 0 {return x *- 1}
    return x
}