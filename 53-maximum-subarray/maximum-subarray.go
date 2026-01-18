// kadanes greedy algo, try to evaluate each idx as a new starting point
// start a new eval if rSum < 0 
// same logic can be applied to gas station problem too ( eval whether a gas station works, drop and start new if it does not work in the middle .. )
func maxSubArray(nums []int) int {
	rSum := 0 
	res := math.MinInt64
	for i := 0; i < len(nums); i++ {
		rSum += nums[i]
		res = max(res, rSum)
		if rSum < 0 {rSum = 0}
	}
	return res
}
// brute force, form all possible subarrays
// tc = n^2
// sc = o(1)
// func maxSubArray(nums []int) int {
//     n := len(nums)
//     res := math.MinInt64
//     sum := 0
//     for i := 0; i < n; i++ {
//         for j := i; j < len(nums); j++ {
//             sum += nums[j]
//             res = max(res, sum)
//         }
//         sum = 0
//     }
//     return res
// }