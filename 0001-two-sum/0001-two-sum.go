/*
    why cant we use running sum pattern here?
    - we can only use running sum pattern when looking for contiguous subarrays!
    - in this problem, we can pick any 2 two numbers, they do not have to be contigous
    - The running sum pattern will give a contigous subarray RANGE and not 2 specific indicies - IT GIVES A SUBARRAY!
    - If the question was find total pairs of 2 contagious numbers that sum to target, running sum can help
*/

// compliment search using hashmap
// time: o(n)
// space: o(n)
func twoSum(nums []int, target int) []int {
    idxMap := map[int]int{}
    for i := 0; i < len(nums); i++ {
        num := nums[i]
        // question to ask yourself
        // How do I make current $num == $target
        // What number adding to $num will be equal to target
        // Thats target-$num
        // and if we have seen such number, return its idx and current $num idx
        if idx, ok := idxMap[target-num]; ok {
            return []int{idx, i}
        }
        // otherwise, add current number in our idxMap as it may be a compliment for another number later.
        idxMap[num] = i
    }
    return nil
}

// if the nums array was sorted AND WE WERE RETURNING THE NUMS THAT ADD UP INSTEAD OF INDICIES, then use two pointers!!!!!
// func twoSum(nums []int, target int) []int {
//     sort.Ints(nums)
//     left := 0
//     right := len(nums)-1
//     for left < right {
//         sum := nums[left] + nums[right]
//         if sum == target {
//             return []int{nums[left] , nums[right]}
//         } else if sum > target {
//             right--
//         } else {
//             left++
//         }
//     }
//     return []int{-1,-1}
// }