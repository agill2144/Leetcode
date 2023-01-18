// as soon as we encounter a 'c', a potential frog ( currently under investigation ) is found
// each new c == potential frog
// increase number of potentialFrogs by 1
// and save if potentialFrogs seen > out 
// potentialFrog != completeFrog
// a potentialFrog becomes complete AFTER we have seen its last char "k"
// what about the chars in the middle? and their sequence ( are they in correct order ? )
// we will count the number of times we have seen each (croak)
// if some char is out of order, then its frequency will be higher than its correct previous char right ?
// example: craok, c = 1, r = 1, a = 1
// but "a" should have arrived AFTER "o", in other words if "a" count > "o" count it means we have seen more 'a' than we have seen "o"
// therefore this is out of order, and we can return -1 then and there!
// so order / sequence check of chars in a string can be done using char freq count and if a char is out of order, its frequency WILL BE HIGHER THAN ITS PREVIOUS NEIGHBORING CHAR
// we will do the above check everytime we increment a char count ( except for 'c'. because there is no char in 'croak' that comes before c)

// when 'k' is found, a potential frog is now a complete frog!
// therefore reduce the count of potential frogs seen currently under investigation by 1
// this works well for "croakcroak" - sequential, not multiple frogs
// Example: croakcroak - wouldnt there be 2 'c's here? and therefore 2 frogs?
// Yes but we ONLY COUNT COMPLETE FROGS

func minNumberOfFrogs(croakOfFrogs string) int {
    c := 0
    r := 0
    o := 0
    a := 0
    k := 0
    potentialFrogs := 0
    out := -1
    
    for i := 0; i < len(croakOfFrogs); i++ {
        char := croakOfFrogs[i]
        if char == 'c' {
            c++
            potentialFrogs++
        } else if char == 'r' {
            r++
            if r > c {return -1}
        } else if char == 'o' {
            o++
            if o > r {return -1}
        } else if char == 'a' {
            a++
            if a > o {return -1}
        } else if char == 'k' {
            k++
            if k > a {return -1}
            if potentialFrogs > out {
                out = potentialFrogs
            }
            potentialFrogs--
        }
    }
    if potentialFrogs != 0 {return -1}
    return out
}