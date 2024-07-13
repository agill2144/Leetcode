func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
    type robot struct {
        pos int
        dir int // +1 going right, -1 going left
        health int
        idx int
    }
    robots := []*robot{}
    for i := 0; i < len(positions); i++ {
        pos := positions[i]
        health := healths[i]
        dir := 1
        if directions[i] == 'L' {dir = -1}
        robots = append(robots, &robot{pos:pos,dir:dir,health:health, idx:i})
    }
    sort.Slice(robots, func(i, j int)bool{
        return robots[i].pos < robots[j].pos
    })
    st := []*robot{}
    for i := 0; i < len(robots); i++ {
        curr := robots[i]
        pushCurr := true
        for len(st) != 0 && st[len(st)-1].dir == 1 && curr.dir == -1 {
            // if both healths are same
            if st[len(st)-1].health == curr.health {
                st = st[:len(st)-1]
                pushCurr = false
                // both are destroyed
                break
            } else if st[len(st)-1].health > curr.health {
            // if top health is more
                // curr is destroyed
                st[len(st)-1].health--
                pushCurr = false
                break
            } else if curr.health > st[len(st)-1].health {
                // if curr health is more
                // then top is destroyed
                st = st[:len(st)-1]
                curr.health--
            }
        }
        if pushCurr {
            st = append(st, curr)
        }
    }
    sort.Slice(st, func(i, j int)bool{
        return st[i].idx < st[j].idx
    })
    out := []int{}
    for i := 0; i < len(st); i++ {
        out = append(out, st[i].health)
    }
    return out
}