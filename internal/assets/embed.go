package assets

import (
	"embed"
	"github.com/go-chi/chi/v5"
	"net/http"
)

//go:embed all:dist
var Assets embed.FS

// Mount To Serve Assets mount the embedded assets to an HTTP server
func Mount(r chi.Router) {
	r.Route("/dist", func(r chi.Router) {
		r.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				next.ServeHTTP(w, r)
			})
		})
		r.Handle("/*", http.FileServer(http.FS(Assets)))
	})
}
