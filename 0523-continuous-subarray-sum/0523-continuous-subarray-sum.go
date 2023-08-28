/*
    approach: running sum
    - if we have a running sum upto idx x
    - how do we check if x is divisble by k ? 
        - if x % k == 0; it means that x is divisible by k
    - if x % k does not result into a 0; it means that this subarray is not divisible by k
    - How do we check previous nested subarrays that created x as running sum ?
    - If the remainder is not 0, it means the remainder is up by some value
        - lets say remainder is 5
        - it means that some running sum made the remainder up by 5
        - we need to make this remainder 0
        - by removing the number/runningSum that caused it to go up by 5
        - THEREFORE we need to search and keep track of each idx running sum remainders
        - if we have seen a 5 as a remainder before ( lets say at a previous idx y )
        - it means, that if we remove runningSum from 0 upto and including y, then we wont have a 
            a remainder up by 5 anymore
        - so we need to search for an idx that caused a remainder to by up some value
        - if we found it, we can check the size by keeping track of each remainder with its idx in a hashmap
        - if we did not find it, we need to save this remainder in the hashmap , maybe this could get used later on 
    
    time = o(n)
    space = o(n)

*/
func checkSubarraySum(nums []int, k int) bool {
    remainderIdx := map[int]int{0:-1}
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        remainder := sum % k
        
        idx, ok := remainderIdx[remainder]
        if ok && i-(idx+1)+1 >= 2{
            return true
        }
        if !ok {
            remainderIdx[remainder] = i
        }
    }
    return false
}