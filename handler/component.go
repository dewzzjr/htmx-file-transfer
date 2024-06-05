package handler

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/dewzzjr/htmx/component"
	"github.com/gorilla/mux"
)

var modals = map[string]func() templ.Component{
	"share":   component.QuickShare,
	"session": component.CreateSession,
}

func Modals(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	compo, ok := modals[name]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	compo().Render(r.Context(), w)
}
