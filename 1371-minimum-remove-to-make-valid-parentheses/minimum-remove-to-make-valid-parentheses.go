type stNode struct {
    val byte
    idx int
}

func minRemoveToMakeValid(s string) string {
    // keep track of which indicies are invalid
    invalidIdxSet := map[int]struct{}{}
    st := []stNode{}
    validOpenings := map[byte]byte{
        ')' : '(',
        '}' : '{',
        ']' : '[',
    }

    for i := 0; i < len(s); i++ {
        char := s[i]
        if char == '[' || char == '(' || char == '{' {
            st = append(st, stNode{char,i})
        } else if char == ']' || char == ')' || char == '}' {
            
            // when is a closing paran invalid? 

            // 1. when there is no corresponding open paran
            if len(st) == 0 {
                // invalid closing paran
                invalidIdxSet[i] = struct{}{}
                continue
            }

            // 2. last open paran is not the correct corresponding paran
            top := st[len(st)-1]
            lastOpen := top.val
            validOpen := validOpenings[char]
            if lastOpen != validOpen {
                invalidIdxSet[i] = struct{}{}
                continue
            }

            st = st[:len(st)-1]
        }
    }

    // or 3. there are open parans in the st have never closed
    for len(st) != 0 {
        top := st[len(st)-1]
        st = st[:len(st)-1]
        invalidIdxSet[top.idx] = struct{}{}
    }

    out := new(strings.Builder)
    for i := 0; i < len(s); i++ {
        _, ok := invalidIdxSet[i]
        if ok {continue}
        out.WriteByte(s[i])
    }
    return out.String()
}