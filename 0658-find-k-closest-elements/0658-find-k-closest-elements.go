/*
    approach: store a pair(val, distAway) in a maxHeap
    - For each n elements
    - Calc the distance : arr[n] - x
    - Create a pair {val, distAway}
    - Store in maxHeap ( maxHeap sorted by distAway so the farthest pair will be on the top )
    - If the maxHeap size > k heap.pop()
    - Once all the elements are done,
    - The maxHeap will be left with k pairs that are the closest to x
    - Loop over maxHeap items and store them in result array and return
    
    time: o(n) * o(logk) = o(nlogk)
    - why k? because in maxHeap we will only store k elements
    - where did the n come from? We loop over all elements to store them in maxHeap therefore o(n) * o(logk)
    space: o(k) - in maxHeap at worse, we will store k elements
    However I still had to sort the final result -- Did I do something wrong?  (+ o(klogk) to time for sorting the final k size output list)
    
    
    approach: 2 pointers
    - left ptr at 0
    - right ptr at n-1
    - while the distance between left and right != k
        - i.e until we do not have a window of size k
        - get leftDist to x
        - get rightDist to x
        - whoever is farther from x, move away from that value 
            - if rightDist is farther from x compared to leftDist, then right--
            - else left++
        - If the distance is the same, then move away from bigger value
    - once we have a valid window of size k
    - then loop from left to right and populate the result arr
    - time: o(n-k)+o(k)
    - space: o(1)
    
    
    approach: Binary search to look for the starting point of the range
- left = 0
- right = len(arr)-k
    - Why not last idx?
    - BECAUSE we are doing binary search to look for THE STARTING IDX OF THE RANGE OF ELEMENTS
    - THE LAST VALID STARTING IDX WITHIN ARR WOULD BE len(arr)-k right... therefore ( do not forget we are doing binary search to find a VALID starting point)
- Then we get the mid
- When we are doing binary search we always use mid to see if thats our target right?
- but since we are looking for a starting point from within our range, we will use mid to see IF mid is a potential starting point...
- but how do we check that?
- if mid == start point then that means mid+k-1 would be the k size range right?
- so if mid is truly the starting point, then distance from mid to x must be smaller than of distance from endPoint to x
    - endPoint here is mid+k ( i.e the last element in this range )
- how do we get the distance from mid to x? -> x-arr[mid]
- how do we get the distance from endPoint to x -> arr[mid+k]-x
- In our two pointers, whichever distance was closer, we moved towards that, ( or whichever distance was farthest from x we moved away from that )
- Same principle applies here, if startDist > endDist - move away from startPtr ( left ptr ) 
- Otherwise move away from endDist (right ptr )
- What about when startDist == endDist?
    - In two pointers and the question itself states, that if two distances are equal use the smaller value
    - Then in our binary search we will do the same
    - But how? and what would that look like...
    - If startDist/leftDist == endDist/rightDist then we obv do not want to lean towards endDist/rightDist -- so right = mid ( move away from current right ptr )
- Once our binary search breaks out of the loop, then we can use the left ptr to pull out arr[left:left+k-1] elements (k closest elements)

Time: o(logn-k) to search + o(k) for the final list generation
Space: o(1)
*/

// binary search to find the starting point of the range
func findClosestElements(arr []int, k int, x int) []int {
    left := 0
    // the last valid starting point inorder to get a valid k size array
    right := len(arr)-k
    
    for left < right {
        mid := left + (right-left)/2
        distStart := x-arr[mid]
        distEnd := arr[mid+k]-x
        if distStart > distEnd {
            // move away from mid because start is way farther from end 
            // so mid which is potential starting point is not correct, since the end of this range is alot closer than the start
            left = mid+1
        } else {
            right = mid
        }
    }
    
    out := []int{}
    for i := left; i < left+k; i++ {
        out = append(out, arr[i])
    }
    return out
}

// two pointer linear approach
// time: o(n-k)
// why -k ? 
//  because as soon as we hit our windowSize of k , we never loop over those inner k window elements, therefore we never looped over entire n length arr but rather n-k len array
// space: o(1)
// func findClosestElements(arr []int, k int, x int) []int {
//     left := 0
//     right := len(arr)-1
//     for right-left+1 > k {
//         rightDist := abs(arr[right]-x)
//         leftDist := abs(arr[left]-x)
//         if (leftDist < rightDist) || (leftDist == rightDist) {
//             right--
//         } else if rightDist < leftDist {
//             left++
//         }
//     }
//     out := []int{}
//     for i := left; i <= right; i++ {
//         out = append(out, arr[i])
//     }
//     return out
// }


// approach: max heap
// func findClosestElements(arr []int, k int, x int) []int {
//     mx := &maxHeap{items: []*item{}}
//     for i := 0; i < len(arr); i++ {
//         it := &item{dist: abs(arr[i]-x), val: arr[i]}
//         heap.Push(mx, it)
//         if mx.Len() > k {
//             heap.Pop(mx)
//         }
//     }
//     out := []int{}
//     for mx.Len() != 0 {
//         out = append(out, heap.Pop(mx).(*item).val)
//     }
//     sort.Ints(out)
//     return out
// }
// type item struct {
//     dist, val int
// }
// type maxHeap struct {
//     items []*item
// }
// func(m *maxHeap) Len() int {return len(m.items)}
// func (m *maxHeap) Swap(i,j int) {m.items[i],m.items[j] = m.items[j],m.items[i]}
// func(m *maxHeap) Less(i, j int)bool{
//     if m.items[i].dist == m.items[j].dist {
//         return m.items[i].val > m.items[j].val
//     }
//     return m.items[i].dist > m.items[j].dist
// }
// func(m *maxHeap) Push(x interface{}) {
//     m.items = append(m.items, x.(*item))
// }
// func(m *maxHeap) Pop() interface{} {
//     out := m.items[len(m.items)-1]
//     m.items = m.items[:len(m.items)-1]
//     return out
// }
// func abs(n int) int {
//     if n < 0 {
//         return n * -1
//     }
//     return n
// }