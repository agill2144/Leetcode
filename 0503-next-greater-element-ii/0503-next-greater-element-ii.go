func nextGreaterElements(nums []int) []int {
    out := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        out[i] = -1
    }
    n := len(nums)
    stack := []int{}
    for i := 0; i < 2 * n; i++ {
        for len(stack) != 0 && nums[i%n] > nums[stack[len(stack)-1]] {
            top := stack[len(stack)-1]
            out[top] = nums[i%n]
            stack = stack[:len(stack)-1]
        }
        
        stack = append(stack, i%n)
        
    }
    return out
}