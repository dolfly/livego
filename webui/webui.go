package webui

import (
	"embed"
	"net/http"
)

//go:embed build/*
var assets *embed.FS

func Assets() http.Handler {
	if assets != nil {
		return http.FileServer(http.FS(assets))
	}
	return http.NotFoundHandler()
}
