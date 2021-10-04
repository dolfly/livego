package webui

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed build/*
var assets embed.FS

func Assets() http.FileSystem {
	fsys, _ := fs.Sub(assets, "build")
	return http.FS(fsys)
}
