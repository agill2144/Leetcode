/*
    we have to find the next smallest number using the same digits in nums array
    12 -> 21
    23 -> 32
    112 -> 121
    1456 -> 1465

    we have to create the next number but the smallest one possible ( using the same digits )
    [1,2,6,5,3,2,1]
    - next number of any number usually means incrementing right handside digit
    - but we cant increment, we have to use the same digits and rearrange in such way that it becomes the next number
    
    - from the right hand side of the array,
    - we can look at each adjacent number
    - [1] -> no number on the right side to compare with, skip
    - [2,1] -> this is already in the highest possible permutation
    - [3,2,1] -> 3,2 is already in the highest possible permutation
    - [5,3,2,1] -> 5,3 is already in highest possible permutation
    - [6,5,3,2,1] -> 6,5 is already in highest possible permutation
    - [2,6,5,3,2,1] -> 2,6 is NOT in the highest possible permutation
    
    - so we have found a position where the number is smaller than its adjacent right number
    - this is the dipping point, i.e breachIdx
    - why are we looking for this dipping point ?
        - because we want to replace this dipping number with the next larger number
        - inorder to create permutation
    - what do we replace this breachIdx number with ?
        - the smallest possible number on the right side ?
        - no, not the smallest, using the smallest will make us go "prev permutation"
        - we want a number that just greater than breachIdx but the smallest num possible
        - therefore from n-1 -> breachIdx, find the next smallest number
    - [1,2,6,5,3,2,1]
         ^     ^     
         b     s
    - b = breachIdx; s = nextSmallest
    - swap the 2 nums
    - [1,3,6,5,2,2,1]
         ^     ^     
         b     s
    - now we have the next greater prefix [1,3 .... ]
    - after the prefix, we want the smallest possible set of numbers
        1. sort everything after breachIdx 
        2. reverse n-1 to breachIdx+1
    - [1,3,1,2,3,5,6]
         ^     ^     
         b     s

    - if we dont find a breachIdx, that means, this entire array is already in the greatest permutation possible
    - reverse the entire array to reset back to smallest possible permutation

    TLDR
    - find the first adacent pair from right side where left is smaller than right
    - if you dont find such a pair, reverse the whole array
    - swap the left element in the pair with the smallest element greater than that to its right
    - then reverse the sub array from the point of swap(after the left element in the pair) till the end


    Replace the smallest number ( from right side ), find the smallest from right - this is the breachIdx
    with the next biggest number possible - this is the next smallest after breachIdx value
    than reverse the array after breachIdx+1 : end
    
    time = o(n)
    space = o(1)
*/

func nextPermutation(nums []int)  {
    breachIdx := -1
    n := len(nums)
    
    // find the first adacent pair from right side where left is smaller than right
    // why smaller than right ?
    // because we want to relace the smallest number from right side with the next biggest number
    for i := n-2; i >= 0; i-- {
        if nums[i] < nums[i+1] {
            breachIdx = i
            break
        }
    }
    
    // look for the next biggest number ( just right of breachIdx value )
    if breachIdx != -1 {
        ns := -1
        for i := n-1; i > breachIdx; i-- {
            if nums[i] > nums[breachIdx] {
                if ns == -1 || nums[i] < nums[ns] {
                    ns = i
                }
            }
        }
        // swap the smallest number from right side ( breachIdx) with the next possible number (ns)
        nums[breachIdx], nums[ns] = nums[ns], nums[breachIdx]
    }
    
    // reverse all elements to the right of breachIdx, so that they are all in increasing order from breachIdx
    // why? because we put the next greater number at breachIdx, and now we want the smallest possible permutation to the right of breachIdx
    left := breachIdx+1
    right := n-1
    for left < right {
        nums[left], nums[right] = nums[right], nums[left]
        left++
        right--
    }
}