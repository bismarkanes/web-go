package main

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/bismarkanes/web-go/application"
	"github.com/bismarkanes/web-go/interfaces/routes"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func main() {
	setupRouter()
}

func setupRouter() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	ping := application.NewPing()
	routePing := routes.NewPing(ping)
	r.Get("/ping", routePing.GetPing)

	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
