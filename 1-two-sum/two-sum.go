// compliment search using hashing
// tc = o(n)
// sc = o(n)
func twoSum(nums []int, target int) []int {
    idx := map[int]int{}
    for i := 0; i < len(nums); i++ {
        diff := target-nums[i]
        if val, ok := idx[diff]; ok {return []int{val, i}}
        idx[nums[i]] = i
    }
    return nil
}
// brute force - n^2
// using nested iterations, check all pairs
// func twoSum(nums []int, target int) []int {
//     n := len(nums)
//     for i := 0; i < n; i++ {
//         for j := i+1; j < n; j++ {
//             if nums[i] + nums[j] == target {
//                 return []int{i, j}
//             } 
//         }
//     }
//     return nil
// }