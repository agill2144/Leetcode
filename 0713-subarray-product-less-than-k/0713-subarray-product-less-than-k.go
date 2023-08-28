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
    if nums == nil || len(nums) == 0 || k <= 1 {
        return 0
    }

    prod := 1
    left := 0
    count := 0
    for right := 0; right < len(nums); right++ {
        prod *= nums[right]
        for prod >= k {
            prod /= nums[left]
            left++
        }
        count += right-left+1
    }
    return count
}