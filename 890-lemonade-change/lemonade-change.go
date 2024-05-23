func lemonadeChange(bills []int) bool {
    fives := 0
    tens := 0
    twentys := 0

    for i := 0; i < len(bills); i++ {
        curr := bills[i]
        diff := curr-5
        if diff == 5 {
            if fives == 0 {return false}
            fives--
        } else if diff == 10 {
            if tens >= 1 {
                tens--
            } else if fives >= 2 {
                fives-=2
            } else {
                return false
            }
        } else if diff == 15 {
            if tens >= 1 && fives >= 1 {
                tens--
                fives--
            } else if fives >= 3 {
                fives -= 3
            } else {
                return false
            }
        }

        if curr == 5 {fives++}
        if curr == 10 {tens++}
        if curr == 20 {twentys++}
    }
    return true
}