package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nathanaelcunningham/gothBase/internal/env"
	"github.com/nathanaelcunningham/gothBase/public"
	"github.com/nathanaelcunningham/gothBase/views/home"
	"github.com/nathanaelcunningham/gothBase/views/layouts"
)

func Routes() http.Handler {
	isDev := env.GetBool("DEV", true)
	r := chi.NewRouter()

	if isDev {
		r.Use(public.Reload)
	}

	r.Handle("/*", public.Public())
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		layouts.Base().Render(r.Context(), w)
	})
	r.Get("/foo", func(w http.ResponseWriter, r *http.Request) {
		home.Home().Render(r.Context(), w)
	})
	return r
}
