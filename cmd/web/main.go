package main

import (
	"log"
	"net/http"
	"time"

	"bookings/pkg/config"
	"bookings/pkg/handlers"
	"bookings/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const numberPort = ":8080"
//Cria a variável do config
//Aqui nós vamos pegar as configurações que estão armazenadas dentro do Config.
var app config.AppConfig
var session *scs.SessionManager

func main() {

	

	//Quando for para ser o site real, mudar para true.
	app.InProduction = false

	//Vamos criar uma sessão com o aplicativo baixado.
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	//Atribuo a sessão criada acima para a sessão do config.
	app.Session = session


	//Aqui criamos uma variável de nome TC - Template Cache, que irá receber o conteúdo do arquivo render
	tc, err := render.CreateTemplateCache()
	
	//Verificamos se há erros ao criar a variável.
	if err != nil {
		log.Fatal("Não foi possível colocar a variável no cache")
	}

	//Aqui eu atribuo o TemplateChace do APP com o Tc criado agora.
	app.TemplateCache = tc
	render.NewTemplate(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	//	http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)
	//http.ListenAndServe(numberPort, nil)

	srv := &http.Server{
		Addr:    numberPort,
		Handler: routers(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
	

}
