package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"POS/backend/api"
	"POS/backend/database/sqlite"
)

//go:embed all:frontend/dist
var frontendAssets embed.FS

type spaFS struct {
	sub fs.FS
}

func (s *spaFS) Open(name string) (fs.File, error) {
	f, err := s.sub.Open(name)
	if err != nil {
		return s.sub.Open("index.html")
	}
	return f, nil
}

func main() {
	sqlite.Init()
	api.SeedDefaultAdmin()

	apiRouter := api.NewRouter()

	var staticFS http.Handler
	if _, err := os.Stat("frontend/dist"); err == nil {
		sub, err := fs.Sub(frontendAssets, "frontend/dist")
		if err != nil {
			log.Fatal("Error al leer assets:", err)
		}
		staticFS = http.FileServer(http.FS(&spaFS{sub: sub}))
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/") {
			apiRouter.ServeHTTP(w, r)
			return
		}
		if staticFS != nil {
			staticFS.ServeHTTP(w, r)
		} else {
			apiRouter.ServeHTTP(w, r)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port
	log.Printf("Servidor iniciado en puerto %s", port)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatal("Error al iniciar servidor:", err)
	}
}
