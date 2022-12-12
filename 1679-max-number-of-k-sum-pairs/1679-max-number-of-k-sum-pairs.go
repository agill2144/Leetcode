/*
    like two sum, but here we are just counting how many pairs exist 
    use compliment search technique
    time: o(n)
    space: o(n)
*/
// func maxOperations(nums []int, k int) int {
//     freqMap := map[int]int{}
//     count := 0
//     for i := 0; i < len(nums); i++ {
//         num := nums[i]
//         if val, ok := freqMap[k-num]; ok {
//             freqMap[k-num]--
//             if val == 1 {delete(freqMap, k-num)}
//             count++
//         } else {
//             freqMap[num]++
//         }
//     }
//     return count
// }


/*
    like two sum, but here we are just counting how many pairs exist 
    use sort + two pointers
    time: o(n) + o(nlogn)
    space: o(1)
*/
func maxOperations(nums []int, k int) int {
    sort.Ints(nums)
    count := 0
    left := 0
    right := len(nums)-1
    for left < right {
        leftNum := nums[left]
        rightNum := nums[right]
        sum := leftNum + rightNum
        if sum == k {
            count++; left++; right--
        } else if sum > k {
            right--
        } else {
            left++
        }
    }
    return count
}