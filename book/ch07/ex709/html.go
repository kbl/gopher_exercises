package ex709

import (
	"book/ch07/ex708"
	"book/ch07/sorting"
	"html/template"
)

var tracks = []*sorting.Track{
	{"Go3", "Moby", "Xoby", 1993, sorting.Length("3m31s")},
	{"Go3", "Moby", "Moby", 1992, sorting.Length("3m32s")},
	{"Go2", "Moby", "Moby", 1992, sorting.Length("3m33s")},
	{"Go2", "Xoby", "Moby", 1993, sorting.Length("3m34s")},
}

var TT = ex708.NewTrackTable(tracks)

var HTMLTemplate = template.Must(template.New("x").Parse(`<html>
<body>
  <table>
    <tr>
      <th><a href="/?s=Title">Title</a></th>
      <th><a href="/?s=Artist">Artist</a></th>
      <th><a href="/?s=Album">Album</a></th>
      <th><a href="/?s=Year">Year</a></th>
      <th><a href="/?s=Length">Length</a></th>
    </tr>
	{{ range .Tracks }}
	<tr>
      <td>{{.Title}}</td>
      <td>{{.Artist}}</td>
      <td>{{.Album}}</td>
      <td>{{.Year}}</td>
      <td>{{.Length}}</td>
	</tr>
	{{ end}}
  </table>
  </body>
</html>`))
