package main

import (
	"booking/pkg/config"
	"booking/pkg/handlers"
	"booking/pkg/render"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"
const environment = "dev"

var app config.AppConfig
var session *scs.SessionManager

// main is the main function
func main() {
	// change this to true when in production
	app.InProduction = false
	app.Env = environment

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Server start in : http://localhost%s", portNumber))
	fmt.Println(fmt.Sprintf("Used Cache template : %v", app.UseCache))
	fmt.Println(fmt.Sprintf("Envirenement : %s", app.Env))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
