func checkSubarraySum(nums []int, k int) bool {
    remaindersIdx := map[int]int{0:-1}
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        rem := sum % k
        if rem < 0 {rem += k}
        val, ok := remaindersIdx[rem]
        if ok && i-(val+1)+1 >= 2 {
            return true
        }
        if !ok {
            remaindersIdx[rem] = i
        }
    }
    return false
}

/*
    approach : brute force
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
//             if sum % k == 0 && j-i+1 >= 2 {return true}
//         }
//     }
//     return false
// }