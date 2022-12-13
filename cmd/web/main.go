package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/yusianglin11010/booking/pkg/config"
	"github.com/yusianglin11010/booking/pkg/handler"
	"github.com/yusianglin11010/booking/pkg/render"
)

const port = ":8000"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false
	repo := handler.NewRepo(&app)

	handler.NewHandler(repo)

	render.NewTemplates(&app)
	// http.HandleFunc("/", handler.Repo.Home)
	// http.HandleFunc("/about", handler.Repo.About)
	fmt.Println(fmt.Sprintf("Starting the app on port %s", port))

	server := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = server.ListenAndServe()
	if err != nil {
		fmt.Println("listen failed", err)
	}
}
