// same intuition as decode strings, basic calculator
// solve / hash out the inner most and combine with parent
// inorder to access parent, push parent to stack when running into an open ( paran
// time: o(s) * o(s) for reversing
// space: o(s)
func reverseParentheses(s string) string {
    st := []*strings.Builder{}
    curr := new(strings.Builder)
    for i := 0; i < len(s); i++ {
        char := s[i]
        if char == '(' {
            // push parent to stack
            st = append(st, curr)
            curr = new(strings.Builder)
        } else if char == ')' {
            // reverse curr
            currRev := reverse(curr)
            // combine with parent
            if len(st) != 0 {
                parent := st[len(st)-1]
                st = st[:len(st)-1]
                parent.WriteString(currRev.String())
                curr = parent
            } else {
                curr = currRev
            }
            
        } else {
            curr.WriteByte(char)
        }
    }
    return curr.String()
}

func reverse(s *strings.Builder) *strings.Builder {
    newS := new(strings.Builder)
    sStr := s.String()
    for i := len(sStr)-1; i >= 0 ; i-- {
        newS.WriteByte(sStr[i])
    }
    return newS
}