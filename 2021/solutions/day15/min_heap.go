package day15

type MinHeap struct {
	fnScore  func(l Location) int
	points   []Location
	minScore int
	minIdx   int
}

func (h *MinHeap) Add(loc Location, score func(l Location) int) {
	h.points = append(h.points, loc)
	if s := score(loc); s < h.minScore {
		h.minScore = s
		h.minIdx = len(h.points) - 1
	}
}

func (h *MinHeap) ExtractMin() Location {
	minPoint := h.points[h.minIdx]

	tmpPoints := h.points[:h.minIdx]
	tmpPoints = append(tmpPoints, h.points[h.minIdx+1:]...)

	return minPoint
}
