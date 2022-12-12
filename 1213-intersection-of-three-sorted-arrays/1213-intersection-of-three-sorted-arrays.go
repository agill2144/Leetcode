func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
    a1 := 0
    a2 := 0
    a3 := 0
    
    out := []int{}
    for a1 < len(arr1) && a2 < len(arr2) && a3 < len(arr3) {
        a1Val := arr1[a1]
        a2Val := arr2[a2]
        a3Val := arr3[a3]
        
        if (a1Val == a2Val && a2Val == a3Val) {
            out = append(out, a1Val)
            a1++; a2++; a3++
        } else if a1Val < a2Val  {
            a1++
        } else if a2Val < a3Val {
            a2++
        } else {
            a3++
        }
        
    }
    return out
    
}