package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type VersionInfo struct {
	Version string
	Commit  string
	Build   string
}
type router struct {
	VersionInfo
	router *chi.Mux
}

func NewHandler(info VersionInfo) *chi.Mux {
	r := router{
		router:      chi.NewRouter(),
		VersionInfo: info,
	}

	r.router.Use(middleware.Recoverer)
	r.router.Get("/", r.handler)
	r.router.Get("/live", r.heartbeatHandler)
	r.router.Get("/version", r.versionHandler)
	return r.router
}

func (rt *router) heartbeatHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (rt *router) versionHandler(w http.ResponseWriter, r *http.Request) {
	version := map[string]string{
		"version": rt.VersionInfo.Version,
		"commit":  rt.VersionInfo.Commit,
		"build":   rt.VersionInfo.Build,
	}
	response, error := json.Marshal(version)
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write(response)

}

func (rt *router) handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
