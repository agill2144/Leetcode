
func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    out := [][]int{}
    for i := 0; i < len(nums); i++ {
        if i != 0 && nums[i] == nums[i-1] {continue}
        left := i+1
        right := len(nums)-1
        for left < right {
            sum := nums[i] + nums[left] + nums[right]
            if sum == 0 {
                out = append(out, []int{nums[i], nums[left], nums[right]})
                left++
                for left < right && nums[left] == nums[left-1] {left++}
                right-- 
                for left < right && nums[right] == nums[right+1] {right--}
            } else if sum < 0 {
                left++
            } else {
                right--
            }
        }
    }
    return out
}

/*
    approach: sort + binary search
    - continue using a set to ensure unique triplets
    time = o(n) * o(nlogn) = o(n^2logn)
    space = o(k) unique triplets in set
*/

// func threeSum(nums []int) [][]int {
//     sort.Ints(nums)
//     out := [][]int{}
//     set := map[[3]int]struct{}{}
//     for i := 0; i < len(nums); i++ { // o(n)
//         target := 0-nums[i]
//         for j := i+1; j < len(nums); j++ { // o(n)
//             k := binarySearch(nums, target-nums[j], j+1) // o(logn)
//             if k != -1 {
//                 tmp := []int{nums[i], nums[j], nums[k]}
//                 if _, ok := set[[3]int{tmp[0], tmp[1], tmp[2]}]; !ok {
//                     set[[3]int{tmp[0], tmp[1], tmp[2]}] = struct{}{}
//                     out = append(out, tmp)
//                 }
//             }
//         }
//     }
//     return out
// }

// func binarySearch(nums []int, target, left int) int {
//     right := len(nums)-1
//     for left <= right {
//         mid := left + (right-left)/2
//         if nums[mid] == target {
//             return mid
//         } else if target > nums[mid] {
//             left = mid+1
//         } else {
//             right = mid-1
//         }
//     }
//     return -1
// }

/*
    approach: better brute force
    - given ith value, how do we make this ith value == 0
    - 0-ith value, this is our new target
    - now we need to find a pair of 2 numbers that add up to target
    - use a linear complement search using a hashmap ( two sum )
    
    time = o(n^2)
    space = o(k) + o(n)
    - k uniq triplets
    - n for the nested idxMap for complement search
*/
// func threeSum(nums []int) [][]int {
//     set := map[[3]int]struct{}{}
//     out := [][]int{}
//     for i := 0; i < len(nums); i++ {
//         target := 0-nums[i]
//         idxMap := map[int]int{}
//         for j := i+1; j < len(nums); j++ {
//             complement := target-nums[j]
//             idx, ok := idxMap[complement]
//             if ok {
//                 tmp := []int{nums[i], nums[idx], nums[j]}
//                 sort.Ints(tmp)
//                 if _, ok := set[[3]int{tmp[0], tmp[1], tmp[2]}]; !ok {
//                     set[[3]int{tmp[0], tmp[1], tmp[2]}] = struct{}{}
//                     out = append(out, tmp)
//                 }
//             } else {
//                 idxMap[nums[j]] = j
//             }
//         }
//     }
//     return out
// }

/*
    approach: brute force
    - check all triplets
    - using nested iterations
    - use set to keep track of unique triplets ( sort the triplets first )
    
    time = o(n^3)
    space = o(k) k uniq triplets
*/
// func threeSum(nums []int) [][]int {
//     set := map[[3]int]struct{}{}
//     out := [][]int{}
//     for i := 0; i < len(nums); i++ {
//         for j := i+1; j < len(nums); j++ {
//             for k := j+1; k < len(nums); k++ {
//                 if nums[i]+nums[j]+nums[k] == 0 {
//                     tmp := []int{nums[i],nums[j],nums[k]}
//                     sort.Ints(tmp)
//                     if _, ok := set[[3]int{tmp[0], tmp[1], tmp[2]}]; !ok {
//                         set[[3]int{tmp[0], tmp[1], tmp[2]}] = struct{}{}
//                         out = append(out, tmp)
//                     }
//                 }
//             }
//         }
//     }
//     return out
// }    