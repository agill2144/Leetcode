// the brute force works....
func rangeSum(nums []int, n int, left int, right int) int {
    sums := []int{}
    for i := 0; i < n; i++ {
        s := 0
        for j := i; j < n; j++ {
            s+=nums[j]
            sums = append(sums, s)
        }
    }
    sort.Ints(sums)
    ans := 0
    mod := 1000000007
    for i := left-1; i < right; i++ {
        ans = (ans + sums[i]) % mod
    }
    return ans
}