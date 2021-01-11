package main

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/bismarkanes/web-go/application"
	"github.com/bismarkanes/web-go/infrastructure/database"
	"github.com/bismarkanes/web-go/interfaces/routes"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func main() {
	// initial database
	db := database.NewDBConnection()
	log.Printf("Db connection %s SUCCESS", db.Name())

	// initial HTTP router handler
	initRouter()
}

func initRouter() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	ping := application.NewPing()
	routePing := routes.NewPing(ping)
	r.Get("/ping", routePing.GetPing)

	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
