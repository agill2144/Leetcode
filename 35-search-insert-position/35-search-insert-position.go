/*
    approach: Binary search
    - Either look for target or ceilingOfTarget/floorOfTarget
    - ceiling : greater than target but the smallest one
        - [1,2,4,5,6] , target=3 --- ceiling would be 4
    - floor : smaller than target but the greatest one 
        - [1,2,4,5,6] , target=3 --- floor would be 2
        
    - for this lets look for floor using binary search
    - if perfect target match is found, return that idx
    - otherwise if the mid < target, save this as a potential answer
    - but continue looking right, as we may have undershot our mid
    - so look right for a greater if possible
        - no need to search left since anything left mid ( thats is already < target will be smaller ) and we want smallest but greater/closest to target
    - if we did not find the exact target 
        - finally return idx+1 because we were looking for smallest/closest from left side to target
        - and we need the next position where its suitable for this target to be inserted
        - and target is best suitable after the largest floor we find
    time: o(logn)
    space: o(1)
    
*/

func searchInsert(nums []int, target int) int {
    left := 0
    right := len(nums)-1
    idx := -1
    
    for left <= right {
        mid := left+(right-left)/2
        if nums[mid] <= target {
            if nums[mid] == target {
                return mid
            }
            idx = mid
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return idx+1
}