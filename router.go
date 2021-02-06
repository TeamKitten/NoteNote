package main

import (
	"net/http"
	"strings"

	"github.com/TeamKitten/NoteNote/api/injector"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

type binaryFileSystem struct {
	fs http.FileSystem
}

func (b *binaryFileSystem) Open(name string) (http.File, error) {
	return b.fs.Open(name)
}

func (b *binaryFileSystem) Exists(prefix string, filepath string) bool {

	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
		if _, err := b.fs.Open(p); err != nil {
			return false
		}
		return true
	}
	return false
}

func BinaryFileSystem(root string) *binaryFileSystem {
	fs := &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: root}
	return &binaryFileSystem{
		fs,
	}
}

func startRouter() {
	router := gin.Default()
	router.Use(static.Serve("/", BinaryFileSystem("dist")))

	api := router.Group("/api")
	{
		h := injector.InjectNoteHandler()
		api.POST("/notes/:id", h.Set)
		api.GET("/notes/:id", h.Get)
	}

	router.Run()
}
