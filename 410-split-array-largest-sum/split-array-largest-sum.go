/*
    approach: binary search
    - we see that we have a range loop
    - ranges are sorted, binary search hint
    - another binary search hint: "largest sum of any subarray is minimized"
        - this is "max is minimized" or "min is maximimzed" pattern type of question
    - therefore for sure binary search
    - "largest sum of any subarray is minimized" i.e max is minimized
    - this means, mid will be our guess (i.e atMax a subarray sum can be $mid )
    - then will check how many subarrays we can make with mid being atMax subarray sum
    - if we end up creating more than k subarrays, it means our atMax was too small
        - increase it , left = mid+1
    - otherwise, this mid value is a potential ans
        - save it, and continue searching on left side of the range
        - since we want smallest such ans
        - i.e right = mid-1

    time = o(n) + o(log$TotalSum-maxValue * n)
    space = o(1)

*/
func splitArray(nums []int, k int) int {
    left := math.MinInt64
    right := 0
    for i := 0; i < len(nums); i++ {
        left = max(left, nums[i])
        right += nums[i]
    }
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        count := 1
        rSum := 0
        for i := 0; i < len(nums); i++ {
            rSum += nums[i]
            if rSum > mid {
                rSum = nums[i]
                count++
            }
        }
        if count <= k {
            ans = mid
            right = mid-1
        } else {
            left= mid+1
        }
    }
    return ans
}

/*
    abs brute force
    - try forming every single subarray sum
    - using recursion, while keeping track of each subarray sum
    - and while keeping track of maxSubarry sum so far
    - and as soon as we have exactly k subarrays
    - compare with a global ans and save the min ; min(global, maxSubarraySumInThisConfiguration)

    time = exponential
    space = o(n)


    approach: range sum
    - one of the things to to always check with array like problems is;
    - does the answer lie within a range that we can assume
    - if it does, and the array elements can be used as subarray and not subseq, there is likely a n^2 solution
    - using a range loop and checking whether some constraint is being met or not
    
    - in this case, our ans does lie within a known range
    - max possible sum is total sum of all elements
    - min possible sum is 0
    - Now we want the smallest such sum, therefore lets be greedy and start with 0
    - if atMax, a subarray sum can be 0 and cannot exceed 0, how many subarrays can we form
        - this is not possible if atMax allowed is 0
        - therefore lets try 1...
        - also not possible
        - turns out the smallest such answer is the largest number in the arr
    - therefore the start of the range is max element in array and end of the range is total sum
    - now we need to check if $atMax, a subarray sum cannot exceed this value, how many subarrays can we form
    - how do we check that ?
        - we need to keep track of a running sum
        - how many subarrays we have formed so far
            - start this with 1, because we are starting to form the first subarray
        - and start a loop on the arr
        - add each element to running sum
        - if running sum exceeds atMax ( i.e runningSum > atMax )
        - we need to split the array
            - count++
            - reset runningSum ( but include current element since it was the one that caused it to go past our atMax value )
            - i.e runningSum = nums[j]
    - now we will have a count of subarrays created if atMax a subarray sum can be $x value ( range loop variable )
    - if we ended up created more subarrays, it means, our atMax was so small that we needed more than k subarrays
    - therefore increase the $atMax per subarray in hopes to reduce the number of subarrays formed.
    - As soon as we have exactly k subarrays, we are looking at the smallest subarray sum ( which was the largest amongst its split )
    - therefore return that $x value

    time = o(n) + o(max-totalSum+1 * n)
    space = o(1)
*/
// func splitArray(nums []int, k int) int {
//     start := math.MinInt64
//     end := 0
//     for i := 0; i < len(nums); i++ {
//         start = max(start, nums[i])
//         end += nums[i]
//     }

//     for i := start ; i <= end; i++ {
//         atMax := i
//         count := 1
//         rSum := 0
//         for j := 0; j < len(nums); j++ {
//             rSum += nums[j]
//             if rSum > atMax {
//                 rSum = nums[j]
//                 count++
//             }
//         }
//         if count > k {
//             continue
//         }
//         return i
//     }
//     return -1
// }