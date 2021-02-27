package ex712

import (
	"github.com/kbl/gopher_exercises/book/ch07/ex711"
	"html/template"
	"net/http"
)

var htmlTemplate = template.Must(template.New("x").Parse(`<html>
<body>
  <table>
    <tr>
      <th>Name</th>
      <th>Price</th>
    </tr>
	{{ range $key, $value := . }}
	<tr>
      <td>{{$key}}</td>
      <td>{{$value}}</td>
	</tr>
	{{ end}}
  </table>
  </body>
</html>`))

type Database struct {
	ex711.Database
}

func (db Database) List(w http.ResponseWriter, req *http.Request) {
	htmlTemplate.Execute(w, db.Database.Database)
}
