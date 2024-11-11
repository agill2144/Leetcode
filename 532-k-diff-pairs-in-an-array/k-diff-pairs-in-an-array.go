func findPairs(nums []int, k int) int {
    sort.Ints(nums)
    count := 0
    slow := 0
    fast := 1
    for fast < len(nums) {
        if slow == fast {fast++; continue}
        diff := nums[fast] - nums[slow]
        if diff == k {
            count++
            slow++
            for slow < fast && nums[slow] == nums[slow-1] {slow++}
            fast++
            for fast < len(nums) && nums[fast] == nums[fast-1] {fast++}
        } else if diff > k {
            slow++
        } else {
            fast++
        }
    }
    return count
}
// func findPairs(nums []int, k int) int {
//     set := map[[2]int]struct{}{}
//     idx := map[int]int{}
//     for i := 0; i < len(nums); i++ {
//         c1 := nums[i]-k
//         _, ok := idx[c1]
//         if ok {
//             tmp := [2]int{nums[i], c1}
//             if c1 < nums[i] {tmp = [2]int{c1, nums[i]}}
//             set[tmp] = struct{}{}
//         }
//         c2 := nums[i]+k
//         _, ok = idx[c2]
//         if ok {
//             tmp := [2]int{nums[i], c2}
//             if c2 < nums[i] {tmp = [2]int{c2, nums[i]}}
//             set[tmp] = struct{}{}
//         }
//         idx[nums[i]] = i
//     }
//     return len(set)
// }