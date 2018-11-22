package ex708

import (
	"book/ch07/sorting"
	"testing"
)

var tracks = []*sorting.Track{
	{"Go3", "Moby", "Xoby", 1992, sorting.Length("3m37s")},
	{"Go3", "Moby", "Moby", 1992, sorting.Length("3m37s")},
	{"Go2", "Moby", "Moby", 1992, sorting.Length("3m37s")},
	{"Go2", "Xoby", "Moby", 1992, sorting.Length("3m37s")},
}

func TestStableSort(t *testing.T) {
	tt := NewTrackTable(tracks)

	t0 := tracks[0]
	t1 := tracks[1]
	t2 := tracks[2]
	t3 := tracks[3]

	tt.Sort("Album")
	tt.Sort("Artist")
	tt.Sort("Title")
	tt.Sort("Year")

	sortedTracks := []*sorting.Track{t2, t3, t1, t0}

	for i, track := range tracks {
		if track != sortedTracks[i] {
			t.Errorf("%v != %v at index %d", track, sortedTracks[i], i)
		}
	}
}
