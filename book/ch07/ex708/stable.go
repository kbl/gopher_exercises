package ex708

import (
	"bytes"
	"fmt"
	"github.com/kbl/gopher_exercises/book/ch07/sorting"
	"sort"
	"text/tabwriter"
)

type TrackTable struct {
	Tracks      []*sorting.Track
	sortHistory []string
}

func (t *TrackTable) Len() int {
	return len(t.Tracks)
}

func (t *TrackTable) Swap(i, j int) {
	t.Tracks[i], t.Tracks[j] = t.Tracks[j], t.Tracks[i]
}

func (t *TrackTable) Less(i, j int) bool {
	for sortI := len(t.sortHistory) - 1; sortI >= 0; sortI-- {
		column := t.sortHistory[sortI]
		switch column {
		case "Title":
			if t.Tracks[i].Title != t.Tracks[j].Title {
				return t.Tracks[i].Title < t.Tracks[j].Title
			}
		case "Artist":
			if t.Tracks[i].Artist != t.Tracks[j].Artist {
				return t.Tracks[i].Artist < t.Tracks[j].Artist
			}
		case "Album":
			if t.Tracks[i].Album != t.Tracks[j].Album {
				return t.Tracks[i].Album < t.Tracks[j].Album
			}
		case "Year":
			if t.Tracks[i].Year != t.Tracks[j].Year {
				return t.Tracks[i].Year < t.Tracks[j].Year
			}
		case "Length":
			if t.Tracks[i].Length != t.Tracks[j].Length {
				return t.Tracks[i].Length < t.Tracks[j].Length
			}
		}
	}
	return false
}

func (t *TrackTable) Sort(column string) {
	t.sortHistory = append(t.sortHistory, column)
	sort.Sort(t)
}

func NewTrackTable(tracks []*sorting.Track) *TrackTable {
	tt := TrackTable{tracks, nil}
	return &tt
}

func (t *TrackTable) String() string {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	var b bytes.Buffer

	tw := new(tabwriter.Writer).Init(&b, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, track := range t.Tracks {
		fmt.Fprintf(tw, format, track.Title, track.Artist, track.Album, track.Year, track.Length)
	}
	tw.Flush()

	return b.String()
}
