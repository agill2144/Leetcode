func majorityElement(nums []int) int {
    count := 0
    ele := -1
    for i := 0; i < len(nums); i++ {
        if count == 0 {
            count = 1
            ele = nums[i]
        } else if nums[i] == ele{ 
            count++
        } else {
            count--
        }
    }
    return ele
}
/*
    approach: brute force
    - freq map
    tc = o(n)
    sc = o(n)
*/
// func majorityElement(nums []int) int {
//     n := len(nums)
//     freq := map[int]int{}
//     for i := 0; i < n; i++ {
//         freq[nums[i]]++
//         if freq[nums[i]] > n/2 {return nums[i]}
//     }
//     return -1
// }