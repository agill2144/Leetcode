/*
    approach: slow and fast ptr
    - slow ptr acts as the idx position where the next non zero item should be placed at
    - fast ptr iterates the nums arrat
    - when fast ptr runs into a non-zero num ( i.e nums[fast] != 0 )
        - then slow ptr needs to take this number at swap it with current slow ptr val
        - nums[slow],nums[fast] = nums[fast],nums[slow]
        - and then move the slow ptr forward to tell where the next non-zero item should be.
    
    time: o(n)
    space: o(1)
    
    [0,1,0,3,12]
     s
       f - fast ptr is not pointing to a 0
----------------------------------------------------
    [1,0,0,3,12]
       s
           f - fast ptr is not pointing to a 0 -- swap!
----------------------------------------------------
    [1,3,0,0,12]
         s
             f - fast ptr is not pointing to a 0 -- swap!
----------------------------------------------------
    [1,3,12,0,0]
         s
                f (for loop breaks)
*/
func moveZeroes(nums []int)  {
   
    slow := 0
    fast := 0
    
    for fast < len(nums) {
        if nums[fast] != 0 {
            nums[slow],nums[fast] = nums[fast],nums[slow]
            slow++
        }
        fast++
    }
}