func maxProduct(nums []int) int {
    sort.Ints(nums)
    n1 := nums[len(nums)-1]
    n2 := nums[len(nums)-2]
    return (n1-1) * (n2-1)
}