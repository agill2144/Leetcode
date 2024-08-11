/*
    1. we already have a set of satisfied customers when store owner is not grumpy - therefore count them first
    2. next, we care about finding a window of size $minutes where we have max customers that are unsatisfied
        - i.e customer count where store owner is grumpy in a $minute size window
        - because this window can be satisfied.
        - if there many such windows, we care about window that gives us the most amount of customers
            when the store owner was grumpy
*/
func maxSatisfied(customers []int, grumpy []int, minutes int) int {
    satisfied := 0
    for i := 0; i < len(customers); i++ {
        if grumpy[i] == 0 {satisfied += customers[i]}
    }

    maxCount := 0
    left := 0
    count := 0 // count of customers when store owner is grumpy
    for i := 0; i < len(customers); i++ {
        if grumpy[i] ==  1 {count+=customers[i]}
        if i-left+1 == minutes {
            maxCount = max(maxCount, count)
            if grumpy[left] == 1 {count -= customers[left]}
            left++
        }
    }
    return maxCount + satisfied
}