package handler

import (
	"net/http"

	"github.com/sabatagung/bookings/pkg/config"
	"github.com/sabatagung/bookings/pkg/models"
	"github.com/sabatagung/bookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// repository type
type Repository struct {
	App *config.AppConfig
}

// create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}
func (h *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	h.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTmpl(w, "home.page.html", &models.TemplateData{})
}

func (h *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic and send data
	stringMap := make(map[string]string)
	stringMap["test"] = "anyeonghaseo!"

	remoteIP := h.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	render.RenderTmpl(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
