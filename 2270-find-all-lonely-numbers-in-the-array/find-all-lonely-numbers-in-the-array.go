func findLonely(nums []int) []int {
    sort.Ints(nums)
    out := []int{}
    for i := 0; i < len(nums); i++ {
        curr := nums[i]
        prev := math.MinInt64+10
        if i-1 >= 0 {prev = nums[i-1]}
        next := math.MaxInt64-10
        if i+1 < len(nums) {next = nums[i+1]}
        if curr == prev || curr == next || curr-1 == prev || curr+1 == next {continue}
        out = append(out,curr)
    }
    return out
}
// func findLonely(nums []int) []int {
//     freq := map[int]int{}
//     for i := 0; i < len(nums); i++ {
//         freq[nums[i]]++
//     }
//     out := []int{}
//     for k, v := range freq {
//         if v > 1 {continue}
//         if freq[k-1] == 0 && freq[k+1] == 0 {
//             out = append(out, k)
//         }
//     }
//     return out
// }