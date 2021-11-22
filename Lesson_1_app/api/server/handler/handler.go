package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type VersionInfo struct {
	Name      string
	Version   string
	GoVersion string
	BuildDate string
	GitLog    string
	GitHash   string
	GitBranch string
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
		"Name":      rt.VersionInfo.Name,
		"Version":   rt.VersionInfo.Version,
		"GoVersion": rt.VersionInfo.GoVersion,
		"BuildDate": rt.VersionInfo.BuildDate,
		"GitLog":    rt.VersionInfo.GitLog,
		"GitHash":   rt.VersionInfo.GitHash,
		"GitBranch": rt.VersionInfo.GitBranch,
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
