func majorityElement(nums []int) int {
    ele := 0
    count := 0
    for i := 0; i < len(nums); i++ {
        if count == 0 {
            count++
            ele = nums[i]
        } else if nums[i] == ele {
            count++
        } else {
            count--
        }
    }
    count = 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == ele {
            count++
        }
    }
    if count > len(nums)/2 {return ele}
    return -1
}

/*
    majority element is a number that appears more than n/2 times
    approach: freq map
    - maintain a freq for each number
    - if at any point a number freq > n/2 times, we have found the majority element
    
    time = o(n)
    space = o(n)
*/
// func majorityElement(nums []int) int {
//     freqMap := map[int]int{}
//     for i := 0; i < len(nums); i++ {
//         freqMap[nums[i]]++
//         if val := freqMap[nums[i]]; val > len(nums)/2 {return nums[i]}
//     }
//     return -1
// }