func subArrayRanges(nums []int) int64 {

    // sum of subarry minimums
    var minTotalSum int64
    st := []int{} // idx
    // nsr
    for i := 0; i < len(nums); i++ {
        for len(st) != 0 && nums[i] < nums[st[len(st)-1]] {
            // process top
            top := st[len(st)-1]
            st = st[:len(st)-1]
            topVal := nums[top]
            nsr := i
            nsl := -1
            if len(st) != 0 {nsl = st[len(st)-1]}
            leftCount := top-nsl
            rightCount := nsr-top
            totalCount := leftCount*rightCount
            minTotalSum += (int64(totalCount) * int64(topVal))
        }
        st = append(st, i)
    }

    for len(st) != 0 {
        // process top
        top := st[len(st)-1]
        st = st[:len(st)-1]
        topVal := nums[top]
        nsr := len(nums)
        nsl := -1
        if len(st) != 0 {nsl = st[len(st)-1]}
        leftCount := top-nsl
        rightCount := nsr-top
        totalCount := leftCount*rightCount
        minTotalSum += (int64(totalCount) * int64(topVal))
    }


    // sum of subarry maximums
    var maxTotalSum int64
    st = []int{}
    // ngr
    for i := 0; i < len(nums); i++ {
        for len(st) != 0 && nums[i] > nums[st[len(st)-1]] {
            // process top
            top := st[len(st)-1]
            st = st[:len(st)-1]
            topVal := nums[top]
            ngr := i
            ngl := -1
            if len(st) != 0 {ngl = st[len(st)-1]}
            leftCount := top-ngr
            rightCount := ngl-top
            totalCount := leftCount*rightCount
            maxTotalSum += (int64(totalCount) * int64(topVal))
        }
        st = append(st, i)
    }

    for len(st) != 0 {
        // process top
        top := st[len(st)-1]
        st = st[:len(st)-1]
        topVal := nums[top]
        ngr := len(nums)
        ngl := -1
        if len(st) != 0 {ngl = st[len(st)-1]}
        leftCount := top-ngl
        rightCount := ngr-top
        totalCount := leftCount*rightCount
        maxTotalSum += (int64(totalCount) * int64(topVal))
    }

    // once we have total max and min sums, the total range sum would be
    // max - min sum 
    return maxTotalSum - minTotalSum
}
