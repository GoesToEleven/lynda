you could also do this:

type year struct {
	Fall, Spring, Summer semester
}

And then in your template, you could have things like this:


{{.Fall.Term}}<br>
{{range .Fall.Courses}}
	{{.Number}} - {{.Name}} - {{.Units}}<br>
{{end}}


{{.Spring.Term}}<br>
{{range .Spring.Courses}}
	{{.Number}} - {{.Name}} - {{.Units}}<br>
{{end}}