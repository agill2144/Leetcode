func majorityElement(nums []int) []int {
    n := len(nums)
    ele1 := math.MinInt64
    ele2 := math.MinInt64
    c1 := 0
    c2 := 0
    for i := 0; i < len(nums); i++ {
        if c1 == 0 && nums[i] != ele2 {
            ele1 = nums[i]
            c1 = 1
        } else if c2 == 0 && nums[i] != ele1 {
            ele2 = nums[i]
            c2 = 1
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
    for i := 0; i < len(nums); i++ {
        if nums[i] == ele1 {c1++}
        if nums[i] == ele2 {c2++}
    }
    out := []int{}
    if c1 > n/3 {out = append(out, ele1)}
    if c2 > n/3 {out = append(out, ele2)}
    return out
}
/*
    approach: freq map
    - maintain freq of each num in a freq map
    - then take another pass to collect
    - all elements whose freq > n/3 times
    time = o(2n)
    space = o(n)
*/
// func majorityElement(nums []int) []int {
//     n := len(nums)
//     atMin := (n/3) + 1
//     freq := map[int]int{}
//     for i := 0; i < len(nums); i++ {
//         freq[nums[i]]++
//     }
//     out := []int{}
//     for k,v := range freq {
//         if v >= atMin {out = append(out, k)}
//     }
//     return out
// }