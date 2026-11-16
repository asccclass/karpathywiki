// router.go
package main

import (
	"net/http"

	SherryServer "github.com/asccclass/sherryserver"
)

func NewRouter(srv *SherryServer.Server, documentRoot string) *http.ServeMux {
	router := http.NewServeMux()

	// Static assets
	staticfileserver := SherryServer.StaticFileServer{StaticPath: documentRoot}
	staticfileserver.AddRouter(router)

	/*
	   // API routes
	   router.HandleFunc("GET /api/notes", GetAll)
	   router.HandleFunc("POST /api/notes", Post)

	   // Go template + HTMX pages
	   router.HandleFunc("GET /", HomeTemplate)
	   router.HandleFunc("GET /partials/page/{page}", PagePartial)
	*/
	return router
}
