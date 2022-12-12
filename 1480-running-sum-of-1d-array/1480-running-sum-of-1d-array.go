// func runningSum(nums []int) []int {
//     rs := 0
//     out := []int{}
//     for i := 0; i < len(nums); i++ {
//         rs += nums[i]
//         out = append(out, rs)
//     }
//     return out
// }

func runningSum(nums []int) []int {
    if len(nums) <= 1 {
        return nums
    }
    for i := 1; i < len(nums); i++ {
        prev := nums[i-1]
        nums[i] = nums[i] + prev
    }
    return nums
}