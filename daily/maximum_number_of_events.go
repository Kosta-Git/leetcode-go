package daily

import (
	"container/heap"
	"sort"
)

type Event struct {
	value    []int
	priority int
	index    int
}

type PriorityQueue []*Event

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Event)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Event, value []int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

type ScoreCompute interface {
	compute(event []int, k int) int
}

type PointAndTimeBasedV5ScoreCompute struct {
	Dates [][]int
}

func (p PointAndTimeBasedV5ScoreCompute) compute(event []int, k int) int {
	if k == 1 {
		return event[2]
	} else {
		p.Dates = append(p.Dates, event)
		overlap := !notOverlapping(event, p.Dates)
		var negative float64 = 0.1
		if overlap {
			negative = 0.5
		}

		eventTime := event[1] - event[0] + 1
		return int(float64(event[2])*negative*float64(len(p.Dates))) - eventTime
	}
}

type PointAndTimeBasedV4ScoreCompute struct {
	Dates [][]int
}

func (p PointAndTimeBasedV4ScoreCompute) compute(event []int, k int) int {
	if k == 1 {
		return event[2]
	} else {
		p.Dates = append(p.Dates, event)
		overlap := notOverlapping(event, p.Dates)
		var negative float64 = 1
		if overlap {
			negative = 0.5
		}

		eventTime := event[1] - event[0]
		if eventTime <= 0 {
			eventTime = 1
		}
		return int(float64(event[2])*negative) - eventTime
	}
}

type PointAndTimeBasedV3ScoreCompute struct{}

func (p PointAndTimeBasedV3ScoreCompute) compute(event []int, k int) int {
	if k == 1 {
		return event[2]
	} else {
		return event[2] * event[1] * event[0]
	}
}

type PointAndTimeBasedV2ScoreCompute struct{}

func (p PointAndTimeBasedV2ScoreCompute) compute(event []int, k int) int {
	if k == 1 {
		return event[2]
	} else {
		eventTime := event[1] - event[0]
		if eventTime <= 0 {
			eventTime = 1
		}
		return (event[2] / eventTime) * 1000
	}
}

type PointBasedScoreCompute struct{}

func (p PointBasedScoreCompute) compute(event []int, k int) int {
	if k == 1 {
		return event[2]
	} else {
		return event[2]
	}
}

type TimeBasedScoreCompute struct{}

func (p TimeBasedScoreCompute) compute(event []int, k int) int {
	if k == 1 {
		return event[2]
	} else {
		return event[1] - event[0]
	}
}

type PointAndTimeBasedScoreCompute struct{}

func (p PointAndTimeBasedScoreCompute) compute(event []int, k int) int {
	if k == 1 {
		return event[2]
	} else {
		eventTime := (event[1] - event[0]) + k
		if eventTime <= 0 {
			eventTime = 1
		}
		return ((event[2] / k) / eventTime) * 100
	}
}

// This solution doesn't work, it gives a good estimate but not the best answer for each case
func maxValue(events [][]int, k int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	computeTypes := []ScoreCompute{
		PointAndTimeBasedV2ScoreCompute{},
	}

	var pqArr []PriorityQueue
	for _, computeType := range computeTypes {
		pq := make(PriorityQueue, len(events))
		for i := 0; i < len(events); i++ {
			pq[i] = &Event{
				value:    events[i],
				priority: computeType.compute(events[i], k),
				index:    i,
			}
		}
		heap.Init(&pq)
		pqArr = append(pqArr, pq)
	}

	var eventArr [][][]int
	for _, pq := range pqArr {
		eventArr = append(eventArr, buildEvents(pq, k))
	}

	m := 0
	for _, events := range eventArr {
		m = Max(m, sum(events))
	}

	return m
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func buildEvents(queue PriorityQueue, k int) [][]int {
	currentEvents := make([][]int, 0)
	for queue.Len() > 0 && len(currentEvents) < k {
		e := heap.Pop(&queue).(*Event)
		if notOverlapping(e.value, currentEvents) {
			currentEvents = append(currentEvents, e.value)
		}
	}
	sort.Slice(currentEvents, func(i, j int) bool {
		return currentEvents[i][0] < currentEvents[j][0]
	})
	return currentEvents
}

func sum(events [][]int) int {
	sum := 0
	for i := 0; i < len(events); i++ {
		sum += events[i][2]
	}
	return sum
}

func notOverlapping(newEvent []int, currentEvents [][]int) bool {
	for _, event := range currentEvents {
		if newEvent[0] <= event[1] && newEvent[1] >= event[0] {
			return false
		}
	}
	return true
}
