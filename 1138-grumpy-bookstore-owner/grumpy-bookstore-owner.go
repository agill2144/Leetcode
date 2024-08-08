func maxSatisfied(customers []int, grumpy []int, minutes int) int {
    satisfied := 0
    n := len(customers)
    for i := 0; i < n; i++ {if grumpy[i] == 0 {satisfied += customers[i]}}    

    left := 0
    maxCustomers := 0
    count := 0 // count number of customers in $minutes window when the store owner is grumpy
    for i := 0; i < n; i++ {
        if grumpy[i] == 1 {count += customers[i]}
        if i-left+1 == minutes {
            maxCustomers = max(maxCustomers, count)
            if grumpy[left] == 1 {count -= customers[left]}
            left++
        }
    }
    return satisfied + maxCustomers
}