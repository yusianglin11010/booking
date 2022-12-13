package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func NoSurf(next http.Handler) http.Handler {
	fmt.Println("NoSurf used")
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler

}

func SessionLoad(next http.Handler) http.Handler {
	fmt.Println("Session used")

	return session.LoadAndSave(next)
}
