/*
    approach: Brute force, BFS
    - usually we have used BFS for connected components on matrixes and or trees
    - but BFS can also be used for arrays
    - We have a start position and we also have a desintation
    - We have to get to the destination in minimum num of jumps as possible
        - almost looking for the shortest/smallest path to a destination
    - How do we figure out the jump count?
        - Thats each level, every level we process, if we did not reach destination in that level, jumps++
    - How and what do we add in our queue?
        - Since we are dealing with numbers in arrays that can be duplicated,
        - we will store index of each number
        - How would we get an index ?
        - We know our starting index is 0, so initially that will be added
        - Now we have to add 0's babies
        - in this context, the babies are all the babies that idx 0 can reach to
        - that would be all indicies from
            - 1..valueAtThatIdx + idx
            - so we enqueue all indicies that we can jump to from idx 0 (while being greedy and starting from the farthest jump first)
            - while we are adding each child index into our queue, how do we make sure, the same child does not get added back again?
            - we could do 2 things
                - maintain a visited set that maintains all the indicies we have in our queue so far
                    - why not just search the queue? because searching in queue takes o(n) time while searching in set will be o(1) time
                - or mark the number negative and have the adding of indicies to queue, filter our negative values at a idx we are evaluating

*/
func jump(nums []int) int {
    visited := map[int]struct{}{
        0: struct{}{},
    }
    q := []int{0}
    jumps := 0
    
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            if dq >= len(nums)-1 {return jumps}

            for i := nums[dq]; i >= 1; i-- {
                nextIdx := dq + i
                if nextIdx >= len(nums)-1 {return jumps+1}
                _, ok := visited[nextIdx]
                if !ok {
                    visited[nextIdx] = struct{}{}
                    q = append(q, nextIdx)
                }
            }
            qSize--
        }
        jumps++
    }
    return jumps
} 

/*
    approach: greedy
    - Using 2 intervals
    - Current interval = nums[0]
    - Next interval = num[0]
    - We will only increase our jump count once we have reached the current interval
    - While our i is headed towards current interval,
        - we will look for the next number that makes us jump the farthest distance possible
        - and we will maintain that with nextInterval
    - Once our i has reached the current interval idx, then we will increase our jump count by 1
    - and we will promote our next interval to current interval -- i.e do not increase the jump count until i == current interval and i has not reached the end
    - Finally return the jump count
    
    time: o(n)
    space: 0(1)

*/
// func jump(nums []int) int {
//     currInt := nums[0]
//     nextInt := nums[0]
//     jumps := 1
//     if len(nums) < 2 {return 0}

//     for i := 1; i < len(nums); i++ {
//         nextInt = int(math.Max(float64(nextInt), float64(i+nums[i])))
//         if i == currInt && i != len(nums)-1 {
//             jumps++
//             currInt = nextInt
//         }
//     }
//     return jumps
// }