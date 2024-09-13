func maxSatisfied(customers []int, grumpy []int, minutes int) int {

    n := len(customers)
    total := 0
    for i := 0; i < n; i++ {
        if grumpy[i] == 0 {total += customers[i]}
    }

    maxCount := 0
    left := 0
    count := 0
    for i := 0; i < n; i++ {
        if grumpy[i] == 1 {
            count += customers[i]
        }
        if i-left+1 == minutes {
            maxCount = max(maxCount, count)
            if grumpy[left] == 1 {
                count-=customers[left]
            }
            left++
        }
    }
    return total + maxCount

}