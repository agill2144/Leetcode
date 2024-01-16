/*
    "subarrays such that the largest sum of any subarray is minimized"
    - binary search on answers hint
    - max ans is minimized
    
    approach: binary search on ans
    - we know our ans lies in a range ( from brute force )
    - mid becomes our $atMax sum allowed
    - evaluate mid;
        - how many subarrays can we create if each allowed subarray sum cannot exceed $mid
    - if we get back more than k subarrays
        - it means, our allowed sum per subarray was too small
        - therefore increase the allowed sum per subarray
        - in-hopes to reduce the number of subarrays needed
        - i.e left = mid+1
    - if we get back <= k subarrays
        - why less than ?
        - because when we were evaluating, we were greedily adding a element to subarray sum (as long as it worked)
        - i.e if rSum <= $mid { we are fine, continue to next element }
        - and like this we may only have had 1 subarray when we needed 3
        - but this means if greedily we were able to fit all elements into 1 subarray without execeeding its sum 
        - it means we can easily split that 1 giant subarr into 3 and still none of them will exceed $max sum
    - it means it worked, save this as potential ans
    - However we want smallest such ans, search left
    
    time = o(log$totalSum * n)
*/
func splitArray(nums []int, k int) int {
    if k > len(nums) {return 0}
    left := 1
    right := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] > left {left = nums[i]}
        right += nums[i]
    }
    ans := 0
    for left <= right {
        mid := left + (right-left)/2
        
        count := 1
        rSum := 0
        for i := 0; i < len(nums); i++ {
            rSum += nums[i]
            if rSum > mid {
                count++
                rSum = nums[i]
            }
        }
        
        if count > k {
            // we needed more than k subbarrays at our given allowed sum per subarray
            // therefore increase sum-allowed-per-subarray to decrease num of subarrays
            left = mid+1
        } else {
            ans = mid
            right = mid-1
        }
    }
    return ans
}

/*
    "subarrays such that the largest sum of any subarray is minimized"
    - binary search on answers hint
    
    approach: brute force
    - try all possible sums starting from 1 to maxPossibleSum ( sum of all numbers )
    - that is, see how many subarrays can we create we a subarray atMax cannot exceed sum=1

    - if number of subarrays is more than k ( when we needed k subarrays )
    - it means our sum allowed per subarray is too small, increase it
    
    - if number of subarray is <= k 
    - it means our sum allowed per subarray worked
    - exit early since we are starting from smallest such sum
    
    time = o(maxPossibleSum * n)
    space = o(1)

*/
// func splitArray(nums []int, k int) int {
//     start := 1
//     end := 0
//     for i := 0; i < len(nums); i++ {end += nums[i]}
    
//     for i := start ; i <= end; i++ {
//         atMax := i
        
//         rSum := 0
//         count := 1
//         for j := 0; j < len(nums); j++ {
//             rSum += nums[j]
//             if rSum > atMax {
//                 count++
//                 rSum = nums[j]
//             }
//         }
        
//         if count > k {
//             continue
//         }
        
//         return i
//     }
//     return -1
// }