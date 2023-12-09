func reversePairs(nums []int) int {
    count := 0
    var dfs func(left, right int)
    dfs = func(left, right int) {
        // base
        if left >= right {return}
        
        
        // logic
        mid := left + (right-left)/2
        dfs(left, mid)
        dfs(mid+1, right)
        
        // left side = i
        // right side = j
        // looking for nums[i] > nums[j]
        i := left
        for i <= mid {
            j := mid+1
            for j <= right {
                if nums[i] > nums[j]*2 {
                    count++
                    j++
                } else {
                    break
                }
            }
            i++
        }
        
        // merge sort
        aux := []int{}
        l := left
        r := mid+1
        for l <= mid && r <= right {
            if nums[l] < nums[r] {
                aux = append(aux, nums[l])
                l++
            } else {
                aux = append(aux, nums[r])
                r++
            }
        }
        for l <= mid {aux = append(aux, nums[l]); l++}
        for r <= right {aux = append(aux, nums[r]); r++}
        
        auxPtr := 0
        for i := left; i <= right; i++ {
            nums[i] = aux[auxPtr]
            auxPtr++
        }
    }
    dfs(0, len(nums)-1)
    return count
    
}