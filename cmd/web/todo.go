package web

import (
	"net/http"
)

func TodoFormHandler(w http.ResponseWriter, r *http.Request) {
	component := TodoForm()
	component.Render(r.Context(), w)
}

func TodoPostHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	task := r.FormValue("task")
	component := TodoPost(task)
	component.Render(r.Context(), w)
}
