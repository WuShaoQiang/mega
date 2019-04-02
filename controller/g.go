package controller

import (
	"html/template"

	"github.com/gorilla/sessions"
)

var (
	homeController home
	templates      map[string]*template.Template
	sessionName    string
	store          *sessions.CookieStore
)

func init() {
	templates = populateTemplates()
	store = sessions.NewCookieStore([]byte("secret"))
	sessionName = "mega"
}

// Startup func
func Startup() {
	homeController.registerRoutes()
}
