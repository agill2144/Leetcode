/*
    approach: brute force
    - using nested for loop, form every single subarray and check its sum
    - if sum == target, count++
    
    time: o(n^2)
    space: o(1)
    
    
    approach: running sum + complement search
    - to do complement search, we need some lookup data structure
        - map would be a good fit for this
    - We will have a running sum as we are looping over the nums array
    - at each runningSum, we will search for a number that will be needed to add to this runningSum to make it == target
    - in other words, complement search
    - what number do I need to add to this runningSum to make it == target
    - target-runningSum - will give you the complement number 
    - search for this complement in our map, if found, take its value and add to count
    - if not found, then save this runningSum and increment its count by 1 in the map
    time: o(n)
    space:o (n)
    
*/

func subarraySum(nums []int, k int) int {
    rsCount := map[int]int{0:1}
    count := 0
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        complement := sum - k
        count += rsCount[complement]
        rsCount[sum]++
    }
    return count
}