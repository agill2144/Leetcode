/*
    approach: binary search
    - median = middle element
    - middle element = there are n/2 (half) elements on both sides  
    - when arr size is odd, median is mid element
    - when arr size is even, median is avg of 2 middle elements (mid+mid+1)/2
    - point is, we only care about 2 middle elements
    - we need to find places in both arrays, where middle elements are sorted
    - that is, place a partition in both arrays
        - ALL elements on left should be <= ALL elements on right
    - if we know that there are supposed to n/2 (half) elements on left and half on right
    - we only really need to place partition on 1 of the array
        - placing a partition on 1 array, gives us the number of elements on left hand side
        - and we just have to satisfy the n/2 num of elements on the left hand side
        - therefore if know n/2 and we know how many elements we picked from 1 array,
        - then second partition placement can be calculated as;
        - half - numOfElementOnLeftSoFar
    - we can place a partition before 0th idx
        - ie. do not pick any elements from this array for left hand side
    - we can plac a partition at nth idx
        - ie. pick all elements from this array for left hand side
    - so partitions can be placed from -1 -> n
    - this is a range
    - ranges are sorted
    - therefore we can use binary search on this partition placement loop
    - for example;
        - a = [1,2,3,4] ; b = [3,4,5,6,7,8]
        - now we need to place 2 partitions ( 1 on each array ), 
        - such that number of elements on left and right side are approx half of full len(n1+n2)
        - lets call this partitionA and partitionB
        - partitionA will be on array "a" and partitionB will be on array "b"
        - then if we put a partition in array "a" after idx 0, ( between 1 and 2 value )
            - [1 | 2,3,4]
        - this means that left half of the array has already allocated 1 element
        - inorder to complete the left half, how many more elements can the left half have?
            - n1+n2/2 - numberOfElementsAlreadyAllocatedOnLeftSide
            - (4+6)/2 - 1
            - 10/2 - 1
            - 5-1
            - 4
        - so the left half can have 4 more elements on the left side ( if we were splitting a pretend merged array into approx equal halves.)
        - this means, from array "b" we can place partition such that there are 4 elements on the left of partition
            - [3,4,5,6 | 7,8]
        - Now combined left half of "a" and "b" we have approx equal halves on the left side and on combined right side.
            - a = [1 | 2,3,4]
            - b = [3,4,5,6 | 7,8]

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
    if len(nums2) < len(nums1) {
        return findMedianSortedArrays(nums2, nums1)
    }
    n1 := len(nums1)
    n2 := len(nums2)
    total := n1+n2
    half := total/2
    left := 0
    right := n1
    for left <= right {
        mid := left + (right-left)/2
        mid2 := half-mid
        l1, l2 := math.MinInt64, math.MinInt64
        if mid-1 >= 0 {l1 = nums1[mid-1]}
        if mid2-1 >= 0 {l2 = nums2[mid2-1]}
        r1, r2 := math.MaxInt64, math.MaxInt64
        if mid < n1 { r1 = nums1[mid] }
        if mid2 < n2 {r2 = nums2[mid2]}

        if l1 <= r2 && l2 <= r1 {
            if total % 2 == 0 {
                return (max(float64(l1),float64(l2)) + min(float64(r1),float64(r2))) / 2.0
            }
            return min(float64(r1),float64(r2))
        }

        if l2 > r1 {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return 0.0
}
/*
    approach: optimized space on brute force
    - we just need 2 middle elements in the worst case
        - i.e our merged array could be even or odd
        - worst case its even and we need 2 middle elements
    - so instead of physically placing them into a merged array
    - we can just keep track of 
        1. the current idx we are placing in a fake array
        2. the current element that should be placed at this idx value in fake array
    - if idx is either one of the mid idx we care about (mid or mid+1)
        - we will save its corresponding value
    - then at the end the logic is the same
    - if total size is even
        - avg of both middle elements
    - if total size is off
        - return the first middle element
    
    time = o(n1+n2)
    space = o(1)
    
*/
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//     total := len(nums1)+len(nums2)
//     mid1 := (total-1)/2
//     mid2 := mid1+1
//     ele1, ele2 := 0,0
//     n1, n2 := 0,0
//     idx := 0
//     for n1 < len(nums1) && n2 < len(nums2) {
//         if nums1[n1] < nums2[n2] {
//             if idx == mid1 {
//                 ele1 = nums1[n1]
//             } else if idx == mid2 {
//                 ele2 = nums1[n1]
//             }
//             n1++
//             idx++
//         } else {
//             if idx == mid1 {
//                 ele1 = nums2[n2]
//             } else if idx == mid2 {
//                 ele2 = nums2[n2]
//             }
//             n2++
//             idx++
//         }
//     }
//     for n1 < len(nums1) {
//         if idx == mid1 {
//             ele1 = nums1[n1]
//         } else if idx == mid2 {
//             ele2 = nums1[n1]
//         }
//         n1++
//         idx++
//     }
//     for n2 < len(nums2) {
//         if idx == mid1 {
//             ele1 = nums2[n2]
//         } else if idx == mid2 {
//             ele2 = nums2[n2]
//         }
//         n2++
//         idx++
//     }
//     if total % 2 == 0 {
//         return (float64(ele1) + float64(ele2)) / 2.0
//     }
//     return float64(ele1)
// }

/*
    approach: brute force
    - using two ptrs, create a 3rd arr
    - find the middle element in the 3rd array
    - if arr size is even
        - median is the avg of 2 middle elements
        - [x,x,x, x,x,x]
    - if arr size is odd
        - median is the middle elment
        - [x,x,x, x, x,x,x]

    n1 = len(nums1)
    n2 = len(nums2)
    n3 = n1+n2
    time = o(n1 + n2)
    space = o(n3)
*/
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//     sortedArr := []int{}
//     n1 := 0
//     n2 := 0
//     for n1 < len(nums1) && n2 < len(nums2) {
//         if nums1[n1] < nums2[n2] {
//             sortedArr = append(sortedArr, nums1[n1])
//             n1++
//         } else {
//             sortedArr = append(sortedArr, nums2[n2])
//             n2++
//         }
//     }
//     for n1 < len(nums1) {sortedArr = append(sortedArr, nums1[n1]); n1++}
//     for n2 < len(nums2) {sortedArr = append(sortedArr, nums2[n2]); n2++}
//     mid := (len(sortedArr)-1) / 2
//     if len(sortedArr) % 2 != 0 {
//         return float64(sortedArr[mid])
//     }
//     return (float64(sortedArr[mid]) + float64(sortedArr[mid+1])) / 2.0    
// }
