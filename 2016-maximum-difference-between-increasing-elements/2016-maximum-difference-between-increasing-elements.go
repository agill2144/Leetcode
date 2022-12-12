func maximumDifference(nums []int) int {
    /*
        we will keep track of a min ( imply idx 0 at start ).
        Then our j loop will start at idx 1 ( j > i )
        Then our j loop will check:
            1. if j val bigger than min val ? 
                - yes, do a diff and save it
            2. else if j val is smalled than min val
                - yes, then save this j val as the min
                - why? because inorder to get the max diff, we need the smallest number on min 
    */
    
    if nums == nil || len(nums) == 0 {return -1}
    diff := 0
    min := nums[0]
    
    for j := 1; j < len(nums); j++ {
        if nums[j] > min {
            if nums[j]-min > diff {
                diff = nums[j]-min
            }
        } else if nums[j] < min {
            min = nums[j]
        }
    }
    if diff == 0 {
        diff = -1
    }
    return diff
}