/*
    approach: nsr + nsl ( process top ) pattern
    - instead of creating each possible subarray
    - we will focus on evaluating each elements contribution to our total final sum
    - each element gets added to our total sum based on how many subarrays it stays as minimum value in
    - example; [2,6,7,3,8,9,1]
    - lets say we are evaluating 3's contribution
    - 3 stays minimum in this window of the array [6,7,3,8,9]
    - how many subarrays are there where 3 is included? 
        - from left side; [6,7,3] [7,3], [3] 
        - from right side; [3,8], [3,8,9]
        - across; [6,7,3,8], [6,7,3,8,9], [7,3,8], [7,3,8,9]
        - total; 9 subarrays
    - if we figure out that there are 9 subarrays where 3 is the min in ALL 9 subarrays
    - what is 3's contribution ? 3*9 = 27 ( we need to add 3, 9 times; i.e 3*9)

    - but how can we form/count all of the above possible subarrays which has 3 ?
        - to count number of subarrays between left and i which includes ith value
            - i-left+1 ; i.e window size
            - i = 3; left = 1; leftCount = 3-1+1 = 3 subarrays on left side
        - to count number of subarrays between right and i which includes ith value
            - right-i+1; i.e window size
            - i = 3; right = 5; rightCount = 5-3+1 = 3 subarrays on right side
        - to count acorss; thats just leftCount*rightCount
        - therefore total count = leftCount*rightCount

    - now how do we find these left and right ptr for each ith value?
        - 3 was the min value in a window as long as no smaller value came in
        - [2,6,7,3,8,9,1]
        - on left; we stopped at idx 0 because its smaller than 3
        - THIS IS NSL ; we are looking for NSL for each ith element
        - same on the right side; 
        - on right; we stopped at last idx, because its smaller than 3
        - THIS IS NSR; we are looking for NSR for each ith element    
    - NSR and NSL can be implemented in 1 pass using "process top" 
        - am i(ith val) your(top-of-stack) nsr ? ; nsl is next top of stack
    - edge cases
        1. stack may have elements left after our intial pass
            - do we need to do something with elements that are still in stack?
            - yes, because we were processing nsr, and if there are elements in stack
            - for example; [1,2,3,4]
            - it means we did not find a ith element that was a nsr of a previous element
            - therefore these elements range/window can be the whole array
            - therefore nsr for each of the remaining element is n (size of array)
    time = o(n)
    space = o(n)
*/
func sumSubarrayMins(arr []int) int {
    mod := 1000000007
    n := len(arr)
    st := []int{} // indices; processing top
    ans := 0
    for i := 0; i < n; i++ {
        for len(st) != 0 && arr[i] < arr[st[len(st)-1]] {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            nsr := i; nsl := -1
            if len(st) != 0 {nsl = st[len(st)-1]}
            leftCount := top-nsl
            rightCount := nsr-top
            totalCount := leftCount * rightCount
            // arr[top] is min in $totalCount number of subarrays
            // and we have to sum of all minimums; i.e totalCount * arr[top]
            sum := arr[top] * totalCount
            ans = (ans+sum) % mod
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
        nsr := n; nsl := -1
        if len(st) != 0 {nsl = st[len(st)-1]}
        leftCount := top-nsl
        rightCount := nsr-top
        totalCount := leftCount*rightCount
        sum := arr[top] *totalCount
        ans = (ans+sum) % mod
    }
    return ans
}

/*
    approach: brute force
    - used nested for loop
    - form all possible subarrays
    - keep track of curr min in each subarray
    - add curr min to total sum each time
    time = o(n^2)
    space = o(1)
*/
// func sumSubarrayMins(arr []int) int {
//     total := 0
//     n := len(arr)
//     mod := 1000000007
//     for i := 0; i < n; i++ {
//         currMin := arr[i]
//         for j := i; j < n; j++ {
//             currMin = min(currMin, arr[j])
//             total = (total + currMin) % mod
//         }
//     }
//     return total
// }