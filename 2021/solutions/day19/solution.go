package day19

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
	"sync"
)

type VectorHash uint64
type Point struct{ X, Y, Z int }
type Vector = Point

func (p Point) VecTo(q Point) Vector {
	return Vector{q.X - p.X, q.Y - p.Y, q.Z - p.Z}
}

func (p1 Point) Substract(p2 Point) Point {
	return Point{
		X: p1.X - p2.X,
		Y: p1.Y - p2.Y,
		Z: p1.Z - p2.Z,
	}
}

func (p Point) GetOrientations() [48]Point {
	var res [48]Point
	x, y, z := p.X, p.Y, p.Z
	for i := 0; i < 8; i++ {
		j := i * 6
		res[j] = Point{x, y, z}
		res[j+1] = Point{x, z, y}
		res[j+2] = Point{y, x, z}
		res[j+3] = Point{y, z, x}
		res[j+4] = Point{z, x, y}
		res[j+5] = Point{z, y, x}
		if i&1 == 1 {
			for k := j; k < j+6; k++ {
				res[k].X = -res[k].X
			}
		}
		if i&2 == 2 {
			for k := j; k < j+6; k++ {
				res[k].Y = -res[k].Y
			}
		}
		if i&4 == 4 {
			for k := j; k < j+6; k++ {
				res[k].Z = -res[k].Z
			}
		}
	}
	return res
}

func (v Vector) Hash() VectorHash {
	return VectorHash((v.X + 2500) + 5000*(v.Y+2500) + 5000*5000*(v.Z+2500))
}

func parseVectors(points [][][48]Point) [][][48][]VectorHash {
	nscanner := len(points)
	vectors := make([][][48][]VectorHash, nscanner)
	for scanner, beaconPoints := range points {
		// Create map of vectors per beacon and orientation
		nbeacons := len(points[scanner])
		vectors[scanner] = make([][48][]VectorHash, nbeacons)

		// Initialize maps
		for beacon := 0; beacon < nbeacons; beacon++ {
			for orient := 0; orient < 48; orient++ {
				vectors[scanner][beacon][orient] = make([]VectorHash, 0, 27)
			}
		}

		// Calculate / add vectors for each pair of beacons
		for firstBeac := 0; firstBeac < nbeacons-1; firstBeac++ {
			for secondBeac := firstBeac + 1; secondBeac < nbeacons; secondBeac++ {
				// point values per orientation
				p1s, p2s := beaconPoints[firstBeac], beaconPoints[secondBeac]
				for orient := range p1s {
					p1ToP2 := p1s[orient].VecTo(p2s[orient])
					// vectors[scanner][firstBeac][orient][p1ToP2.hash()] = struct{}{}
					vectors[scanner][firstBeac][orient] = append(vectors[scanner][firstBeac][orient], p1ToP2.Hash())
					p2ToP1 := p2s[orient].VecTo(p1s[orient])
					// vectors[scanner][secondBeac][orient][p2ToP1.hash()] = struct{}{}
					vectors[scanner][secondBeac][orient] = append(vectors[scanner][secondBeac][orient], p2ToP1.Hash())
				}
			}
		}
		// Sort vectors for faster comparisons
		for beacon := 0; beacon < nbeacons; beacon++ {
			for orient := 0; orient < 48; orient++ {
				sort.Slice(vectors[scanner][beacon][orient], func(i, j int) bool {
					return vectors[scanner][beacon][orient][i] < vectors[scanner][beacon][orient][j]
				})
			}
		}
	}
	return vectors
}

func parsePoints(rows []string) [][][48]Point {
	var i, j int
	scannerBeaconPoints := make([][][48]Point, 0)
	for i < len(rows) {
		i++ // skip scanner id
		scannerBeaconPoints = append(scannerBeaconPoints, make([][48]Point, 0, 30))
		for ; i < len(rows) && rows[i] != ""; i++ {
			var x, y, z int
			fmt.Sscanf(rows[i], "%d,%d,%d", &x, &y, &z)
			p := Point{X: x, Y: y, Z: z}
			scannerBeaconPoints[j] = append(scannerBeaconPoints[j], p.GetOrientations())
		}
		i++
		j++
	}

	return scannerBeaconPoints
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func sharesSpace(v1, v2 []VectorHash) bool {
	var count int
	var l, r int
	for l != len(v1) && r != len(v2) && min(len(v1), len(v2))+count >= 11 {
		if v1[l] == v2[r] {
			count++
			if count == 11 {
				return true
			}
			l++
			r++
		} else {
			if v1[l] < v2[r] {
				l++
			} else if v2[r] < v1[l] {
				r++
			}
		}
	}
	return false
}

func compareScanners(
	vectors [][][48][]VectorHash,
	points [][][48]Point,
	scannerPos []Point,
	rootScanner, otherScanner int,
) bool {
	// For each beacon in first
	for rootBeacon := range vectors[rootScanner] {
		// If there exists a beacon + orientation in second such that there
		// are 11 shared vectors, then there is a match
		// Orientation of first doesn't matter, the second beacon is
		// exhaustively searched for all orientations
		firstVecs := vectors[rootScanner][rootBeacon][0]
		for otherBeacon := range vectors[otherScanner] {
			for orient := range vectors[otherScanner][otherBeacon] {
				if !sharesSpace(firstVecs, vectors[otherScanner][otherBeacon][orient]) {
					continue
				}
				// root and other scanner are within the same space. Shift the
				// orientation of the other scanner so that its aligned with root.

				// The root and other scanner are matching. The 'other' scanner
				// will now become a root for further iterations, so we adjust the
				// first orientation of each beacon so that it matches the root,
				// and also the position of each point as well. This will ensure
				// a shared field in the end.
				// Also adjust the point locations to align with the root
				p1 := points[rootScanner][rootBeacon][0]
				p2 := points[otherScanner][otherBeacon][orient]
				dx, dy, dz := p2.X-p1.X, p2.Y-p1.Y, p2.Z-p1.Z

				for otherBeacon := range vectors[otherScanner] {
					// Use the right orientation for vectors/points
					vectors[otherScanner][otherBeacon][0] = vectors[otherScanner][otherBeacon][orient]
					points[otherScanner][otherBeacon][0] = points[otherScanner][otherBeacon][orient]

					// Adjust locations
					p := points[otherScanner][otherBeacon][0]
					points[otherScanner][otherBeacon][0] = Point{
						X: p.X - dx,
						Y: p.Y - dy,
						Z: p.Z - dz,
					}
				}

				// Part 2: keep track of scanner positions
				scannerPos[otherScanner] = Point{X: -dx, Y: -dy, Z: -dz}
				return true
			}
		}
	}
	return false
}

func Solve() (part1, maxDist int) {
	inputScanner := bufio.NewScanner(strings.NewReader(Input))
	inputScanner.Split(bufio.ScanLines)

	var rows []string
	for inputScanner.Scan() {
		s := inputScanner.Text()
		rows = append(rows, s)
	}

	points := parsePoints(rows)
	vectors := parseVectors(points)
	nscanner := len(points)

	seen := make([]bool, nscanner)
	seenCount := 1
	seen[0] = true
	cur := []int{0}
	next := []int{}

	scannerPos := make([]Point, nscanner)
	scannerPos[0] = Point{0, 0, 0}

	var nextMtx sync.Mutex
	for seenCount != nscanner {
		next = next[:0]
		for _, rootScanner := range cur {
			var wg sync.WaitGroup
			for otherScanner := 0; otherScanner < nscanner; otherScanner++ {
				if seen[otherScanner] {
					continue
				}
				wg.Add(1)
				go func(other int) {
					defer wg.Done()
					if compareScanners(vectors, points, scannerPos, rootScanner, other) {
						seen[other] = true
						nextMtx.Lock()
						next = append(next, other)
						seenCount++
						nextMtx.Unlock()
					}
				}(otherScanner)
			}
			wg.Wait()
		}

		cur, next = next, cur
	}
	uniquePoints := make(map[Point]bool)
	for scanner := range points {
		for beacon := range points[scanner] {
			uniquePoints[points[scanner][beacon][0]] = true
		}
	}

	part1 = len(uniquePoints)

	for first := 0; first < nscanner-1; first++ {
		for second := 0; second < nscanner; second++ {
			p1, p2 := scannerPos[first], scannerPos[second]
			maxDist = max(maxDist, abs(p2.X-p1.X)+abs(p2.Y-p1.Y)+abs(p2.Z-p1.Z))
		}
	}

	return
}
