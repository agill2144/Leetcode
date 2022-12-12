/*
    approach: slow and fast pointers
    - Slight variation of remove duplicates from sorted array
        - where we used slow pointer as a position where the next uniq element should go.
    - The only difference in this one is that we have to remove duplicates but allow 2 dupes at max.
    - We will use the same technique but along with slow and fast pointer we will have a counter
    - This counter will be keeping track of number of times fast pointer saw this element
    - counter goes up by 1 if value at fast == value at fast-1
    - counter resets to 1 if value at fast != value at fast-1
    - As long as count <= k (2), we will copy the value at fast to slow and move slow by 1
    - Slow will hold on at a position if count > k -- i.e slow has seen this character k times so far and waiting for next element to be placed
    - As soon as fast pointer resets count to 1 because fast value is not same fast-1 value
    
    time: o(n)
    space: o(1)
    
    the other brute force to create a separate list and write it back to nums list
*/
func removeDuplicates(nums []int) int {
    if nums == nil || len(nums) == 0 {
        return 0
    }
    
    k := 2
    slow := 1
    fast := 1
    count := 1
    n := len(nums)
    
    for fast < n {
        if nums[fast] == nums[fast-1] {
            count++
        } else {
            count = 1
        }
        if count <= k {
            nums[slow] = nums[fast]
            slow++
        }
        fast++
    }
    return slow
}