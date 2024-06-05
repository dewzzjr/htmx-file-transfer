package handler

import (
	"net/http"

	"github.com/dewzzjr/htmx/component"
)

func AppPages(w http.ResponseWriter, r *http.Request) {
	props := component.AppProp{
		Sessions: []component.SessionProp{
			{Name: "Session 1"},
			{Name: "Session 2"},
		},
		Create: component.ModalProp{ID: "create", Get: "/modals/share"},
	}
	component.App(props).Render(r.Context(), w)
}
