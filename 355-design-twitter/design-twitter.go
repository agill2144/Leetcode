type tweet struct {
    id int
    time int
}

type Twitter struct {
    time int
    followings map[int]*set
    tweets map[int][]*tweet
}


func Constructor() Twitter {
    return Twitter{
        followings: map[int]*set{},
        tweets: map[int][]*tweet{},
    }
}


func (this *Twitter) PostTweet(userId int, tweetId int) {
    this.Follow(userId, userId)
    t := &tweet{tweetId, this.time}
    this.time++
    this.tweets[userId] = append(this.tweets[userId], t)
}


func (this *Twitter) GetNewsFeed(userId int) []int {
    mn := &minHeap{items: []*tweet{}}
    followings := this.followings[userId]
    if followings == nil {return nil}
    for id, _ := range followings.items {
        for _, tweet := range this.tweets[id] {
            heap.Push(mn, tweet)
            if mn.Len() > 10 {
                heap.Pop(mn)
            }
        }
    }
    out := []int{}
    for mn.Len() != 0 {
        out = append(out, heap.Pop(mn).(*tweet).id)
    }
    for i := 0; i < len(out)/2; i++ {
        out[i], out[len(out)-1-i] = out[len(out)-1-i],out[i]
    }
    return out
}


func (this *Twitter) Follow(followerId int, followeeId int)  {
    if this.followings[followerId] == nil {this.followings[followerId] = newSet()}
    this.followings[followerId].add(followeeId)
}


func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    if this.followings[followerId] == nil {return}
    this.followings[followerId].remove(followeeId)
}


type set struct {items map[int]bool}
func newSet() *set {return &set{items: map[int]bool{}}}
func (s *set) add(x int) {s.items[x] = true}
func (s *set) contains(x int) bool{
    _, ok := s.items[x]
    return ok
}
func (s *set) remove(x int) {delete(s.items, x)}
/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */

type minHeap struct {
	items []*tweet
}

func (m *minHeap) Less(i, j int) bool {
	return m.items[i].time < m.items[j].time
}
func (m *minHeap) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}
func (m *minHeap) Len() int {
	return len(m.items)
}
func (m *minHeap) Push(x interface{}) {
	m.items = append(m.items, x.(*tweet))
}
func (m *minHeap) Pop() interface{} {
	out := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return out
}
