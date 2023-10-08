func reversePairs(nums []int) int {
    count := 0
    var dfs func(left, right int)
    dfs = func(left, right int) {
        // base
        if left == right {return}
        
        // logic
        mid := left + (right-left)/2
        dfs(left, mid)
        dfs(mid+1, right)
        
        
        j := mid+1
        for i := left; i <= mid; i++ {
            for j <= right {
                if nums[i] <= nums[j]*2 {
                    break
                }
                j++
            }
            count += j-(mid+1)
        }
        
        l := left
        r := mid+1
        aux := []int{}
        for l <= mid && r <= right {
            if nums[l] < nums[r] {
                aux = append(aux, nums[l])
                l++
            } else {
                aux = append(aux, nums[r])
                r++
            }
        }
        
        for l <= mid {aux = append(aux, nums[l]);l++}
        for r <= right {aux = append(aux, nums[r]);r++}
        
        
        auxPtr := 0
        for i := left; i <= right; i++ {
            nums[i] = aux[auxPtr]
            auxPtr++
        }
        
    }
    dfs(0, len(nums)-1)
    return count
}