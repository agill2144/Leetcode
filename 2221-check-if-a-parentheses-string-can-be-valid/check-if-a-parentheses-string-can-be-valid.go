func canBeValid(s string, locked string) bool {
    open := 0
    close := 0
    n := len(s)
    for i := 0; i < n; i++ {
        if locked[i] == '0' || s[i] == '(' {
            open++
        } else if s[i] == ')' {
            open--
            if open < 0 {return false}
        }

        if locked[n-1-i] == '0' || s[n-1-i] == ')' {
            close++
        } else if s[n-1-i] == '(' {
            close--
            if close < 0 {return false}
        }
    }

    return open % 2 == 0 && close % 2 == 0
}

/*
The stack intuition is same as valid parans with asteriks, when traversing the first time

if we have a closing paran; do we have any open? do we have asteriks seen before?
similarly
if we have a closing paran; do we have any open? do we have any unlocks seen before?
when traversing the second time, when we have open parans left

we could only close them if the asteriks idxs appeared after open paran idxs
similary
we can only close open parans, if unlock idxs appeared after open paran idxs
at the end, we must have closed all open parans ( i.e len of open paran stack should be 0 )
but its possible that we still have left over unlock idxs.
In the asteriks problem, we could ignore extra asteriks because we were allowed to discard them
but in this question, we cannot discard them.
Therefore if number of remaining unlocked idxs are even in size, we can assume that we can do some flips and make them valid, but if we had odd number of unlock idxs, its impossible to use-them-all and make them-all-valid.


*/
// func canBeValid(s string, locked string) bool {
//     if len(s) % 2 != 0 {return false}
//     flexible := []int{}
//     open := []int{}
//     for i := 0; i < len(s); i++ {
//         if locked[i] == '0' {
//             flexible = append(flexible, i)
//         } else if s[i] == '(' {
//             open = append(open, i)
//         } else if s[i] == ')' {
//             if len(open) > 0 {
//                 open = open[:len(open)-1]
//             } else if len(flexible) > 0 {
//                 flexible = flexible[:len(flexible)-1]
//             } else {
//                 return false
//             }
//         }
//     }
//     for len(open) > 0 && len(flexible) > 0 && flexible[len(flexible)-1] > open[len(open)-1] {
//         flexible = flexible[:len(flexible)-1]
//         open = open[:len(open)-1]
//     }
//     return len(open) == 0 && len(flexible) % 2 == 0
// }