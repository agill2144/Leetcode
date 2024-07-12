func maximumGain(s string, x int, y int) int {
    higher := "ab";hPoints := x
    lower := "ba";lPoints := y
    if y > x {
        higher = "ba";hPoints = y
        lower = "ab";lPoints = x
    }
    // first pass only for higher score
    st := []byte{}
    points := 0
    for i := 0; i < len(s); i++ {
        st = append(st, s[i])
        for len(st) >= 2 {
            top := st[len(st)-1]
            prev := st[len(st)-2]
            str := fmt.Sprintf("%v%v", string(prev), string(top))

            if str == higher {
                st = st[:len(st)-2]
                points += hPoints
                continue
            }
            break
        }

    }
    st2 := []byte{}
    for i := 0; i < len(st); i++ {
        curr := st[i]
        st2 = append(st2, curr)
        for len(st2) >= 2 {
            top := st2[len(st2)-1]
            prev := st2[len(st2)-2]
            str := fmt.Sprintf("%v%v", string(prev), string(top))
            if str == lower {
                points += lPoints
                st2 = st2[:len(st2)-2]
                continue
            }
            break
        }
    }
    return points
}