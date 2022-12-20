func canVisitAllRooms(rooms [][]int) bool {
    st := []int{0}
    seen := make([]bool, len(rooms))
    seen[0] = true
    
    for len(st) != 0 {
        top := st[len(st)-1]; st = st[:len(st)-1]
        keysInTopRoom := rooms[top]
        for _, key := range keysInTopRoom {
            if ok := seen[key]; !ok {
                seen[key] = true
                st = append(st, key)
            }
        }
    }
    
    for _, val := range seen {
        if !val {return false}
    }
    
    return true
    
}