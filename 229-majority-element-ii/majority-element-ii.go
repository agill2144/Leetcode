func majorityElement(nums []int) []int {
    c1 := 0
    c2 := 0
    ele1 := math.MinInt64
    ele2 := math.MinInt64
    for i := 0; i < len(nums); i++ {
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
            if c1 < 0 {c1 = 0}
            if c2 < 0 {c2 = 0}
        }
    }

    // verify
    c1, c2 = 0,0
    for i := 0; i < len(nums); i++ {
        if nums[i] == ele1 {c1++}
        if nums[i] == ele2 {c2++}
    }
    out := []int{}
    if c1 >= (len(nums)/3)+1 {out = append(out, ele1)} 
    if c2 >= (len(nums)/3)+1 {out = append(out, ele2)} 
    return out
}