/*
    intuition:
    - The idea is to use 2 pointers ( slow and fast )
    - slow starts at idx 0 and fast starts at 1
    - the idea is that the slow pointer acts as the left wall
    - If the fast pointer wall is less than slow pointer wall, we can trap water ( slow-fast height diff )
    - If the wall on the right side/fast-ptr becomes >= slow ptr wall, we can no longer use slow ptr as a left wall, therefore we must save what we have 
    - and move the slow ptr to the new left wall, i.e where fast ptr is
    - But what if the walls are in descending order, i.e they are always going down from left to right
    - then slow ptr will always be greater than fast ptr height, so we will trap but will never move the slow ptr to a new position
    - therefore we will never really save the accumulated trapped
    - Could we have a test case where this would fail?
    - Yes, when there is a 1 peek that is largest and on the right side of the peek, we next heights may never be >= the peek height
    - But could have little waters after the peek that can be trapped but never saved
    - In this case we will run a loop from the back ( 2nd pass )
    - Until the peek ( which will be the slow ptr )
    - And new slow ptr will become the right wall
    - And fast ptr will start from 2nd last position.
    

    approach: 2 pass using 2 pointers ( slow and fast ptrs )
    - We want to position our 2 pointers such that we know we can trap some water
    - fast pointer will always move
    - if height at fast pointer is smaller than height at left
    - this means we can trap some water between left and right - so we will take whatever we can trap here
        - trap += height[left]-height[fast]
    - when height at fast is not less than height at slow, this means we are done trapping as much water as we can , so add our trapped value to final result
    - move slow to fast and reset trap value
    - always move fast pointer regardless of which path we take
    - but this does not work in a single pass everytime... because we could have a slow pointer pointing the tallest bar in the middle of the array
    - and all bars to right are always less than slow height ( i.e height[fast] is always less than height[slow] because slow is at the tallest point )
    - Therefore another pass from the back is needed to check if we can trap water the other until this peek ( if there is a peek value slow paused at )
    - so loop back until peek idx and repeat the same checks
   
   time: o(2n) - but 2 is constant
   space: o(1)
   
*/ 
func trap(height []int) int {
    n := len(height)
    result := 0
    trap := 0
    
    slow := 0 // acting as the left wall
    fast := slow+1
    
    for fast < n {
        if height[slow] > height[fast] {
            trap += height[slow]-height[fast]
        } else {
            result += trap
            trap = 0
            slow = fast
        }
        fast++
    }
    
    peek := slow
    slow = n-1 // acting as the right wall now
    fast = slow - 1
    trap = 0
    for fast >= peek {
        if height[slow] > height[fast] {
            trap += height[slow]-height[fast]
        } else {
            result += trap
            trap = 0
            slow = fast
        }
        fast--
    }
    
    return result
}