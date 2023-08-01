package queue

import (
    "container/list"
)

const (
    radiant uint8 = 82
    dire    uint8 = 68
)

// TODO: Fix this solution I have no clue why test 4 doesnt work

func predictPartyVictory(senate string) string {
    queue := createQueue(senate)
    solved, winner := isSolved(queue)
    for !solved {
        for e := queue.Front(); e != nil; e = e.Next() {
            if queue.Len() <= 1 {
                break
            }

            if e == queue.Back() {
                if e.Value != queue.Front().Value {
                    queue.Remove(queue.Front())
                    break
                }
            }

            itemToDelete := findNextDifferentElement(e)
            if itemToDelete == nil {
                queue.MoveToFront(e)
                break
            } else {
                queue.Remove(itemToDelete)
            }
        }
        solved, winner = isSolved(queue)
    }

    if winner == radiant {
        return "Radiant"
    } else {
        return "Dire"
    }
}

func createQueue(senate string) *list.List {
    queue := list.New()
    for i := 0; i < len(senate); i++ {
        queue.PushBack(senate[i])
    }
    return queue
}

func findNextDifferentElement(e *list.Element) *list.Element {
    for e.Next() != nil {
        if e.Next().Value != e.Value {
            return e.Next()
        }
        e = e.Next()
    }
    return nil
}

func isSolved(queue *list.List) (bool, uint8) {
    if queue.Len() == 0 {
        return false, 0
    }

    if queue.Len() == 1 {
        return true, queue.Front().Value.(uint8)
    }

    foundDire, foundRadiant := false, false
    for e := queue.Front(); e != nil; e = e.Next() {
        if e.Value.(uint8) == radiant {
            foundRadiant = true
        } else {
            foundDire = true
        }
    }

    if foundRadiant && foundDire {
        return false, 0
    }

    return true, queue.Front().Value.(uint8)
}
