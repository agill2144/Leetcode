/*
    binary search on answer hints:
    - brute force is looping over a range of sorted numbers
    - answer lies within a known range
    - if we pick an answer to evaluate we get some x result
        - if we increase the answer value and evaluate again, x is either increasing or decreasing
        - if we decrease the answer value and evaluate again, x is either increasing or decreasing
        - in this case;
        - when we pick a smaller divisor; sum is a large number
        - when we start to increase the divisor; sum starts going down
        - when we monotonically increase the answer-to-eval; our resultant-sum starts to monotonically decrease

    approach: binary search on answers
    - we can clearly see that the outter loop is running a naive range loop
    - from 1 to maxVal
    - we can cut this time in half by using binary search
    - our answer lies in a range ( 1 to maxVal )
    - ranges are sorted!
    - therefore binary search can be applied
    - the inner for loop remains the same
    - mid will be evaluated the same way
    - if our sum is <= threshold
        - then we know for sure that increasing mid will definitely work
        - but we want the smallest possible mid
        - therefore save this mid as potential ans and look left
        - i.e right = mid-1
    - otherwise our sum was > threshold
        - meaning our divisor ( mid ) is still too small
        - pick a bigger divisor
        - i.e left = mid+1
    
    time = o(n) + o(logMaxVal * n)
    space = o(1)
*/
func smallestDivisor(nums []int, threshold int) int {
    maxVal := math.MinInt64
    for i := 0; i < len(nums); i++ {
        if nums[i] > maxVal {
            maxVal = nums[i]
        }
    }
    ans := -1
    left := 1
    right := maxVal
    for left <= right {
        mid := left + (right-left)/2
        sum := 0
        for i := 0; i < len(nums); i++ {
            q := int(math.Ceil(float64(nums[i])/float64(mid)))
            sum += q
        }
        if sum <= threshold {
            ans = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
}
/*
    approach: brute force
    - because we want the smallest such divisor, lets be greedy and start with the smallest such divisor ( i.e 1 )
    - and we have to go until maxVal in nums array.
    - why maxVal from nums array ?
        - [1,2,5,9]
        - dividing each number with 9 will result in the smallest quotient ; [1,1,1,1]
        - therefore resulting in the smallest sum ; 1+1+1+1 = 4
        - therefore leading to an answer for SURE!
        - therefore we do not need to go beyond the maxVal 
    - for each divisor, we will simulate what the question is asking
        - divide each element
        - add it to a sum
    - after evaluating a divisor, we have a sum, check if its <= threshold
    - if it is, return that divisor and exit
    - we started from the smallest and therefore as soon as we run into an answer, we can call it done
    
    time; 
    o(n) to look for maxVal
    + o(maxVal * n)
    
    space;
    o(1)
*/
// func smallestDivisor(nums []int, threshold int) int {
//     maxVal := math.MinInt64
//     for i := 0; i < len(nums); i++ {
//         if nums[i] > maxVal {
//             maxVal = nums[i]
//         }
//     }
//     for i := 1; i <= maxVal; i++ {
//         sum := 0
//         for j := 0; j < len(nums); j++ {
//             q := int(math.Ceil(float64(nums[j])/float64(i)))
//             sum += q
//         }
//         if sum <= threshold {
//             return i
//         }
//     }
//     return -1
// }