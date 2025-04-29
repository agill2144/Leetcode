func majorityElement(nums []int) []int {
    ele1, ele2 := math.MinInt64, math.MinInt64
    c1, c2 := 0,0
    n := len(nums)
    for i := 0; i < n; i++ {
        if c1 == 0 && nums[i] != ele2 {
            c1 = 1
            ele1 = nums[i]
        } else if c2 == 0 && nums[i] != ele1 {
            c2 = 1
            ele2 = nums[i]
        } else if nums[i] == ele1 {
            c1++
        } else if nums[i] == ele2 {
            c2++
        } else {
            c1--
            c2--
        }
    }
    c1 = 0
    c2 = 0
    for i := 0; i < n; i++ {
        if nums[i] == ele1 {c1++}
        if nums[i] == ele2 && ele2 != ele1 {c2++}
    }
    out := []int{}
    if c1 > n/3 {out = append(out, ele1)}    
    if c2 > n/3 {out = append(out, ele2)}    
    return out
}