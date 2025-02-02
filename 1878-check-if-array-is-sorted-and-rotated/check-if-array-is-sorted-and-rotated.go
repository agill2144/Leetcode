func check(nums []int) bool {
    startIdx := -1
    // find the dipping point
    // i.e start of sorted array
    for i := 0; i < len(nums)-1; i++ {
        nextIdx := i+1
        if nums[i] > nums[nextIdx] {
            startIdx = nextIdx
            break
        }
    }
    if startIdx == -1 {return true}
    // now count how many elements successfully processed
    // as in, count number of curr elements seen 
    // as long as curr and next element are sorted
    // once curr has seen all, we are good, things are sorted
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