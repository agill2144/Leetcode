/*
    approach: binary search on answers
    does the answer lie within a known range?
    - yes, from smallest size of 1 to biggest size ( max element ) 
    - "Return the maximum possible positive integer"
        - binary search hint, ( but does not fit max val is min || min val is max )
        - therefore checking if linear range loop works
        - and in this case, it does
    - then we can replace the linear loop with a binary search loop
    time = o( log(max) * n )
    space = o(1)
*/
func maxLength(ribbons []int, k int) int {
    right := math.MinInt64
    left := 1
    for i := 0; i < len(ribbons); i++ {
        right = max(right, ribbons[i])
    }
    ans := 0
    for left <= right {
        mid := left + (right-left) / 2
        // if the answer to evaluate is mid ( size of ribbon is mid )
        // how many ribbons of the same size can we get ?
        count := 0
        for i := 0; i < len(ribbons); i++ {
            count += (ribbons[i]/mid)
        }
        // when does mid work?
        // if we were able to create atleast ( or more ) than k ribbons
        // mid works
        if count >= k {
            // save this as a potential ans
            ans = mid
            // but keep searching right, since we want the highest such answe
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return ans
}

/*
    approach: linear scan over a range of possible answers
    - we have ribbons, of a specific size
    - we can cut a ribbon into many pieces
    - after cutting a ribbon, we will have some pieces of different lenghts ( depending on where we cut )
    - we want to find out, the max size of a cut we can have which results into atleast k ribbons
        - we can collect several k from same ribbon
        - and we can also collect k from different ribbons
    - we want the max such lenght
    - so, then lets be greedy and start with highest possible length we can get
        - i.e max size of a ribbon
        - [9,7,5] = 9 ; k=3
    - then at each element, we will ask, how many 9 sized ribbons can i get at this ith ribbob
        - how many 9's make up this ith element
        - ithElement/9
        - 9/9 = 1 ( meaning we can have 1 ribbon this element of size 9 )
        - 7/9 = 0 ( meaning we cannot get a ribbon of size 9 from this element )
        - 5/9 = 0 ( meaning we cannot get a ribbon of size 9 from this element )
        - we could only collect 1 ribbon of size 9 if our desired size was 9.
        - we need atleast ( at min ) k such ribbons of desired size ;
    - lets try 8
        - 9/8 = 1 
        - 7/8 = 0
        - 5/8 = 0
        - didnt work, could not collect k=3 ribbons of size 8
    - lets try 7 ( results into 2 ; not equal to k )
    - lets try 6 ( results into 2 ; not equal to k )
    - lets try 5
        - 9/5 = 1 ( we can get 1 ribbon out of 9 of size 5 )
        - 7/5 = 1 ( we can get 1 ribbon out of 7 of size 5 )
        - 5/5 = 1 ( we can get 1 ribbon out of 5 of size 5 )
        - we got 3 total ribbons of size 5
        - this meets our k constraint and is the max possible ans
        - therefore we can return 5        
    - we can keep going until size 1
        - this is our last valid size as a potential answer

    - the key takeway was;
        - answer lies within a range
        - that range is from 1 ( smallest possible answer ) till 9 ( max possible answer )

    time = o( max(ribbons) * n )
    space = o(1)
*/
// func maxLength(ribbons []int, k int) int {
//     maxSize := math.MinInt64
//     start := 1
//     for i := 0; i < len(ribbons); i++ {
//         maxSize = max(maxSize, ribbons[i])
//     }

//     for i := maxSize; i >= start; i-- {
//         desiredSize := i
//         count := 0
//         for j := 0; j < len(ribbons); j++ {
//             count += (ribbons[j] / desiredSize)
//         }
//         if count >= k {return desiredSize}
//     }
//     return 0
// }