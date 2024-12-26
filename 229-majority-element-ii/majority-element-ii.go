func majorityElement(nums []int) []int {
    c1, c2 := 0, 0
    e1, e2 := math.MinInt64, math.MinInt64
    n := len(nums)
    for i := 0; i < n; i++ {
        if c1 == 0 && nums[i] != e2 {
            c1 = 1
            e1 = nums[i]
        } else if c2 == 0 && nums[i] != e1 {
            c2 = 1
            e2 = nums[i]
        } else if nums[i] == e1 {
            c1++
        } else if nums[i] == e2 {
            c2++
        } else {
            c1--
            c2--
        }
    }
    c1 = 0
    c2 = 0
    for i := 0; i < n; i++ {
        if nums[i] == e1 {c1++}
        if nums[i] == e2 {c2++}
    }
    out := []int{}
    if c1 > n/3 {out = append(out, e1)}
    if c2 > n/3 {out = append(out, e2)}
    return out
}