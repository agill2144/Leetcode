func check(nums []int) bool {
    startIdx := -1
    for i := 0; i < len(nums)-1; i++ {
        nextIdx := i+1
        if nums[i] > nums[nextIdx] {
            startIdx = nextIdx
            break
        }
    }
    if startIdx == -1 {return true}
    count := 1
    ptr := startIdx
    n := len(nums)
    for count != n {
        curr := nums[ptr%n]
        next := nums[(ptr+1)%n]
        if curr > next {return false}
        count++
        ptr++
    }
    return true
}