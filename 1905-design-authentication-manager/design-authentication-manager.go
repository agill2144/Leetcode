type AuthenticationManager struct {
    leases map[string]int
    ttl int
}


func Constructor(timeToLive int) AuthenticationManager {
    return AuthenticationManager{
        leases: map[string]int{},
        ttl: timeToLive,
    }
}


func (this *AuthenticationManager) Generate(tokenId string, currentTime int)  {
    _, ok := this.leases[tokenId]
    if !ok {
        this.leases[tokenId] = currentTime+this.ttl
    }
}


func (this *AuthenticationManager) Renew(tokenId string, currentTime int)  {
    expiryTime, ok := this.leases[tokenId]
    isExpired := expiryTime <= currentTime
    if !ok || isExpired {return}
    this.leases[tokenId] = currentTime+this.ttl
}


func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
    count := 0
    for _, v := range this.leases {
        if currentTime < v {count++}
    }
    return count
}


/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */