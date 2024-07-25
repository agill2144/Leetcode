func sortArray(nums []int) []int {
    var dfs func(left, right int)
    dfs = func(left, right int){
        // base
        if left >= right {return}

        // logic
        mid := left + (right-left)/2
        dfs(left, mid)
        dfs(mid+1, right)

        
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
        for l <= mid { aux = append(aux, nums[l]); l++ }
        for r <= right { aux = append(aux, nums[r]); r++ }

        i := left
        for a := 0; a < len(aux); a++ {
            nums[i] = aux[a]
            i++
        }

    }
    dfs(0, len(nums)-1)
    return nums
}