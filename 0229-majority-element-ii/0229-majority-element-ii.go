func majorityElement(nums []int) []int {
    count := 0
    ele := math.MinInt64
    count2 := 0
    ele2 := math.MinInt64
    for i := 0; i < len(nums); i++ {
        if count == 0 && nums[i] != ele2 {
            count = 1
            ele = nums[i]
        } else if count2 == 0 && nums[i] != ele {
            count2 = 1
            ele2 = nums[i]
        } else if nums[i] == ele {
            count++
        } else if nums[i] == ele2 {
            count2++
        } else {
            count--
            count2--
        }
    }

    out := []int{}
    t := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == ele {t++}
    }
    if t > len(nums)/3 {out = append(out, ele)}
    t = 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == ele2 {t++}
    }
    if t > len(nums)/3 {out = append(out, ele2)}
    return out
}