/*
    approach: complement search
    - this is like two sum in a lot of ways (i.e complement search, sorting + two ptrs, sorting + binary search)
    - the differences is we would have to count all pairs vs just finding 1 pair
    - and the input array may contain duplicates, and so we have to count dupes as 1 pair only
        - for example (1,2) and (2,1) is the same pair ( values matter, their idx and how they are used as pairs do not matter )
    - The other difference is that we have find a pair, whose abs(diff) == k
    - if k = 2 and input array is [3,1,4,1,5]
        - i = 0; what number do we need to subtract from 3 (idx 0) to make it equal to 2
            - 3-x = 2; -x = 2-3; x = 1
            - so we need a 1. ( 3-1 does equal to 2 )
            - there is another way
            - 5-3 is also 2, i.e x = 2+3 ( 5-3 does equal to 2 )
            - so we have two complements to search for
        - SEARCH! so lets put nums in a idx map for now so we can search easily
    - To eliminate duplicate pairs
        - we will do what we did in 3sum
        - set of sorted pairs ( where each pair is sorted )
        - if we find a pair, we will sort this pair and check if we have seen/used this pair before ( check in our set of pairs )
        - increment our counter, if we have not seen this pair and then add it this sorted pair to our set of sorted pairs.
    - What if k is 0
        - nums =    [3,1,4,1,5]
        - idxMap =  {3:0,1:3,4:2,5:4}
        - and i = 0
        our two complements to search for will be;
        - 3+0 = 3
        - 3-0 = 3
        ... so 3 does exist, but we cannot use the same idx right!
        lucky for us, we have an idx map that tells us where 3's position is which is idx 0
        - so only consider a complement if found and its idx != current i
            - this way we wont add (3,3)...
        when i goes to idx 1 ( value 1 )
        our two complements to search for will be;
        - 1+0 = 1
        - 1-0 = 1
        - we have both complements, so therefore we will spearate them into 2 chunks of checks
        1. complement1 ; 1+0 = 1
        - 1 does exist and its idx in idxMap is 3 which != current i (2)
        - then check whether we have seen this pair in our sorted set. no we have not
        - use it, increment our counter and move on
        1. complement2 ; 1-0 = 1
        - 1 does exist and its idx in idxMap is 3 which != current i (2)
        - then check whether we have seen this pair in our sorted set. yes we have ( we just added this pair )
        - so skip
        
        when i goes to idx 3 ( value 1 )
        our two complements to search for will be;
        - 1+0 = 1
        - 1-0 = 1
        - we have both complements, so therefore we will spearate them into 2 chunks of checks
        1. complement 1 ; 1+0 = 1
        - 1 does exist and its idx in idxMap is 3 which == current i (3)
        - cant use 2 numbers from the same idx! skip and move on

    time: o(n) + o(n2log2) = o(n)
    time: o(n) + o(n) = o(n) ( was able to remove sorting 2log2 because we are just sorting a pair of 2 nums )
    space: o(n)
*/
func findPairs(nums []int, k int) int {
    idxMap := map[int]int{}
    for i := 0; i < len(nums); i++ {
        idxMap[nums[i]] = i
    }
    set := newSetOfPairs()
    count := 0
    
    for i := 0; i < len(nums); i++ {

        comp1 := nums[i]+k
        idx, comp1Found := idxMap[comp1]
        if comp1Found && i != idx {
            tmp := [2]int{comp1, nums[i]}
            if !set.containsPair(tmp) {
                count++
                set.addPair(tmp)
            }
        }
        
        comp2 := nums[i]-k
        idx2, comp2Found := idxMap[comp2]
        if comp2Found && i != idx2 {
            tmp := [2]int{comp2, nums[i]}
            if !set.containsPair(tmp) {
                count++
                set.addPair(tmp)
            }
        }
        
    }
    return count
}

type setOfPairs struct{
    items map[[2]int]struct{}
}
func newSetOfPairs() *setOfPairs{
    return &setOfPairs{items: map[[2]int]struct{}{}}
}
func (s *setOfPairs) addPair(n [2]int) {
    // sort it first
    if n[1] < n[0] {
        n[0], n[1] = n[1],n[0]
    }
    s.items[n] = struct{}{}
}
func (s *setOfPairs) containsPair(n [2]int)  bool {
    // sort it first
    if n[1] < n[0] {
        n[0], n[1] = n[1],n[0]
    }
    _, ok := s.items[n]
    return ok
}


/*
    approach 2: sort + two pointers
    - sort the nums first
    - then use two pointers
        - in two sum sorted input, we placed two pointers at the beginning and end of array
        - in this we need to place them in at same spot ( from beginning )
        - why ?
            - [3,1,4,1,5]
            -  ^       ^
            - sort it
            - [1,1,3,4,5]
            -  ^   ^
            - now if we place 2 pointers at the end like 2 sum
            - we will move the right ptr back to value 3 ( because 5-1 > 2, so right move back, 4-1 > 2 , so right move back, 3-1 == 2 )
            - but now we have lost 5 and 3 as pair...
        - therefore we will place ptr1 at idx 0 and ptr2 at idx 1
        - then diff ; nums[ptr2] - nums[ptr1]
        - if diff == k, move both ptrs!
            - However because there could be dupes, keep ptr2 moving by 1 if nums[ptr2] == nums[ptr2-1] ( just like we did in 3Sum / 2Sum to avoid dupes in sorted array )
        - if diff > k, move the ptr1 closer to ptr2
            - ptr1++
            - moving ptr2 ahead will only increase the diff since its a sorted array
        - otherwise diff < k , then move ptr2 ahead by 1 because diff is smaller than k, we need a bigger value 
        
    time: o(nlogn) + o(n)
    space: o(1)
*/
// func findPairs(nums []int, k int) int {
//     sort.Ints(nums)
//     count := 0
//     i := 1
//     j := 0
//     for i < len(nums) {
//         if i == j {i++; continue}
//         diff := nums[i]-nums[j]
//         if diff == k {
//             count++
//             i++
//             j++
//             for i < len(nums) && nums[i] == nums[i-1] {i++}
//         } else if diff > k {
//             j++
//         } else {
//             i++
//         }
//     }
//     return count
// }


