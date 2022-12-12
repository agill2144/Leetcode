/*
matrix traversals are all about watching for boundaries!

We are going in either or directions
1. up
2. down

when going up, what boundaries do we hit when we are going out of bounds?

    We will hit
        top row or
     _ _ _ _ _ 
    |         |
    |         |
    |         |  right col
    |         |
    |_ _ _ _ _|

So if going up ever reaches row == 0 ( top row )
then go to next col ( c++ ) and change direction and go down
if going up every ever hits right col, then row++ and change direction to go down


when going down, what boundaries do we hit when we are going out of bounds?
         _ _ _ _ _ 
        |         |
        |         |
left col|         |
        |         |
        |_ _ _ _ _|
            bottom
            row
            

if going down we ever hit col 0, then row++ and change direction to go up
if going down we ever hit bottom last row ( m-1 ), then c++ and change direction


*/

func findDiagonalOrder(mat [][]int) []int {
    if mat == nil {
        return nil
    }
    up := true
    r := 0
    c := 0
    m := len(mat)
    n := len(mat[0])
    idx := 0
    out := []int{}
    
    for idx < m*n {
        idx++
        out = append(out, mat[r][c])
        if up {
            if c == n-1 {
                up = false
                r++
            } else if r == 0 {
                up = false
                c++
            } else {
                r--
                c++
            }
        } else {
            if r == m-1 {
                c++
                up = true
            } else if c == 0 {
                up = true
                r++
            } else {
                r++
                c--
            }
        }
    }
    return out
}