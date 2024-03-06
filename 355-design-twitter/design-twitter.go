type Twitter struct {
    tweets map[int][][]int // userId: [[tweetId, time]]
    followers map[int]*set // userId: {userId}
    time int

}


func Constructor() Twitter {
    return Twitter{
        tweets: map[int][][]int{},
        followers: map[int]*set{},
    }
}


func (this *Twitter) PostTweet(userId int, tweetId int)  {
    this.Follow(userId, userId)
    this.tweets[userId] = append(this.tweets[userId], []int{tweetId, this.time})
    this.time++
}


func (this *Twitter) GetNewsFeed(userId int) []int {
    if this.followers[userId] == nil {return nil}
    mn := &minheap{items: [][]int{}}
    for userId, _ := range this.followers[userId].items {
        tweetsByUserId := this.tweets[userId]
        for i := len(tweetsByUserId)-1; i >= 0; i-- {
            heap.Push(mn, tweetsByUserId[i])
            if mn.Len() > 10 {heap.Pop(mn)}
        }
    }
    out := make([]int, mn.Len())
    for i := mn.Len()-1; i >= 0; i-- {
        out[i] = heap.Pop(mn).([]int)[0]
    }
    return out
}


func (this *Twitter) Follow(followerId int, followeeId int)  {
    if this.followers[followerId] == nil {
        this.followers[followerId] = newSet()
    }
    this.followers[followerId].add(followeeId)
}


func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    if this.followers[followerId] == nil {return}
    this.followers[followerId].remove(followeeId)
}


type set struct {
    items map[int]struct{}
}

func newSet() *set {
    return &set{items:map[int]struct{}{}}
}
func (s *set) contains(item int) bool {
    _, ok := s.items[item]
    return ok
}
func (s *set) add(item int) {
    s.items[item] = struct{}{}
}
func (s *set) remove(item int) {
    delete(s.items, item)
}

// heap implementation ( sort by tweet.time )
// we will use min heap
type minheap struct {
    items [][]int
}
func (m *minheap) Len() int {return len(m.items)}
func (m *minheap) Less(i, j int) bool {return m.items[i][1] < m.items[j][1]}
func (m *minheap) Swap(i, j int) { m.items[i],m.items[j] = m.items[j], m.items[i]}
func (m *minheap) Push(x interface{}) {m.items = append(m.items, x.([]int))}
func (m *minheap) Pop() interface{} {
	i := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return i
}




/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */