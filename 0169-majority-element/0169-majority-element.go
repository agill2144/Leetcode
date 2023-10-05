/*
    The majority element is the element that appears more than ⌊n / 2⌋ times
    therefore an element that has a count of > n/2, is the majority element
    
    approach: freqMap
    - we can keep track of each number freq in a hashmap
    - as soon the value of a number in hashmap == (n/2) + 1
    - return that element
    - otherwise return -1 ( no such element > n/2 counts )
*/

func majorityElement(nums []int) int {
    n := len(nums)
    freqMap := map[int]int{}
    for i := 0; i < len(nums); i++ {
        freqMap[nums[i]]++
        if freqMap[nums[i]] > n/2 {
            return nums[i]
        }
    }
    return -1
}