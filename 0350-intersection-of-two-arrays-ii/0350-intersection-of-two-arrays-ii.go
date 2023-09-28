/*
    collecting equal values from both arrays
    approach: two pointers
    - since the values matter, their positioning does not matter, we can always sort!
    - sort opens up 2 options = two ptrs and binary search
    - put ptrs on each array
    - and if the ptr values are same in both array
    - save this value, move both ptrs ahead
    - if not same values, which ptr to move ?
        - smaller value ptr moves forward. why ?
        - [5,x,x,x]
        - [7,x,x]
        - if ptr is at 5 and 7, these values are not equal
        - moving ptr at 7 is not going to run into 5 because the array is sorted in asc order
        - moving ptr at 5 can potentinally run into a 7, therefore smaller value moves ahead
    
    - if we are counting sorting
    time = o(n1logn1) + o(n2logn2) + o(min(n1,n2))
    space = o(n1) + o(n2)
*/
func intersect(nums1 []int, nums2 []int) []int {
    sort.Ints(nums1)
    sort.Ints(nums2)
    n1 := 0
    n2 := 0
    out := []int{}
    for n1 < len(nums1) && n2 < len(nums2) {
        if nums1[n1] == nums2[n2] {
            out = append(out, nums1[n1])
            n1++
            n2++
        } else if nums1[n1] < nums2[n2] {
            n1++
        } else {
            n2++
        }
    }
    return out
}

/*
    approach: we are collecting equal numbers from both array
    - pick the largest array and create freqMap
        - why not a set?
        - because a number in larger array could be only once, but 5 times in smaller array
    - loop over the smaller array
    - check if this element exists in freqMap and that we still have counts available 
    - if we do, store this number in output array and decrement the freqmap val by 1
    
    
    time = o(n1) + o(n2)
    space = o(n1) 
*/
// func intersect(nums1 []int, nums2 []int) []int {
//     if len(nums2) > len(nums1) {
//         return intersect(nums2, nums1)
//     }
    
//     // nums1 is always the larger array
//     freqMap := map[int]int{}
//     for i := 0; i < len(nums1); i++ {
//         freqMap[nums1[i]]++
//     }
    
//     out := []int{}
//     for i := 0; i < len(nums2); i++ {
//         if val, ok := freqMap[nums2[i]]; ok && val != 0 {
//             out = append(out, nums2[i])
//             freqMap[nums2[i]]--
//         }
//     }
//     return out
// }