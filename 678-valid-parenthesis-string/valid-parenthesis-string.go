func checkValidString(s string) bool {
    asteriks := []int{}
    opens := []int{}
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            opens = append(opens,i)
        } else if s[i] == '*' {
            asteriks = append(asteriks, i)
        } else {
            if len(opens) > 0 {
                opens = opens[:len(opens)-1]
            }else if len(asteriks) > 0{
                // use previously seen asterik as a open paran
                asteriks = asteriks[:len(asteriks)-1]
            } else {
                return false

            }
        }
    }
    for len(opens) != 0 && len(asteriks) != 0 {
        // we only have opens left, therefore we need to use asteriks as closing
        // therefore asteriks SHOULD HAVE CAME AFTER open paran
        if opens[len(opens)-1] < asteriks[len(asteriks)-1] {
            asteriks = asteriks[:len(asteriks)-1]
            opens = opens[:len(opens)-1]
            continue
        }
        break
    }
    return len(opens) == 0 
}