func minNumberOfFrogs(croakOfFrogs string) int {

    c := 0
    r := 0
    o := 0
    a := 0
    k := 0
    activeFrogs := 0
    ans := -1
    
    for i := 0; i < len(croakOfFrogs); i++ {
        char := croakOfFrogs[i]
        if char == 'c' {
            c++
            activeFrogs++
            // as soon as we encounter a c, a potential frog is found
            // each new c == potential frog
            // increase number of active frogs by 1
            // and save if activeFrogs seen > ans 
            ans = max(ans, activeFrogs)
        } else if char == 'r' {
            r++
        } else if char == 'o' {
            o++
        } else if char == 'a' {
            a++
        } else if char == 'k' {
            // when k is found, a potential frog has finished croaking!
            // therefore reduce the count of active frogs
            // this works well for "croakcroak" - sequential, not multiple frogs
            k++
            activeFrogs--
        }
        // to do sequential check of char, each char must appear before its prev char
        // when a char appears out of order, i.e not sequential, then its freq will be greater than its prev char
        // check each char freq with its prev char
        // if a char freq > prev char freq, then a char is out of order ( not sequential )
        if r > c || o > r || a > o || k > a {return -1}
    }
    
    // if an activeFrog is still waiting to be end ( i.e missing k ), activeFrogs wont be 0
    // when thats the case, return -1
    if activeFrogs != 0 {return -1}
    return ans
}

func max(x, y int) int {
    if x > y {return x}
    return y
}