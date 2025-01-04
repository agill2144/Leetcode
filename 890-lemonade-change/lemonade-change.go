func lemonadeChange(bills []int) bool {
    fives := 0
    tens := 0
    twenties := 0
    for i := 0; i < len(bills); i++ {
        if bills[i] == 5 {
            fives++
        } else if bills[i] == 10 {
            tens++
            if fives == 0 {return false}
            fives--
        } else if bills[i] == 20 {
            twenties++ 
            if tens >= 1 && fives >= 1 {
                tens--
                fives--
                continue
            }
            if fives >= 3 {
                fives -= 3
                continue
            }
            return false
        }
    }
    return true 
}