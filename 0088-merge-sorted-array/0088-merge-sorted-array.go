func merge(nums1 []int, m int, nums2 []int, n int)  {
    
    /*
    approach 1
    merge them blindly
    and then sort at the end
    
        time: o(m+nlogm+n)
        space: o(1)
    */

    /*
    approach 2
    create another array of size m+n
    and use 2 pointers and start pushing items from n1 and n2 in asc order
    
        time : o(m+n)
        space: o(m+n)
    */
    
    
    /*
    approach 3
    using three pointers, fill out the nums1 array in desc order ( from right to left )
    */
    last := m+n-1
    mPtr := m-1
    nPtr := n-1
    
    for last >= 0 {
        
        if mPtr >= 0 && nPtr >= 0 {
            mVal := nums1[mPtr]
            nVal := nums2[nPtr]
            
            if mVal > nVal {
                nums1[last] = mVal
                mPtr--
            } else {
                nums1[last] = nVal
                nPtr--
            }
        }else if nPtr >= 0 {
            nums1[last] = nums2[nPtr]
            nPtr--
        }else if mPtr >= 0 {
            nums1[last] = nums1[mPtr]
            mPtr--
        }
        last--
    }
    
}