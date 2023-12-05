/*

*/
func missingNumber(nums []int) int {
    actualSum := 0
    n := len(nums)
    for i := 0; i < n; i++ {
        actualSum += nums[i]
    }
    expectedSum := n*(n+1)/2
    return expectedSum-actualSum
}


/*
    approach: set
    - nums are supposed to be within 0 and n ( inclusive )
    - toss nums in a set so we can easily check whether a ele exists
    - then run a loop from 0 to n and return the first number we dont find in set
    time: o(2n)
    space = o(n)
*/
// func missingNumber(nums []int) int {
//     set := map[int]struct{}{}
//     n := len(nums)
//     for i := 0; i < len(nums); i++ {
//         set[nums[i]] = struct{}{}
//     }
//     for i := 0; i <= n; i++ {
//         if _, ok := set[i]; !ok {return i}
//     }
//     return n
// }