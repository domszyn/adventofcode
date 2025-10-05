package day20

import (
	"container/heap"
	"math"

	"github.com/domszyn/adventofcode/2019/toolbox"
)

type State struct {
	CameFrom   Connection
	Connection Connection
	Level      int
}
type StateItem = toolbox.Item[State]

func (sm SpaceMap) findMinDistance() int {
	// for _, c1 := range sm.Connections {
	// 	for _, c2 := range sm.Connections {
	// 		if c1.Label == c2.Label && c1.Location != c2.Location {
	// 			sm.Distances[c1][c2] = 1
	// 		}
	// 	}
	// }

	distances := make(map[Connection]int)
	for _, c := range sm.Connections {
		distances[c] = math.MaxInt64
	}
	distances[sm.Start] = 0

	openSet := make(toolbox.PriorityQueue[State], 0)
	openSetMap := make(map[State]*StateItem)

	startState := State{Connection: sm.Start}
	startItem := &StateItem{State: startState, Priority: distances[sm.Start]}
	heap.Push(&openSet, startItem)
	openSetMap[startState] = startItem

	for openSet.Len() > 0 {
		current := heap.Pop(&openSet).(*StateItem).State
		currentLabel := current.Connection
		delete(openSetMap, current)

		if current.Connection.Label == sm.End.Label {
			return distances[current.Connection]
		}

		for nextConnection, dist := range sm.Distances[current.Connection] {
			if dist == 0 || dist == math.MaxInt64 {
				continue
			}

			newDist := distances[currentLabel] + dist

			if newDist < distances[nextConnection] {
				distances[nextConnection] = newDist

				nextState := State{Connection: nextConnection}
				if item, found := openSetMap[nextState]; found {
					openSet.Update(item, nextState, newDist)
				} else {
					item = &StateItem{State: nextState, Priority: newDist}
					heap.Push(&openSet, item)
					openSetMap[nextState] = item
				}
			}
		}
	}

	return -1
}

func (sm SpaceMap) findMinDistanceWithLevels() int {
	distances := make(map[State]int)

	openSet := make(toolbox.PriorityQueue[State], 0)
	openSetMap := make(map[State]*StateItem)

	startState := State{Connection: sm.Start}
	distances[startState] = 0
	startItem := &StateItem{State: startState, Priority: distances[startState]}
	heap.Push(&openSet, startItem)
	openSetMap[startState] = startItem

	for openSet.Len() > 0 {
		currentItem := heap.Pop(&openSet).(*StateItem)
		currentState := currentItem.State
		delete(openSetMap, currentItem.State)
		if distances[currentState] < currentItem.Priority {
			continue
		}

		if currentState.Connection.Label == sm.End.Label && currentState.Level == 0 {
			return distances[currentState]
		}

		for nextConnection, dist := range sm.Distances[currentState.Connection] {
			if dist == 0 {
				continue
			}

			if (nextConnection.Label == sm.Start.Label || nextConnection.Label == sm.End.Label) && currentState.Level > 0 {
				continue
			}

			newDist := distances[currentState] + dist
			nextState := State{
				Connection: nextConnection,
				Level:      currentState.Level,
				CameFrom:   currentState.Connection,
			}

			if dist, found := distances[nextState]; !found || newDist < dist {
				distances[nextState] = newDist

				if item, found := openSetMap[nextState]; found {
					openSet.Update(item, nextState, newDist)
				} else {
					item = &StateItem{State: nextState, Priority: newDist}
					heap.Push(&openSet, item)
					openSetMap[nextState] = item
				}
			}
		}

		if portal, found := sm.Portals[currentState.Connection]; found {
			nextLevel := currentState.Level + int(currentState.Connection.LevelChange)
			if nextLevel < 0 || nextLevel > len(sm.Connections) {
				continue
			}

			nextState := State{
				Connection: portal,
				Level:      nextLevel,
				CameFrom:   currentState.Connection,
			}
			newDist := distances[currentState] + 1

			if dist, found := distances[nextState]; !found || newDist < dist {
				distances[nextState] = newDist
				if item, found := openSetMap[nextState]; found {
					openSet.Update(item, nextState, newDist)
				} else {
					item = &StateItem{State: nextState, Priority: newDist}
					heap.Push(&openSet, item)
					openSetMap[nextState] = item
				}
			}
		}
	}

	return -1
}
