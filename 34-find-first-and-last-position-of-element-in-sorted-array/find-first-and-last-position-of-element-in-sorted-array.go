/*
    approach: linear search
    - as soon as we run into target
    - check if start idx is set
        - if not, save this idx as start idx
    - and keep saving this same idx also as end idx
        - its possible that there is only 1 occurence of this element
        - so save this idx everytime
    - keep traversing next set of elements until element != target
    - return the start and end indicies
    time = o(n)
    space = o(1)
*/
func searchRange(nums []int, target int) []int {
    start, end := -1, -1
    for i := 0; i < len(nums); i++ {
        num := nums[i]
        if num == target {
            if start == -1 {
                start = i
            }
            end = i
        }
    }
    return []int{start, end}
}