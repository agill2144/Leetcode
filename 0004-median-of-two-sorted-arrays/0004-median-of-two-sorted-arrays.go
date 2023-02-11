/*
    approach : find parition ( split ) such that all elements on the left of partition are < elements on right side of partition
    - binary search to not search for a target, 
    - but instead our target is where to place a partition such that all elements on the left of partition are < elements on right side of partition 
    - to find the correct partition, use binary search
    - if this was a 1d array, then to split the array in half, we would do len/2
        - this would create 2 sides of equal size ( approx equal )
    - then if we have 2 arrays, pretend like they are 1 merged array
        - a = [1,2,3,4] ; b = [3,4,5,6,7,8]
        - now we need to place 2 partitions ( 1 on each array ), such that number of elements on left and right side are approx half of full len(n1+n2)
        - lets call this partitionA and partitionB
        - partitionA will be on array "a" and partitionB will be on array "b"
        - then if we put a partition in array "a" after idx 0, ( between 1 and 2 values )
        - this means that left half of the array has already allocated 1 element
        - inorder to complete the left half, how many more elements can the left half have?
            - n1+n2/2 - numberOfElementsAlreadyAllocatedOnLeftSide
            - 5+6/2 - 1
            - 11/2 -1
            - 5-1 = 4
        - so the left half can have 4 more elements on the left side ( if we were splitting a pretend merged array into approx equal halves.)
        - this means, from array "b" we can place partition such that there are 4 elements on the left of partition
        - Now combined left half of "a" and "b" we have approx equal halves on the left side and on combined right side.
    - Now we need to identify where the partition we placed in both "a" and "b" are correct
        - partition is correct IF ALL ELEMENTS to the left <= ALL ELEMENTS ON RIGHT SIDE
        - we have 2 arrays... both arrays are sorted.
        - we only have to check numbers immediately to the left and compare with the rights number of partition
        - we will have 4 pointers ( 2 on array "a" and 2 on array "b")
            - l1 = immediate left of partitionA in "a" array
            - l2 = immediate left of parititonB in "b" array
            - r1 = immediate right of partitionA in "a" array
            - r2 = immediate right of partitionB in "b" array
        - if both left values are less than both right values, this means we have found the perfect split in both arrays.
            - then if total number of elements are odd (n1+n2);
                - median is min(r1, r2)
            - if total number of elements are even (n1+n2)
                - median is max(l1, l2) + min(r1,r2) / 2
        
    
    time: o(log(min(nums1,nums2)))
    space: o(1)
    
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if len(nums2) < len(nums1) {return findMedianSortedArrays(nums2, nums1) }
    
    n1 := len(nums1)
    n2 := len(nums2)
    
    left := 0
    right := n1
    
    for left <= right {
        partX := left + (right-left)/2
        partY := (n1+n2)/2 - partX
        
        l1 := math.MinInt64
        if partX != 0 {l1 = nums1[partX-1]}
        l2 := math.MinInt64
        if partY != 0 {l2 = nums2[partY-1]}
        
        r1 := math.MaxInt64
        if partX < n1 {r1 = nums1[partX]}
        r2 := math.MaxInt64
        if partY < n2 {r2 = nums2[partY]}
        
        if l1 <= r2 && l2 <= r1 {
            if (n1+n2) % 2 == 0 {
                return (math.Min(float64(r1), float64(r2)) + math.Max(float64(l1), float64(l2))) / 2.0
            } else {
                return math.Min(float64(r1), float64(r2))
            }
        } else if l1 > r2 {
            right = partX-1
        } else {
            left = partX+1
        }
    }
    return -1
}   


/*
    approach : merge + sort and then grab middle
    - merge both arrays into a new m+n array
    - Sort the new array
    - if len(newArray) is even, grab the middle and return
    - or else grab the 2 middle elements and return the avg between them (even len array)
    
    time: o(m+nlogm+n)
    space: o(1)
*/
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//     merged := []int{}
//     for i := 0; i < len(nums1); i++ {
//         merged = append(merged, nums1[i])
//     }
//     for i := 0; i < len(nums2); i++ {
//         merged = append(merged, nums2[i])
//     }
//     sort.Ints(merged)
//     midIdx := len(merged) / 2
//     if len(merged) % 2 != 0 {
//         // odd array len, return the mid
//         return float64(merged[midIdx])
//     }
//     var avg float64 = float64(merged[midIdx]+merged[midIdx-1])/2.0 
//     return avg
// }


/*
    approach : merge using two pointers so no need to sort, and then grab middle
    - merge both arrays into a new m+n array using 2 pointers
    - if len(newArray) is even, grab the middle and return
    - or else grab the 2 middle elements and return the avg between them (even len array)
    
    time: o(m+n)
    space: o(1)
*/
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//     merged := []int{}
//     n1 := 0
//     n2 := 0
//     for n1 < len(nums1) && n2 < len(nums2) {
//         if nums1[n1] < nums2[n2] {
//             merged = append(merged, nums1[n1])
//             n1++
//         } else {
//             merged = append(merged, nums2[n2])
//             n2++
//         }
//     }
//     if n1 < len(nums1) {
//         merged = append(merged, nums1[n1:]...)
//     }
//     if n2 < len(nums2) {
//         merged = append(merged, nums2[n2:]...)
//     }
    
//     midIdx := len(merged) / 2
//     if len(merged) % 2 != 0 {
//         // odd array len, return the mid
//         return float64(merged[midIdx])
//     }
//     var avg float64 = float64(merged[midIdx]+merged[midIdx-1])/2.0 
//     return avg
// }



/*
    brute force: 2 heaps ( inspired by Find median from data stream )
    - idea is the same, we have 1 sorted array ( we merge nums1 and nums2 ) - o(n1+n2) time and space
    - median, means we only care about middle elements
    - we can use 2 heaps (min and max)
        - where min is used on the right half of the array
        - and max is used on the left half of the array
        - this way middle elements are at the top of both heaps
    
    if N = n1+n2 elements
    time ;
        - o(N) to merge both into 1 array ( we could get away from doing this too)
        - for each element in N elements, we did 
            - heap.Push + [heap.Pop() + heap.Push()] + [heap.Pop() + heap.Push()]
            - o(5logN)
        - o(N) + o(N5logN)
    space: o(2N) or o(N)
*/
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//     merged := []int{}
//     n1 := 0
//     n2 := 0
//     for n1 < len(nums1) && n2 < len(nums2) {
//         if nums1[n1] < nums2[n2] {
//             merged = append(merged, nums1[n1])
//             n1++
//         } else {
//             merged = append(merged, nums2[n2])
//             n2++            
//         }
//     }
//     for n1 < len(nums1) { merged = append(merged, nums1[n1]) ; n1++ }
//     for n2 < len(nums2) { merged = append(merged, nums2[n2]) ; n2++ }
    
//     left := &maxHeap{items: []int{}}
//     right := &minHeap{items: []int{}}
    
    
//     for i := 0; i < len(merged); i++ {
//         heap.Push(left, merged[i])
//         // ensure left top <= right top
//         if left.Len() != 0 && right.Len() != 0 && 
//         left.items[0] > right.items[0] {
//             heap.Push(right, heap.Pop(left))
//         }
//         // ensure its balanced, so that the median is at the top of surface of both heaps
//         // otherwise median could get burried somewhere deep inside a heap
//         if left.Len() > right.Len() + 1 {
//             heap.Push(right, heap.Pop(left))
//         } else if right.Len() > left.Len() {
//             heap.Push(left, heap.Pop(right))
//         }

//     }
//     if len(left.items) > len(right.items) {
//         return float64(left.items[0])
//     } else if len(right.items) > len(left.items) {
//         return float64(right.items[0])
//     }
//     return (float64(left.items[0]) + float64(right.items[0]))/2.0
// }


// func abs(x int) int {
//     if x < 0 {return x * -1}
//     return x
// }




// type maxHeap struct {items []int}
// func(m *maxHeap) Less(i, j int)bool {return m.items[i] > m.items[j]}
// func(m *maxHeap) Swap(i, j int)     {m.items[i], m.items[j] = m.items[j], m.items[i]}
// func(m *maxHeap) Len()int           {return len(m.items)}
// func(m *maxHeap) Push(x interface{}) {m.items = append(m.items, x.(int))}
// func(m *maxHeap) Pop()interface{} {
//     out := m.items[len(m.items)-1]
//     m.items = m.items[:len(m.items)-1]
//     return out
// }

// type minHeap struct {items []int}
// func(m *minHeap) Less(i, j int)bool {return m.items[i] < m.items[j]}
// func(m *minHeap) Swap(i, j int)     {m.items[i], m.items[j] = m.items[j], m.items[i]}
// func(m *minHeap) Len()int           {return len(m.items)}
// func(m *minHeap) Push(x interface{}) {m.items = append(m.items, x.(int))}
// func(m *minHeap) Pop()interface{} {
//     out := m.items[len(m.items)-1]
//     m.items = m.items[:len(m.items)-1]
//     return out
// }