/*

    1. From the back of the nums array
    2. Find the breach 
        - What does a breach mean?
        - breach = nums[i] < nums[i+1]
        - example: 3,1,4,8,7,6,2,1
        - breach is 4 because 4 < 8
        
    3. Find the next greatest number ( but smallest among all - i.e the next avail number ) of breach on right side of breach
        - example: 3,1,4,8,7,6,2,1
        - breach is 4 because 4 < 8
        - next number avail after breach is 6
    4. If found, Swap with the number found un step3
        - example: 3,1,4,8,7,6,2,1
        - breach is 4 because 4 < 8
        - next number avail after breach is 6
        - after swapping - 3,1,6,8,7,4,2,1
        - there is a chance that we have 
        - 321 - then there is no breach, so there is nothing to swap ( i will reach -1 )
    5. Finally reverse all numbers from breach + 1 till end 
        - Why?
        - after swapping - 3,1,6,8,7,4,2,1
        - But the next permutation is 3,1,6,1,2,4,7,8
        - So all numbers after breach must be reversed
    
    time: o(n)
    space: o(1)


*/

func nextPermutation(nums []int)  {    
    
    breachIdx := len(nums)-1
    for breachIdx >= 0 {
        if breachIdx == len(nums)-1 {breachIdx--; continue}
        if nums[breachIdx] < nums[breachIdx+1] {
            break
        }
        breachIdx--
    }
    
    // we have a breach, swap it with its next avail number
    i := len(nums)-1 
    nextAvailIdx := -1
    if breachIdx != -1 {
        // if we found a breach ( i.e breach < numsToRight )
        // there is 100% a number greater than breach towards the right
        for i > breachIdx {
            if nums[i] > nums[breachIdx] {
                if nextAvailIdx == -1 {
                    nextAvailIdx = i
                } else {
                    if nums[i] < nums[nextAvailIdx] {
                        nextAvailIdx = i
                    }
                }
            }
            i--
        }
        nums[breachIdx], nums[nextAvailIdx] = nums[nextAvailIdx], nums[breachIdx] 
    }
    
    // reverse after breach ( if there was one, otherwise reverse the whole thing )
    upToIdx := breachIdx+1
   
    nums = reverse(upToIdx, len(nums)-1, nums)
}


func reverse(left, right int, nums []int) []int {
    for left < right {
        nums[left],nums[right] = nums[right],nums[left]
        left++
        right--
    }
    return nums
}