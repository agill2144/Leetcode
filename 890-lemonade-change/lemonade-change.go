func lemonadeChange(bills []int) bool {
    fives := 0
    tens := 0
    for i := 0; i < len(bills); i++ {
        currBill := bills[i]
        if currBill == 5 {
            fives++
        } else if currBill == 10 {
            if fives == 0 {return false}
            tens++
            fives--
        } else if currBill == 20 {
            if tens >= 1 && fives >= 1 {
                tens--
                fives--
            } else if fives >= 3 {
                fives -= 3
            } else {
                return false
            }
        }
    }
    return true
}