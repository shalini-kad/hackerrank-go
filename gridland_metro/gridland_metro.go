package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

/* naive solution: does not do well with large numbers
func main() {
	var rows, cols, numTracks int
	fmt.Scanf("%d %d %d", &rows, &cols, &numTracks)

	gridland := make([][]int, rows)
	for i := 0; i < rows; i++ {
		gridland[i] = make([]int, cols)
	}

	for i := 0; i < numTracks; i++ {
		var r, c1, c2 int
		fmt.Scanf("%d %d %d", &r, &c1, &c2)
		for j := c1; j <= c2; j++ {
			gridland[r-1][j-1] = 1
		}
	}

	numSpotsOpen := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if gridland[i][j] != 1 {
				numSpotsOpen++
			}
		}
	}

	fmt.Println(numSpotsOpen)
}
*/

type trackRange struct {
	start int
	end   int
}

func main() {
	// use bufio for buffered io, faster than fmt.Scan
	in := bufio.NewReader(os.Stdin)

	var rows, cols, numTracks int
	fmt.Fscan(in, &rows, &cols, &numTracks)

	gridland := make(map[int]*list.List)

	for i := 0; i < numTracks; i++ {
		var (
			row      int
			newTrack trackRange
		)

		fmt.Fscan(in, &row, &newTrack.start, &newTrack.end)
		row -= 1
		newTrack.start -= 1
		newTrack.end -= 1

		if gridland[row] == nil {
			gridland[row] = list.New()
		}

		mergeTrack(gridland[row], newTrack)
	}

	// count empty spaces
	numSpaces := 0
	for _, trackList := range gridland {
		rowSpaces := cols
		for e := trackList.Front(); e != nil; e = e.Next() {
			track := e.Value.(trackRange)
			trackLength := track.end - track.start + 1
			rowSpaces -= trackLength
		}
		numSpaces += rowSpaces
	}

	// add spaces for blank lines
	numSpaces += (cols * (rows - len(gridland)))

	fmt.Println(numSpaces)
}

func mergeTrack(tracks *list.List, newTrack trackRange) {
	var e, next *list.Element

	for e = tracks.Front(); e != nil; e = next {
		track := e.Value.(trackRange)

		if newTrack.start < track.start {

			if newTrack.end < track.start {
				// starts and ends before - insert before
				tracks.InsertBefore(newTrack, e)
				break
			} else if newTrack.end <= track.end {
				// starts before, but ends inside - merge
				track.start = newTrack.start
				break
			} else {
				// starts before and ends after - supercede existing
				next = e.Next()
				tracks.Remove(e)
			}

		} else if newTrack.start == track.start {

			if newTrack.end <= track.end {
				// encompassed by existing - noop
				break
			} else {
				// ends after - supercede existing
				next = e.Next()
				tracks.Remove(e)
			}

		} else {

			if newTrack.end <= track.end {
				// encompassed by existing - noop
				break
			} else if newTrack.start <= track.end {
				// starts inside and ends after - merge and continue
				newTrack.start = track.start
				next = e.Next()
				tracks.Remove(e)
			} else {
				// after current track - continue
				next = e.Next()
			}
		}
	}

	if e == nil {
		tracks.PushBack(newTrack)
	}
}
