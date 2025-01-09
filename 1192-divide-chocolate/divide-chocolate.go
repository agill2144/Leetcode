/*
    approach: binary search on answers
    - our brute force has a outter loop running from min->totalSum
    - this is a range
    - ranges are sorted
    - we can replace this linear range loop with binary search

    time = o( log(sum-min+1) * n )
    space = o(1)

    Key take away
    - always check wether our ans lies within a range or not
        - see if a simple range loop can give us a brute force solution
    - minimum total sweetness is maximized
    - min val is maximized ( binary search on answers ) 
    - max val is minimized ( binary search on answers )

*/
func maximizeSweetness(sweetness []int, k int) int {
    left := math.MaxInt64
    right := 0
    for i := 0; i < len(sweetness); i++ {
        left = min(sweetness[i],left)
        right += sweetness[i]
    }

    ans := 0
    for left <= right {
        mid := left + (right-left)/2
        // if at min, each chunk sweetness level is mid
        // how many chunks can we create?
        chunks := 0
        sum := 0
        for j := 0; j < len(sweetness); j++ {
            sum += sweetness[j]
            // is atMin sweetness level met ?
            if sum >= mid {
                // start creating a new chunk
                sum = 0
                chunks++
            }
        }

        // when does atMin work ?
        // if we have created atleasr k+1 chunks
        if chunks >= k+1 {
            // save this as potential answer
            ans = mid
            // continue searching in right, since we want max possible ans
            left = mid+1
        } else {
            // when we could not create atleast k+1 chunks (or more)
            // there is no point in continuing our search on right hand side our range
            // because all the next numbers are higher 
            // meaning if atMin=x didnot work, why would x+1, x+2, x+3 .. x+max work?
            // it would not, therefore search left for a smaller sweetness level
            right = mid-1
        }
    }
    return ans
}
/*
    approach: brute force
    - always check wether our ans lies within a range or not
    - in this question it does
    - max possible sweetness is sum of all numbers
    - min possible sweetness is smallest num in array
    - try each sweetness from min -> max ( inclusive ) and calc;
        - number of chunks that can be created if **atMin** our sweetness level is i (i from min to max)
        - if we are able to make atleast k+1 chunks, then we good, save this i as potential ans
        - but continue further , as we may be able to get a better ans
        - also, the question states to maximize our sweetness level
        - when do we stop ?
            - as soon as number of chunks being created are less than k+1, we can stop
            - it means that, our desired sweetness level we were asking for was so big that we could not create k+1 chunks
                - meaning each chunk was able to have desired sweetness level, but doing so,
                        we could not meet the requirement of creating k+1 chunks
    time = o((sum-min+1) * n)    
    space = o(1)
*/

// func maximizeSweetness(sweetness []int, k int) int {
//     start := math.MaxInt64
//     end := 0
//     for i := 0; i < len(sweetness); i++ {
//         start = min(sweetness[i],start)
//         end += sweetness[i]
//     }

//     ans := 0
//     for i := start; i <= end; i++ {
//         // if at min, each chunk sweetness level is i
//         // how many chunks can we create?
//         atMin := i
//         chunks := 0
//         sum := 0
//         for j := 0; j < len(sweetness); j++ {
//             sum += sweetness[j]
//             // is atMin sweetness level met ?
//             if sum >= atMin {
//                 // start creating a new chunk
//                 sum = 0
//                 chunks++
//             }
//         }

//         // when does atMin work ?
//         // if we have created atleasr k+1 chunks
//         if chunks >= k+1 {
//             ans = atMin
//             continue
//         }

//         // we can stop trying if we end up creating < k+1 chunks
//         // because all the next numbers are higher 
//         // meaning if atMin=x didnot work, why would x+1, x+2, x+3 .. x+max work?
//         // it would not, therefore break and exit
//         break
//     }
//     return ans
// }



/*
    approach: ultimate brute force
    - using for-loop based recursion, make all possible chunks
    - keep track of number of chunks
    - keep track of minSweetness
    - when num of chunks == k+1; update our global ans;

    time = n^k ?
    space = o(k+1); max depth of recursion stops at k+1th node
*/
// func maximizeSweetness(sweetness []int, k int) int {
//     if k > len(sweetness) {return -1}
//     out := math.MinInt64
//     var dfs func(start int, chunks int, minSweetness int)
//     dfs = func(start, chunks, minSweetness int) {
//         // base
//         if chunks == k+1 {
//             out = max(out, minSweetness)
//             return
//         }

//         // logic
//         chunkSweetness := 0
//         for i := start; i < len(sweetness); i++ {
//             chunkSweetness += sweetness[i]            
//             dfs(i+1,chunks+1,min(minSweetness, chunkSweetness))
//         }
//     }
//     dfs(0, 0,math.MaxInt64)
//     return out
// }