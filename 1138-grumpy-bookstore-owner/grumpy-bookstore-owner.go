func maxSatisfied(customers []int, grumpy []int, minutes int) int {
    satisfied := 0
    for i := 0; i < len(customers); i++ {
        if grumpy[i] == 0 {satisfied += customers[i]}
    }
    left := 0 
    c := 0
    maxC := 0
    for i := 0; i < len(customers); i++ {
        if grumpy[i] == 1 {c += customers[i] }
        if i-left+1 == minutes {
            maxC = max(maxC, c)
            if grumpy[left] == 1 {c -= customers[left]}
            left++
        }
    }

    return satisfied + maxC
}

/*
1,0,1,2,1,1,7,5
0,1,0,1,0,1,0,1

7,1,1,1 = 10

*/