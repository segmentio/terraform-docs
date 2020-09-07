## Changelog

{{ range .Versions }}
{{ if .NoteGroups }}
{{ range .NoteGroups -}}
### {{ .Title }}
{{ range .Notes -}}
{{ .Body }}
{{ end -}}
{{ end -}}
{{ end -}}
{{- if .CommitGroups -}}
{{ range .CommitGroups }}
### {{ .Title }}
{{ range .Commits -}}
{{ if not (contains .Subject "<code>") -}}
- {{ if .Scope }}**{{ .Scope }}:** {{ end }}{{ if .Subject }}{{ .Subject }}{{ else }}{{ .Header }}{{ end }}
{{ end -}}
{{ end -}}
{{ end }}
{{ else }}
{{ range .Commits -}}
- {{ if .Scope }}**{{ .Scope }}:** {{ end }}{{ if .Subject }}{{ .Subject }}{{ else }}{{ .Header }}{{ end }}
{{ end }}
{{ end -}}
{{ end -}}
