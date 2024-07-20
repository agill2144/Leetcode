/*
    sliding window:
    - the question itself tells us to create a window
    - a window of fixed size; k
    - and keep track of max in each window
    - the problem with tracking max in our window
        - say we have a window of size k right now
        - we have left ptr, i ptr and we have our max val in this window
        - and lets assume the max value is the left ptr right now
        - when the window slides forward ( since our size is met and we cannot expand our window )
        - left ptr is no longer part of our new window, therefore the max val is now lost
        - we could run another loop after sliding the window from left; from left -> i ptr
            - to find our next max
        - this is again coming down to brute force time complexity
        - therefore not optimal
        - therefore we need some data structure that can help us keep track of max val in our window
    - data structures that can help keep track of max val:
        - max heap
        - this can reduce our time complexity to o(n log n)
    - when we slide our window we simply need our next max value 
    - we can leverage a monotonica increasing/decreasing datastructure to keep track of max value(s) in our window
    - since we want next max value to take place, it makes sense if we keep value in mono-decreasing order
        - meaning highest value at the front, rest are all decreasing
    - mono increasing / decreasing data structures:
        - stack
        - dq
    - lets evaluate stack:
        - if our max is at the top of our stack, meaning top is higest, and elements underneath will be smaller
            - from stacks perspective; this is mono decreasing stack (starting from top of stack)
        - so we want to store elements in increasing order,
        - if curr > top of stack; push it to stack as is
        - if curr < top of stack, pop until this condition is true ( since smaller values should be at the bottom of highest val )
            - push curr val
            - re-push all elements that we have popped
            - this is inefficient
        
        - lets try storing them in mono descreasing manner;
            - ie decreasing from top of stack
            - meaning the max value is at the bottom of stack
            - if curr > top of stack; pop, because we have a better max for a upcoming window
                - keep doing this until curr > top
            - now our max is at the bottom, next max above the max and so on .... [10,9,8,7,6] <- top of stack
            - we can easily answer; "whats the max in our window"
            - that is st[0]
            - when left ptr moves forward and st[0] is no longer part of our window; we need to remove st[0]
            - therefore we need some better data structure that can
                1. push to top in o(1) time
                2. pop from top in o(1) time
                3. push to front in o(1) time (we do not need this in this question)
                4. pop from front in o(1) time
            - this is why a double-ended-queue is better!
    - double ended queue:
        - store elements in mono decreasing
        - i.e max value is a value that is part of our window
        - and this max value is dq[0]
        - and elements after that are smaller in value!
        - from top (back) of dq; its mono increasing
        - from front of dq; its mono decreasing
        - if curr element is > top(back) of dq;
            - pop until the above true
            - it means; we have a better value for the next set of windows
        - if our window size is met ( i-left+1 == k )
            - max is at front of dq; dq[0]
            - we need to slide our window; shrink from left first
                - if left ptr is == dq[0]
                - rm dq[0] 
            - left++ 
    time = o(n)
    space = o(n)
*/

func maxSlidingWindow(nums []int, k int) []int {
    n := len(nums)
    dq := []int{} // idx
    out := []int{}
    left := 0 // idx
    for i := 0; i < n; i++ {
        for len(dq) != 0 && nums[i] > nums[dq[len(dq)-1]] {
            dq = dq[:len(dq)-1]
        }
        dq = append(dq, i)
        if i - left + 1 == k {
            out = append(out, nums[dq[0]])
            if dq[0] == left {dq = dq[1:]}
            left++
        }
    }
    return out
}

/*
    brute force:
    - form all possible subarrays of size k
    - keep track of max val
    - save max each time we have a subarray of size k

    time = o(n * k)
    - however k could be == n
    - therefore time = o(n*n) = o(n^2) in the worst case

    space = o(1)
*/
// func maxSlidingWindow(nums []int, k int) []int {
//     out := []int{}
//     n := len(nums)
//     for i := 0; i < n-k+1; i++ {
//         maxVal := math.MinInt64
//         for j := i; j < n; j++ {
//             maxVal = max(maxVal, nums[j])
//             if j-i+1 == k {break}
//         }
//         out = append(out, maxVal)
//     }
//     return out
// }