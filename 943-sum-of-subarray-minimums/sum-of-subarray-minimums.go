/*
    for each subarray we need a minimum
    instead of forming each subarray
    find out the number of subarrays if ith ele was a min ele
    "whats will be the contribution of ith element if we knew number of subarrays where ith ele is min"
    
    basically, evaluate each ith element ( like trapping rain water; evaluate how much we can compute on top of this building )
    to evaluate each ith element
    we need number of subarrays ith element is min
    then its total contribution is nums[i] * numOfSubarrays

    approach:
    - [0,3,1,2,4]
    - when i is at val 1
    - i stays minimum till idx 0 on the left hand side
        - this means, all subarrays ending at i(val 1) from left hand sides are
        - [1], [3,1], [0,3,1]
        - total 3 subarrys on left hand side of i
    - i stays minimum till last idx on the right hand side
        - this means, all subarrays starting with i(val 1) in right hand side are
        - [1], [1,2], [1,2,4]
    - this far, we have counted subarrays ending at i and subarrays starting from i
    - but i could be a middle element too, and those subarrays are
        - [0,3,1,2], [0,3,1,2,4], [3,1,2], [3,1,2,4]
    - now we need sum of min in these subarrays
    - well we already were evaluating for ith element(val 1) 
    - meaning if we have n subarrays where 1 is the min
    - what is the sum of all those n subarrays min
    - n subarrays, each has a min value of 1
    - total sum = n * minVal = n * 1
    - in this case our n is; 9
    - n = [0,3,1],[3,1],[1],[1,2],[1,2,4],[0,3,1,2],[0,3,1,2,4],[3,1,2],[3,1,2,4]
    - therefore total sum where 1 is the min element is 9*1 = 9
    - now repeat the rest for each ith element

    - [0,3,1,2,4]

    - to figure out the number of subarrays where i is the min element
    - this is NSR and NSL
    - NSL of val 1 will be at idx -1
    - how many elements/subarrays on left if NSL is at idx -1 ?
        - idx - NSL
        - 2-(-1) = 3
    - NSR of val 1 will be at idx 5
    - how many elements/subarrays on right if NSR is at idx 5 ?
        - NSR - idx
        - 5-2 = 3
    - total subarrays = 3*3 = 9
    
    - For NSR, NSL ; we can use our basic "process top NSR/NSL via stack" template
    
    time = o(n)
    space = o(n)
*/
func sumSubarrayMins(nums []int) int {
    mod := 1000000007
    st := []int{}
    sum := 0
    for i := 0; i < len(nums); i++ {
        curr := nums[i]
        for len(st) != 0 && curr < nums[st[len(st)-1]] {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            nsr := i
            nsl := -1
            if len(st) != 0 {nsl = st[len(st)-1]}
            leftCount := top-nsl // number of starting positions
            rightCount := nsr-top // number of ending positions
            totalCount := leftCount*rightCount // total number of subarrays
            // if nums[top] = 2 ; and total subarrays are 12
            // we want to find the sum of this value(2) in subarrays where this value(2) is min
            // value is 2, it shows up in 12 subarrays, therefore its sum is 12*2 = 24
            currSum := nums[top] * totalCount
            sum = (sum + currSum) % mod
        }
        st = append(st, i)
    }
    // monotonic stack edge cases;
    // do we need to do something with elements that are still in stack?
    // yes, because we were processing nsr, and if there are elements in stack
    // for example; [1,2,3,4]
    // it means we did not find a ith element that was a nsr of a previous element
    // therefore these elements range/window can be the whole array
    // therefore nsr for each of the remaining element is n (size of array)

    for len(st) != 0 {
        top := st[len(st)-1]
        st = st[:len(st)-1]        
        nsr := len(nums)
        nsl := -1
        if len(st) != 0 {nsl = st[len(st)-1]}
        leftCount := top-nsl // number of starting positions
        rightCount := nsr-top // number of ending positions
        totalCount := leftCount*rightCount // total number of subarrays
        // if nums[top] = 2 ; and total subarrays are 12
        // we want to find the sum of this value(2) in subarrays where this value(2) is min
        // value is 2, it shows up in 12 subarrays, therefore its sum is 12*2 = 24
        currSum := nums[top] * totalCount
        sum = (sum + currSum) % mod

    }
    return sum

}


/*
    pre-compute nsr and nsl using helpers

    edge case: duplicate subarray counting
     0  1  2  3
    [71,55,82,55]



*/

// func sumSubarrayMins(arr []int) int {
//     nsr := nsr(arr)
//     nsl := nsl(arr)
//     mod := 1000000007
//     count := 0
//     for i := 0; i < len(arr); i++ {
//         rw := nsr[i]-i
//         lw := i-nsl[i]
//         count = (count + (arr[i]*(lw*rw))) % mod
//     }
//     return count
// }


// func nsr(nums []int) []int{
//     n := len(nums)
//     out := make([]int, n)
//     st := []int{}
//     for i := n-1; i >= 0; i-- {
//         out[i] = n
//         curr := nums[i]
//         for len(st) != 0 && nums[st[len(st)-1]] >= curr {
//             st = st[:len(st)-1]
//         }
//         if len(st) != 0 {
//             out[i] = st[len(st)-1]
//         }
//         st = append(st, i)
//     }
//     return out
// }


// func nsl(nums []int) []int{
//     n := len(nums)
//     out := make([]int, n)
//     st := []int{}
//     for i := 0; i < n; i++ {
//         out[i] = -1
//         curr := nums[i]
//         for len(st) != 0 && nums[st[len(st)-1]] > curr {
//             st = st[:len(st)-1]
//         }
//         if len(st) != 0 {
//             out[i] = st[len(st)-1]
//         }
//         st = append(st, i)
//     }
//     return out
// }




/*
    brute force
    approach: create all possible subarrays
    - create all subarrays
    - keep track of min in that subarray
    - keep adding min to the total sum

    time = o(n^2)
    space = o(1)
*/
// func sumSubarrayMins(arr []int) int {
//     mod := 1000000007
//     total := 0
//     for i := 0; i < len(arr); i++ {
//         minSoFar := arr[i]
//         for j := i; j < len(arr); j++ {
//             minSoFar = min(minSoFar, arr[j])
//             total = (total+minSoFar) % mod
//         }
//     }
//     return total
// }