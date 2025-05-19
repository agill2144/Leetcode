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

// sorted = 1. binary search
// NOT POSSIBLE - WE NEED TO RETURN OG IDX, SORTING FUCKS IT UP ( unless we idx them before, then like .. why do this lol)
// func twoSum(nums []int, target int) []int {
//     sort.Ints(nums)
//     for i := 0; i < len(nums); i++ {
//         t := target-nums[i]
//         left := i+1
//         right := len(nums)-1
//         for left <= right {
//             mid := left + (right-left)/2
//             if nums[mid] == t && mid != i {return []int{i, mid}}
//             if t > nums[mid] {
//                 left = mid+1
//             } else {
//                 right = mid-1
//             }
//         }
//     }
//     return nil
// }


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