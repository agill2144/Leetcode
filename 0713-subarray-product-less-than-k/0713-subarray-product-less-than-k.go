/*
        sliding window  OR  Running Sum ?
        Check if the numbers have negative
        if they dont , sliding window could work. 
     
        In this case we wont have negatives
        1 <= nums[i] <= 1000
        
        What type of sliding window?
        - Dynamic
        
        What defines keep increasing the window size?
        - If the window/running prod < k
        
        When do we start shriking?
        - if the window prod >= k ( while windowProd >= k, slide the window forward )
        
        When we slide the window forward, an element leaves our subarray, what state needs to be updated?
        - Since we maintain running prod, the num going out of the window will need to be divided from running prod
        
        What are we looking to return once we do have a window of running prod < k ?
        - count number of subarrays 
*/
func numSubarrayProductLessThanK(nums []int, k int) int {
    if nums == nil || len(nums) == 0 {return 0}
    /*
        k = 1; nums = [1,2,3,4,1]
        in this case prod will always be atleast 1
        but we want strictly less than k ( not including k)
        
        k = 1; nums = [1,1,1,1]
        same problem as above
        - our running prod will always be atleast 1 and never less than 1
        
        k = 1; nums = [0,1]
        - cannot have 0's because of the constraints
        - nums[i] >= 1 && <= 1000
    */
    if k <= 1 {return 0}
    
    left := 0
    count := 0
    prod := 1
    for i := 0; i < len(nums); i++ {
        prod *= nums[i]
        for prod >= k {
            prod /= nums[left]
            left++
        }
        count += i-left+1
    }
    return count
}