func permute(nums []int) [][]int {

    arr := [][]int{}
    var helper func(start int)
    helper = func(start int) {
        // base
        if start == len(nums) {
            newL := make([]int, len(nums))
            copy(newL, nums)
            arr = append(arr, newL)
            return
        }
        
        // logic
        for i := start; i < len(nums); i++ {
            nums[i],nums[start] = nums[start],nums[i]
            helper(start+1)
            nums[i],nums[start] = nums[start],nums[i]            
        }
    }
    
    helper(0)
    return arr
    
}