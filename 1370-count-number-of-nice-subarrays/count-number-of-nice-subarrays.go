
// time = o(n)
// space = o(n)
// similar approach as runningSum pattern;
// https://leetcode.com/problems/subarray-sum-equals-k/
func numberOfSubarrays(nums []int, k int) int {
    return count(nums, k) - count(nums, k-1)
}

func count(nums []int, k int) int {
    left := 0
    c := 0
    oddCount := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] % 2 != 0 {
            oddCount++
        }
        for oddCount > k {
            if nums[left] % 2 != 0 {
                oddCount--
            }
            left++
        }
        c += (i-left+1)
    }
    return c
}