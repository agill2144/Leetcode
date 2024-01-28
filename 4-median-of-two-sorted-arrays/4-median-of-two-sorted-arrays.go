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
    if len(nums2) < len(nums1) {return findMedianSortedArrays(nums2, nums1)}
    
    n1 := len(nums1)
    n2 := len(nums2)
    total := n1+n2
    half := total/2
    // partX = points to nums1 array
    // partY = points to nums2 array
    
    left := 0
    right := n1
    for left <= right {
        mid := left + (right-left)/2
        partX := mid
        partY := half-partX

        // find out the left numbers from both arrays ( to the left of partX and partY )
        l1 := math.MinInt64
        if partX != 0 {l1 = nums1[partX-1]}
        l2 := math.MinInt64
        if partY != 0 {l2 = nums2[partY-1]}
        // find out the right numbers from both arrays ( to the right of partX and partY )
        r1 := math.MaxInt64
        if partX != n1 {r1 = nums1[partX]}
        r2 := math.MaxInt64
        if partY != n2 {r2 = nums2[partY]}
        
        // are we looking at correct partitions?
        // is this a perfect split ?
        if l1 <= r2 && l2 <= r1 {
            if total % 2 == 0 {
                return (math.Max(float64(l1), float64(l2)) + math.Min(float64(r1), float64(r2))) / 2.0
            }
            // odd len
            return math.Min(float64(r1), float64(r2))
        }
        
        if l1 > r2 {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return -1
}

/*
    approach: optimizing last brute force ( using the same merge sorted array intuition )
    - instead of actually allocating a 3rd array
    - use an imaginary array
    - median = middle element
    - therefore if total len of combined array is $total
        - the middle element will be at total/2 idx position
        - if the total len was even, we need 2 middle positions
        - in that case, the middle elements will be at total/2 and (total/2)-1
        - therefore idx1 = total/2 and idx2 = (total/2)-1
    - so we know at what idx positions in this imaginary array the middle elements will be at
    - then we dont really need the actual array
    - we can just keep track of the idx we are currently placing an element in our imaginary sorted array
    - create element vars for both idx1 and idx2 ( ele1 and ele2 )
    - if current idx we are placing at is;
        - either idx1 or idx2
        - save the smaller value in the corresponding idx1,idx2 element vars
    - once we have both elements (ele1 and ele2)
    - we can calc our median
    - if total len is even
        - return (ele1+ele2) / 2
    - if total len is odd
        - we only care about one of the element vars
        - and that will be ele1
        - return ele1
    
    time = o(n1+n2)
    space = o(1)
*/
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//     total := len(nums1) + len(nums2)

//     idx1 := total/2
//     ele1 := math.MinInt64
    
//     idx2 := (total/2)-1
//     ele2 := math.MinInt64
    
    
//     idx, n1, n2 := 0,0,0
//     for n1 < len(nums1) && n2 < len(nums2) {
//         n1Val := nums1[n1]
//         n2Val := nums2[n2]
//         if n1Val < n2Val {
//             if idx == idx1 {
//                 ele1 = n1Val
//             } else if idx == idx2 {
//                 ele2 = n1Val
//             }
//             idx++
//             n1++
//         } else {
//             if idx == idx1 {
//                 ele1 = n2Val
//             } else if idx == idx2 {
//                 ele2 = n2Val
//             }
//             idx++
//             n2++
//         }
//         if ele1 != math.MinInt64 && ele2 != math.MinInt64 {
//             break
//         }
//     }
    
//     for n1 < len(nums1) {
//         n1Val := nums1[n1]
//         if idx == idx1 {
//             ele1 = n1Val
//         } else if idx == idx2 {
//             ele2 = n1Val
//         }
//         idx++
//         n1++        
//         if ele1 != math.MinInt64 && ele2 != math.MinInt64 {
//             break
//         }
//     }
    
//     for n2 < len(nums2) {
//         n2Val := nums2[n2]
//         if idx == idx1 {
//             ele1 = n2Val
//         } else if idx == idx2 {
//             ele2 = n2Val
//         }
//         idx++
//         n2++        
//         if ele1 != math.MinInt64 && ele2 != math.MinInt64 {
//             break
//         }
//     }
    
//     if total % 2 == 0{
//         return (float64(ele1) + float64(ele2)) / 2.0
//     }
//     return float64(ele1)
// }

/*
    median = middle element
    middle when arr size is even = mid + mid+1 / 2
    middle when arr size is odd  = mid
    
    approach: brute force
    - merge 2 sorted arr
    - find the middle
    
    time = o(n1+n2)
    space = o(n1+n2)
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
    
//     for n1 < len(nums1) {merged = append(merged, nums1[n1]); n1++}
//     for n2 < len(nums2) {merged = append(merged, nums2[n2]); n2++}

//     mid := 0+(len(merged)-1-0)/2
//     if len(merged) % 2 == 0 {
//         return (float64(merged[mid]) + float64(merged[mid+1])) / 2.0
//     }
//     return float64(merged[mid])
// }
