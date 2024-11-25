func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    result := [][]int{}
    n := len(nums)
    
    for i := 0; i < n-3; i++ {
        // Skip duplicates for i
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        
        for j := i+1; j < n-2; j++ {
            // Skip duplicates for j
            if j > i+1 && nums[j] == nums[j-1] {
                continue
            }
            
            // Use two pointers for the remaining two numbers
            left, right := j+1, n-1
            
            for left < right {
                sum := nums[i] + nums[j] + nums[left] + nums[right]
                
                if sum == target {
                    result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
                    
                    // Skip duplicates for left
                    for left < right && nums[left] == nums[left+1] {
                        left++
                    }
                    // Skip duplicates for right
                    for left < right && nums[right] == nums[right-1] {
                        right--
                    }
                    
                    left++
                    right--
                } else if sum < target {
                    left++
                } else {
                    right--
                }
            }
        }
    }
    
    return result
}