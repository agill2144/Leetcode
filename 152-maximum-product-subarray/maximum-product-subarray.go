/*
    approach: brute force
    - form all subarrays
    time = o(n^2)
    space = o(1)
    
    
    approach: prefix / suffix product
    - There are 4 main things to observe
    1. If all numbers are positive , we want to use all numbers
    2. If there are even number of negatives, we want to use all numbers
    3. If there are odd number of negatives,
        - form subarrays with even negatives
        - because even negatives multiplied will result into positive number
        - [x -x x -x x x -x]
           --------- is 1 possibility ( even number of negatives )
                  --------- is another possibility ( even number of negatives )
        
        - How do we create these subarrays efficiently ?
        - runningProduct ( prefixRunningProduct and suffixRunningProduct )
            - we need answers from both possibilities

    4. There are 0s in the middle
        - any number multiplied by a 0 will result into a 0
        - therefore as soon as we run into a 0, we know we cant use this element
        - therefore stop this subarray and start forming a new one

    
    - maintain a prefix , suffix product
    - loop thru numbers
    - simultaneously update both prefix and suffix prods
        prefix *= curr ith number
        suffix *= curr ith number from back of the array
    - save the max out of the two
    - if we prefix or suffix prods have turned to 0, it means its times to drop the subarrays
    - and start forming new ones
    
    time = o(n)
    space = o(1)
    
*/
func maxProduct(nums []int) int {
    n := len(nums)
    leftRp := 1
    rightRp := 1
    res := math.MinInt64
    for i := 0; i < n; i++ {
        leftRp *= nums[i]
        rightRp *= nums[n-1-i]
        res = max(res, max(leftRp, rightRp))
        if leftRp == 0 {leftRp = 1}
        if rightRp == 0 {rightRp = 1}
    }
    return res
}
