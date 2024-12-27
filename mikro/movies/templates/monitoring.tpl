upstream metadata {
{{ range service "metadata" }}
  server {{ .Address }}:{{ .Port }}
{{ end }}
}

upstream movie {
{{ range service "movie" }}
  server {{ .Address }}:{{ .Port }}
{{ end }}
}

upstream rating {
{{ range service "rating" }}
  server {{ .Address }}:{{ .Port }}
{{ end }}
}

