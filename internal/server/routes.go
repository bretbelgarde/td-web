package server

import (
	"net/http"
	"td-web/cmd/web"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/js/", http.FileServer(http.FS(web.JS)))
	mux.Handle("/css/", http.FileServer(http.FS(web.CSS)))
	mux.HandleFunc("GET /{$}", web.TodoFormHandler)
	mux.HandleFunc("POST /todo", web.TodoPostHandler)

	return mux
}
