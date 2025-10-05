package day18

import (
	"container/heap"
	"math"
	"strings"

	"github.com/domszyn/adventofcode/2019/toolbox"
)

type SearchStateItem = toolbox.Item[SearchState]

func (g Grid) findShortestDistance(startState SearchState) int {
	gScore := make(map[SearchState]int)
	gScore[startState] = 0

	fScore := make(map[SearchState]int)
	fScore[startState] = g.h(startState)

	openSet := make(toolbox.PriorityQueue[SearchState], 0)
	openSetMap := make(map[SearchState]*SearchStateItem)

	startItem := &SearchStateItem{State: startState, Priority: fScore[startState]}
	heap.Push(&openSet, startItem)
	openSetMap[startState] = startItem

	for openSet.Len() > 0 {
		currentState := heap.Pop(&openSet).(*SearchStateItem).State
		delete(openSetMap, currentState)

		if currentState.HasAllKeys() {
			return gScore[currentState]
		}

		var possibleMoves []struct {
			lastKeyIndex int
			nextKey      rune
			distance     int
		}

		for lkIdx, lk := range currentState.LastKeys {
			currentPoint := g.getLocation(lk)

			for nextKey := 'a'; nextKey <= 'z'; nextKey++ {
				if currentState.HasKey(nextKey) {
					continue
				}

				nextKeyPoint := g.Keys[nextKey]
				pathInfo := g.Cache[currentPoint][nextKeyPoint]

				if pathInfo.Distance != math.MaxInt64 && (currentState.KeyMask&pathInfo.DoorsNeeded) == pathInfo.DoorsNeeded {
					possibleMoves = append(possibleMoves, struct {
						lastKeyIndex int
						nextKey      rune
						distance     int
					}{
						lkIdx, nextKey, pathInfo.Distance,
					})
				}
			}
		}

		for _, move := range possibleMoves {
			nextState := currentState

			if len(nextState.LastKeys) == 1 {
				nextState.LastKeys = string(move.nextKey)
			} else {
				var sb strings.Builder
				for i, c := range nextState.LastKeys {
					if i == move.lastKeyIndex {
						sb.WriteRune(move.nextKey)
					} else {
						sb.WriteRune(c)
					}
				}
				nextState.LastKeys = sb.String()
			}
			nextState.KeyMask = addKey(nextState.KeyMask, move.nextKey)

			tentativeGScore := gScore[currentState] + move.distance
			gScoreNeighbor, found := gScore[nextState]

			if !found || tentativeGScore < gScoreNeighbor {
				gScore[nextState] = tentativeGScore
				fScore[nextState] = tentativeGScore + g.h(nextState)

				if item, found := openSetMap[nextState]; found {
					openSet.Update(item, nextState, fScore[nextState])
				} else {
					item = &SearchStateItem{State: nextState, Priority: fScore[nextState]}
					heap.Push(&openSet, item)
					openSetMap[nextState] = item
				}
			}
		}
	}

	return -1
}
