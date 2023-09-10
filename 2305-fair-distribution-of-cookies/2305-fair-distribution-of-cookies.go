func distributeCookies(cookies []int, k int) int {
    totalMax:=1000000 // keep current best result
    childs:=make([]int,k)

    var dfs func(ind int)int 
    dfs = func(ind int)int { //ind is index of cookies
        if len(cookies)==ind{ // all cookies are distributed, so find unfairness
            maximal:=0
            for i:=0;i<len(childs);i++{
                maximal = max(maximal,childs[i])
            }
            totalMax = min(totalMax,maximal) // update best result
            return maximal
        }

        minimal:=1000000
        cook:=cookies[ind]
        for i:=0;i<len(childs);i++{
            childs[i]+=cook
            if childs[i]<totalMax{ // skip variants that leads to worse result
                minimal = min(minimal,dfs(ind+1))
            }
            childs[i]-=cook
        }
        return minimal
    }

    dfs(0)
    return totalMax
}

func max( a,b int)int {
    if a<b {
        return b
    }
    return a
}

func min (a, b int)int{
    if a<b {
        return a
    }
    return b
}