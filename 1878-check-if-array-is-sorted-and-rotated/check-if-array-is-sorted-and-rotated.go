func check(nums []int) bool {
    n := len(nums)
    end := -1
    for i := 0; i < 2*n; i++ {
        if nums[i%n] > nums[(i+1)%n] {end = i}
    }
    if end == -1 {return true}
    start := end+1
    for start%n != end%n {
        curr := nums[start % n ]
        next := nums[(start+1) % n]
        if curr > next {return false}
        start++
    }
    return true
}