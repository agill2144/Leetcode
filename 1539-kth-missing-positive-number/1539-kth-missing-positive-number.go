/*
    approach: brute force
    - create a list of missing numbers
    - if kth missing lies within our missing list
        - if k-1 idx exists within the missing numbers list
        - return value at missing[k-1]
    - otherwise kth missing lies outside missing list
        - so that means arr[len(arr)-1]+k, HOWEVER
        - we have take into account that number of missing elements we have found
        - therefore arr[len(arr)-1]+k-len(missing)
    
    time = o(maxVal)
    space = o(maxVal)
*/
func findKthPositive(arr []int, k int) int {
    missing := []int{}
    counter := 1
    i := 0
    for i < len(arr) {
        if arr[i] == counter {
            i++
            counter++
        } else {
            missing = append(missing, counter)
            counter++
        }
    }
    if k-1 <= len(missing)-1 {return missing[k-1]}
    return arr[len(arr)-1]+k-len(missing)
}