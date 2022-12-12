func numSubarrayProductLessThanK(nums []int, k int) int {
    
    /*
        Subarray == sliding window  ||  prefix/running sum/prod
        Check if the number have negative
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
    
    if nums == nil || len(nums) == 0 || k <= 1 {
        return 0
    }
    count := 0
    windowProd := 1
    left := 0
    for right := 0; right < len(nums); right++ {
        windowProd *= nums[right]
        for windowProd >= k {
            // shrink
            leftNum := nums[left]
            windowProd /= leftNum
            left++
        }
        
        // when we get here our window prod < k
        // For every right, we update left and prod to maintain this invariant. Then, the number of intervals with subarray product less than k and with right-most coordinate right, is right - left + 1. We'll count all of these for each value of right.
        count += right-left+1
    }
    return count
}