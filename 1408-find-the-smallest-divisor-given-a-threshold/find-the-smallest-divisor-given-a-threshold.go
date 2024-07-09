func smallestDivisor(nums []int, threshold int) int {
    start := 1
    end := math.MinInt64
    for i := 0; i < len(nums); i++ {
        end = max(end, nums[i])
    }
    left := start
    right := end
    ans := -1
    for left <= right {
        divisor := left + (right-left)/2
        sum := 0
        for j := 0; j < len(nums); j++ {
            q := int(math.Ceil(float64(nums[j]) / float64(divisor)))
            sum += q
        }

        if sum > threshold {
            left = divisor+1
        } else {
            ans = divisor
            right = divisor-1
        }
    }
    return ans
}


// func smallestDivisor(nums []int, threshold int) int {
//     start := 1
//     end := math.MinInt64
//     for i := 0; i < len(nums); i++ {
//         end = max(end, nums[i])
//     }
//     for i := start; i <= end; i++ {
//         divisor := i
//         sum := 0
//         for j := 0; j < len(nums); j++ {
//             q := int(math.Ceil(float64(nums[j]) / float64(divisor)))
//             sum += q
//         }
//         if sum <= threshold {return divisor}
//     }
//     return -1
// }