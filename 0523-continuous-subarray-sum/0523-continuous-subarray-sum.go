/*
    approach: running sum / prefix sum
    - to get rid of nested iterations, consider; binarySearch, hashing, DP, SlidingWindow, TwoPtrs, RunningSum
    - This problem sounds closer to runningSum as we do not have a definitive as to when to shrink our sliding window
    - at each runningSum, we will be checking if this runningSum is divisible by k
    - If the remainder is 0, great, it means from idx 0 to this idx there is a subarray which is divisble by k
        - double check the size >= 2, and return if true
    - but what if the remainder is not 0 or the size is < 2
    - Then we need to look for whether we have seen this remainder before at some previous idx runningSum
    - Why ?
    - Because turns out that the only time we seen duplicate remainders AGAIN and AGAIN is when;
        - there was some number added to the runningSum which is divisible by k
    - for eg, if the remainder at idx 2 was 7
        - and the remainder at idx 14 is again 7
        - it means from idx 3, the sum added that made to idx 14 is a number that IS divisible by k
    - The only time when we get same remainder when dividing by k is if some number was added to runningSum that itself was divisible by k
    
    - example
    k = 5
    runningSum = 23 % 5 = 3
    runningSum = (23+3) % 5 = 26 % 5 = 1
    runningSum = (23+3+17) % 5 = 43 % 5 = 3 <--- remainder is again 3
    Which means after 23, we added some numbers which are divisible by K for the same remainder to show up again
    
    So the reason we have to look up for the same remainder is because some number was added to runningSum which was divisible by k
    No matter what numbers we add to runningSum, we can keep adding any multiples of k to runningSum, the remainders of all WILL BE THE SAME
    
    eg; 23 + 5 = 28 % 5 = 3
    eg; 23 + 10 = 33 % 5 = 3
    eg; 23 + 15 = 38 % 5 = 3
    eg; 23 + (5+5+5+5+5) = 48 % 5 = 3
    
    Turns out that adding any number that is divisible by the divisor will result into the same remainder

*/
func checkSubarraySum(nums []int, k int) bool {
    // either add that we have seen a remainder 0 at idx -1
    // OR handle in the loop when remainder becomes 0, we know from idx 0 till current ith position we have a good subarray
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