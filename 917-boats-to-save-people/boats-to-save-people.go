func numRescueBoats(people []int, limit int) int {

    sort.Ints(people)
    left := 0
    right := len(people)-1
    count := 0
    for left <= right {
        if people[left] + people[right] <= limit {
            left++
            right--
            count++
        } else if  people[right] <= limit {
            right--
            count++
        } else {
            left++
            count++
        }
    }
    return count
}