func longestOnes(nums []int, k int) int {
    ones := 0
    maxWin := 0
    left := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 1 {ones++}
        size := i-left+1
        zeros := size-ones
        if zeros <= k {
            maxWin = max(maxWin, i-left+1)
        } else {
            if nums[left] == 1 {ones--}
            left++
        }
    }
    return maxWin
}
// func longestOnes(nums []int, k int) int {
//     ones := 0
//     maxWin := 0
//     left := 0
//     for i := 0; i < len(nums); i++ {
//         if nums[i] == 1 {ones++}
//         size := i-left+1
//         zeros := size-ones
//         for zeros > k {
//             if nums[left] == 1 {ones--}
//             left++
//             size = i-left+1
//             zeros = size-ones
//         }
//         maxWin = max(maxWin, i-left+1)
//     }
//     return maxWin
// }