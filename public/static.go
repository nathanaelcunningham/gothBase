package public

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/aarol/reload"
	"github.com/nathanaelcunningham/gothBase/internal/env"
)

//go:embed css/* js/*
var publicFS embed.FS

func Public() http.Handler {
	isDev := env.GetBool("DEV", true)

	if isDev {
		return http.StripPrefix("/public/", http.FileServer(http.Dir("public")))
	}
	contentStatic, err := fs.Sub(publicFS, ".")
	if err != nil {
		panic(err)
	}
	return http.StripPrefix("/public/", http.FileServer(http.FS(contentStatic)))
}

func Reload(handler http.Handler) http.Handler {
	reloader := reload.New("public/js/")
	return reloader.Handle(handler)
}
