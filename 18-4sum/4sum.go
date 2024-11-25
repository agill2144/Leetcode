func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    result := [][]int{}
    n := len(nums)
    
    for i := 0; i < n-3; i++ {
        if i > 0 && nums[i] == nums[i-1] { continue}
        for j := i+1; j < n-2; j++ {
            /*
                [2,2,2,3,4....]
                 i j
                - we want the first j to get processed even if jth val is same as j-1th val
                - because this is the first time we are making quads starting with [2,2,x,y]

                [2,2,2,3,4....]
                 i   j
                - now we can skip this jth ptr pointing at the same value as prev value
                - because continuing with this will create another quad starting with [2,2,x,y]
                    that we had discovered in the last iteration ( above )

                - therefore j > i+1 for j to skip curr position if curr position value is same as prev
            */
            if j > i+1 && nums[j] == nums[j-1] {continue}
            left, right := j+1, n-1
            for left < right {
                sum := nums[i] + nums[j] + nums[left] + nums[right] 
                if sum == target {
                    result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})                    
                    left++
                    for left < right && nums[left] == nums[left-1] {left++}
                    right--
                    for left < right && nums[right] == nums[right+1] {right--}                    
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