
type activeEvent struct {
    startStation string
    startTime int
}

type UndergroundSystem struct {
    activeIds map[int]*activeEvent
    avgTravelTimes map[string]map[string][]int
}


func Constructor() UndergroundSystem {
    return UndergroundSystem{
        activeIds: map[int]*activeEvent{},
        avgTravelTimes: map[string]map[string][]int{},
    }
}


func (this *UndergroundSystem) CheckIn(id int, stationName string, t int)  {
    this.activeIds[id] = &activeEvent{startStation:stationName, startTime:t}
}


func (this *UndergroundSystem) CheckOut(id int, stationName string, t int)  {
    event := this.activeIds[id]
    if event == nil {return}
    delete(this.activeIds, id)
    startS := event.startStation
    startT := event.startTime
    endS := stationName
    endT := t
    travelTime := endT-startT
    if this.avgTravelTimes[startS] == nil {
        this.avgTravelTimes[startS] = map[string][]int{}
    }
    if this.avgTravelTimes[startS][endS] == nil {
        this.avgTravelTimes[startS][endS] = make([]int, 2)
    }
    this.avgTravelTimes[startS][endS][0] += travelTime
    this.avgTravelTimes[startS][endS][1]++
}


func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
    totalTimes := this.avgTravelTimes[startStation][endStation][0]
    count := this.avgTravelTimes[startStation][endStation][1]
    return float64(totalTimes)/float64(count)
}


/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */