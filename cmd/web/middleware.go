package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// Aqui vai ser uma função que irá fazer escrever algo entre uma requisição de uma página e outra.
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)

	})
}

//NoSurf adiciona proteção nas requisições POST
func NoSurf(next http.Handler) http.Handler {
	csrHandler := nosurf.New(next)

	csrHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrHandler

}

//SessionLoad carrega as sessões em cada requisição.
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
