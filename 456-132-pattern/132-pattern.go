func find132pattern(nums []int) bool {
    n := len(nums)
    numsK := math.MinInt64
    numsJ := math.MinInt64
    st := []int{} // idx
    for i := n-1; i >= 0; i-- {


        // once we have numsK, numsJ becomes irrelevant
        // because numsJ was the largest value once discovered; initialized/defined numsK val
        // therefore numsK < numsJ is already satisfied when numsK gets set
        // if this number is < numsK, it means we have found nums[i] < nums[k]
        if numsK != math.MinInt64 && nums[i] < numsK && numsK < numsJ {return true}

        // we have a ele thats > than immediate neighbors
        // meaning we have the peak element of 132 pattern
        // i.e numsJ value is num[i]
        // and numsK value should be just smaller on right side of numsJ
        // i.e looking back, last-in values in stack
        for len(st) != 0 && nums[i] > nums[st[len(st)-1]] {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            numsK = nums[top]
            numsJ = nums[i]
        }
        
        // push this value into st, as it could become a numsK value when we run into a new peak element
        st = append(st, i)
    }
    return false
}


// WORST FUCKING QUESTION EVER
// time = o(n)
// space = o(n)
// func find132pattern(nums []int) bool {
//     st := []int{} // indices, viable numsK candidates
//     n := len(nums)
//     numsK := math.MinInt64
//     for i := n-1; i >= 0; i-- {
//         // once we have numsK, numsJ becomes irrelevant
//         // because numsJ was the largest value once discovered and processed/setted numsK val
//         // therefore nums[k] < nums[j] is already satisfied is numsK is set
//         // if this number is < numsK, it means we have found nums[i] < nums[k]
//         // return true
//         if nums[i] < numsK {return true}

//         // the iterator ptr is acting as our numsJ val
//         /// numsJ is supposed to the largest and its supposed to set numsK value
//         // as soon as we run into a ele > top of stack(potential numsK vals), we have found nums[k] < nums[j] (and j < k)
//         // process the stack until nums[i] > top of stack, and each popped element is a better numsK val
//         // therefore each popped element becomes numsK ( overriden each time )
//         for len(st) != 0 && nums[i] > nums[st[len(st)-1]] {
//             numsK = nums[st[len(st)-1]]
//             st = st[:len(st)-1]
//         }

//         // push this value as potential numsK val
//         st = append(st, i)
//     }
//     return false
// }