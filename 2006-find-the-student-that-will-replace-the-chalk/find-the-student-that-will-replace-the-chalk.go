func chalkReplacer(chalk []int, k int) int {
    total := 0
    for i := 0; i < len(chalk); i++ {total += chalk[i]; if total > k {break}}
    // we could simulate taking a round trip in a for loop
    // or, we can do it using mod operator
    // if k = 22 and total = 11;
    // it means, we can take 2 full trips from idx 0 to n-1
    // so now we can use a for loop to remove total from k ( while k >= total { k-= total })
    // then we can use a idx ptr to figure out which idx position we will end up at
    // otherwise, we can use mod operator
    // if k = 24 and total = 11
    // then k % total ; 24 % 11; leaves a remainder of 2 
    // with this mod operator; we have essentially ran the trip from idx 0 to n-1 2 times ( becuase quotient will also be 2 )
    // leaving us with remaining k value ( i.e the remainder ) = 2
    // now we can use a idx ptr to simulate 
    k %= total
    n := len(chalk)
    for i := 0; i < n; i++ {
        if k >= chalk[i] {
            k -= chalk[i]
        } else {
            return i
        }
    }
    return -1
}