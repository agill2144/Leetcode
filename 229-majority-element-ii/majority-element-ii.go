/*
    same problem as majority majority-algorith-1
    here we have to find numbers repeating more than n/3 times
    there can be multiple numbers and we have to return all of them

    approach: moores voting algorithm
    - VERY IMPORTANT INSIGHT
    - if there are n total elements
    - and we have find elements that repeat > n/3 times
    - that means at max we can only have 2 elements, maybe less, but never more than 2
    
    eg; n = 10, n/3 = 3
    - means that a element must repeat more than 3 times to be considered a majority element
    - THIS ALSO MEANS
    - at min there must be atleast 4 counts of the same element
    - how many such elements can we take if at min this element must be repeat atleast 4 times 
    - thats 2 elements at max
        - 4 + 4 = 8
        - there can x repeating 4 times, y repeating 4 times, but never another z repeating 4 times because we only have 10 elements in total
    
    eg; n = 21, n/3 = 7
    - meaning, at min, there must be 8 repeats ( 7+1 ) for an element to be considered majority element
    - how many times we can take a element that is repeating at min 8 times before we blow past array size ?
        - 8+8 = 16
        - we can have x repeating 8 times, y repeating 8 times, but never another z repeating 8 more times, that would exceed our input array
    
    - Therefore this problem boils down to tracking 2 elements ( maybe less ) that are repeating more than n/3 times
    - therefore we can use the same algo as majority-element-1
        - instead of 1 count and 1 num
        - we will have 2 counts and 2 num vars
        
    Moores voting algorith
    1. Find a candidate
    2. Verify the candidate
    
    In this case, we have 2 candidates
    1. Find 2 candidates
    2. Verify both candidates candidate
    3. Save a candidate if they are repeating > n/3 times    
*/
func majorityElement(nums []int) []int {
    ele1 := math.MinInt64
    count1 := 0
    ele2 := math.MinInt64
    count2 := 0
    
    for i := 0; i < len(nums); i++ {
        if count1 == 0 && nums[i] != ele2 {
            count1 = 1
            ele1 = nums[i]
        } else if count2 == 0 && nums[i] != ele1 {
            count2 = 1
            ele2 = nums[i]
        } else if nums[i] == ele1 {
            count1++
        } else if nums[i] == ele2 {
            count2++
        } else {
            count1--
            count2--
        }
    }
    
    count1 = 0
    count2 = 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == ele1 {count1++}
        if nums[i] == ele2 {count2++}
    }
    out := []int{}
    if count1 >= (len(nums)/3)+1 {out = append(out, ele1)} 
    if count2 >= (len(nums)/3)+1 {out = append(out, ele2)} 
    return out
}