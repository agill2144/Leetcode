func checkSubarraySum(nums []int, k int) bool {
    remIdx := map[int]int{}
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        rem := sum % k
        idx, ok := remIdx[rem]
        if (ok && i-(idx+1)+1 >= 2) || (rem == 0 && i >= 1) {
            return true
        }
        if !ok {
            remIdx[rem] = i
        }
    }
    return false
}
/*
    approach : brute force
    - We need a subarray sum divisible by k
    - Which means we need to maintain sum
    - but this is running sum from 0 to some x idx
    - and it may or may not be divisible by k
    - But what about the other subarrays between ( not including ) idx 0 and idx x ?
    - how do we find those ?
    - in brute force, we used nested iterations to form every single subarray
    - using nested iterations, form all possible subarray
    - and check if its a good subarray    
    time = o(n^2)
    space = o(1)
*/
// func checkSubarraySum(nums []int, k int) bool {
//     for i := 0; i < len(nums); i++ {
//         sum := 0
//         for j := i; j < len(nums); j++ {
//             sum += nums[j]
//             if sum % k == 0 && j-i+1 >= 2 {
//                 return true
//             }
//         }
//     }
//     return false
// }