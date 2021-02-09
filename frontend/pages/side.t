{{ define "side"}}


{{range $item := .Notes}}
<li>{{ $item }}</li>
{{end}}


{{ end }}