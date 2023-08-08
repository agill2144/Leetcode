// when we cannot determine a sorted half, shrink our window! FUCKING SHIT
// because of this worst case TC = o(n) ; example: [1,1,1,1,1,1,2,1,1,1,1] ; target = 2
// our target is surrounded by n-1 dupes, therefore our left and right ptr see most of the elements , i.e seeing all elements
// avg tc = o(logn)
// space = o(1)

func search(nums []int, target int) bool {
    left := 0
    right := len(nums)-1
    
    for left <= right {
        mid := left+(right-left)/2
        if nums[mid] == target {return true}
        
        // we are at a point where we cannot confidently make a decision on picking a sorted side
        // therefore shrink our window and continue searching
        if nums[mid] == nums[left] && nums[mid] == nums[right] {
            left++
            right--
            continue
        }
        
        // now classic ; search in rotated sorted array by finding the sorted half
        if nums[left] <= nums[mid] {
            if target >= nums[left] && target <= nums[mid] {
                right = mid-1
            } else {
                left = mid+1
            } 
        } else {
            if target <= nums[right] && target >= nums[mid] {
                left = mid+1
            } else {
                right = mid-1
            }
        }
    }
    return false
}