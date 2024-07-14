/*
    input is sorted, binary search ? 
    
    approach:
    - search for left most of or equal to x
    - then deploy 2 ptrs ( idx and idx+1 )
    - only collect elements using 2 ptrs that are closer dist to x
    - if both dist are equal, collect left ( discard right )
    
    time = o(logn) + o(k)
    space = o(1)
    
    but there is a better binary search ( fastest one yet )
    
    approach: better binary search
    - instead of searching for the target x, we will now search for a starting position of our k size window
    - if mid position is our starting point, then the next elements would be mid+k-1 
        - now we need to check whether this is our k size window of closest elements possible
        - how ?
        - we know the start position ( mid idx )
        - we also know the end position ( mid+k-1 idx )
        - we can compute distances for both
        - if there is a large imbalance in their distances, it means one side has too many far-away elements
        - if start dist > end dist , move away from start position ; i.e left = mid+1
        - otherwise it means end dist too far away from x, move away from right side ; i.e right = mid-1

        - what if the distances were equal ? we cant decide which side to pick ... 
            - shrink window like search in rotated sorted 2 ?  No
            - instead of shrinking the window, we know that if 2 distances are equal, we can discard higher val ( right value )
            - inorder to gurantee k elements, we will take 1 extra in our window each time
            - when we run into a equal dist, we can for sure discard right element and still have exactly k elements left ( since we took 1 extra on right side )
            - therefore the range becomes mid = start, mid+k = end of the window
        - keep doing this until binary search ends
    - once binary search is over, left:left+k-1 are all k closest to x

    time = o(logn-k) + o(k) to extract later
    space = o(1)
    
*/
func findClosestElements(arr []int, k int, x int) []int {
    left := 0
    right := len(arr)-k
    for left <= right {
        mid := left + (right-left)/2
        if mid + k >= len(arr) {break}
        startDist := x-arr[mid]
        endDist := arr[mid+k]-x
        if startDist > endDist {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return arr[left:left+k]
}

/*
    input is sorted, two pointers ?
    approach: two pointers
    - left and right ptr
    - check their dist
    - move away from farther one
    - if both dist are equal, move away from higher value ( i.e right )
    - until number of elements within left and right == k
    
    time = o(n-k) + o(k)
    space = o(1)
*/
// func findClosestElements(arr []int, k int, x int) []int {
//     left := 0
//     right := len(arr)-1
//     for right - left + 1 != k {
//         leftDist := abs(arr[left]-x)
//         rightDist := abs(arr[right]-x)
//         if leftDist > rightDist {
//             left++
//         } else {
//             right--
//         }
//     }
//     return arr[left:right+1]
// }

/*
    "k" closest
    approach: maxHeap
    - anything "k" closest/smallest/largest/farthest consider heap
    - heap will keep track of closest elements to x
    - compute the dist for each element
    - store a pair in maxHeap {num, dist}
        - maxHeap will store by dist
        - if dist of i == j , then sort by num
        - maxHeap will gurantee the farthest pair is at the top
        - while keep the closest one at the bottom of the tree safe
    - if our heap exceeds size k, drop the farthest pair ( i.e top of heap )
    
    time = o(nlogk) + o(klogk) for sorting
    space = o(k)
*/
// func findClosestElements(arr []int, k int, x int) []int {
//     mx := &maxHeap{items: [][]int{}}
//     for i := 0; i < len(arr); i++ {
//         dist := abs(arr[i]-x)
//         heap.Push(mx, []int{arr[i], dist})
//         if mx.Len() > k {
//             heap.Pop(mx)
//         }
//     }
//     out := []int{}
//     for mx.Len() != 0 {
//         val := heap.Pop(mx).([]int)[0]
//         out = append(out, val)
//     }
//     sort.Ints(out)
//     return out
// }


/*
    type Interface interface {
        sort.Sort
        Push()
        Pop()
    }
*/
// type maxHeap struct {
//     items [][]int // {num, dist}
// }

// func (m *maxHeap) Less(i, j int) bool {
//     iDist := m.items[i][1]
//     jDist := m.items[j][1]
//     if iDist == jDist {
//         return m.items[i][0] > m.items[j][0]
//     }
//     return iDist > jDist
// }

// func (m *maxHeap) Swap(i, j int) {
//     m.items[i], m.items[j] = m.items[j], m.items[i]
// }

// func (m *maxHeap) Len() int {
//     return len(m.items)
// }

// func (m *maxHeap) Push(x interface{}) {
//     m.items = append(m.items, x.([]int))
// }

// func (m *maxHeap) Pop() interface{} {
//     out := m.items[len(m.items)-1]
//     m.items = m.items[:len(m.items)-1]
//     return out
// }

/*
    approach: brute force
    - compute distances for each number and store them in a map
    - store by {dist : [nums]}
    - then run a loop from minDist -> maxDist collecting elements until we have k elements
    - once we have k elements, sort our array, return the array
*/
// func findClosestElements(arr []int, k int, x int) []int {
//     dists := map[int][]int{}
//     minDist := math.MaxInt64
//     maxDist := math.MinInt64
//     for i := 0; i < len(arr); i++ {
//         diff := abs(arr[i]-x)
//         dists[diff] = append(dists[diff], arr[i])
//         minDist = min(diff, minDist)
//         maxDist = max(diff, maxDist)
//     }
//     out := []int{}
//     for i := minDist; i <= maxDist; i++ {
//         for _, v := range dists[i] {
//             out = append(out, v)
//             if len(out) == k {sort.Ints(out); return out}            
//         }
//     }
//     return out
// }


func abs(x int )int {
    if x < 0 {return x *-1}
    return x
}

func min(x, y int) int {
    if x < y {return x}
    return y
}

func max(x, y int) int {
    if x > y {return x}
    return y
}