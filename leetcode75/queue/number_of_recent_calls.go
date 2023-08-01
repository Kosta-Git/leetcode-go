package queue

type RecentCounter struct {
    ping         int
    previousPing *RecentCounter
}

func Constructor() RecentCounter {
    return RecentCounter{
        ping:         0,
        previousPing: nil,
    }
}

func (r *RecentCounter) Ping(t int) int {
    if r.previousPing == nil {
        r.previousPing = &RecentCounter{
            ping:         t,
            previousPing: nil,
        }
        return 1
    }

    r.previousPing = &RecentCounter{
        ping:         t,
        previousPing: r.previousPing,
    }

    return r.deleteOldPings(t)
}

func (r *RecentCounter) deleteOldPings(currentTime int) int {
    if r.previousPing == nil {
        return 0
    }

    if r.previousPing.ping+3000 < currentTime {
        r.previousPing = nil
        return 0
    }

    return 1 + r.previousPing.deleteOldPings(currentTime)
}
