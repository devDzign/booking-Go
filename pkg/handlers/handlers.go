package handlers

import (
	"booking/pkg/config"
	"booking/pkg/models"
	"booking/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repositories
func NewRepo(a *config.AppConfig) *Repository {

	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, request *http.Request) {
	remoteIp := request.RemoteAddr
	m.App.Session.Put(request.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIp := m.App.Session.GetString(request.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{StringMap: stringMap})
}
