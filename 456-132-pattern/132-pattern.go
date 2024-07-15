func find132pattern(nums []int) bool {
    n := len(nums)
    st := []int{} // idx
    numsJ := math.MinInt64
    numsK := math.MinInt64
    for i := n-1; i >= 0; i-- {
        if numsK != math.MinInt64 && nums[i] < numsK && numsK < numsJ {
            return true
        }
        // we have ran into a number thats larger that last seen
        // therefore we imply that this ( nums[i] ) is numsJ val and we are looking for numsK in st ( largest value )
        for len(st) != 0 && nums[i] > nums[st[len(st)-1]] {
            numsJ = nums[i]
            top := st[len(st)-1] // new 2nd largest; new kth val
            st = st[:len(st)-1]
            numsK = nums[top]
        }
        st = append(st, i) // potential k value
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