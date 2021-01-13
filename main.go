package main

import (
	"net/http"

  "os"

	"github.com/asaskevich/govalidator"
	"github.com/bismarkanes/web-go/application"
	"github.com/bismarkanes/web-go/infrastructure/database"
  "github.com/bismarkanes/web-go/infrastructure/utils"
  "github.com/bismarkanes/web-go/interfaces/handlers"
  "github.com/go-chi/chi"
  "github.com/go-chi/chi/middleware"
  _ "github.com/joho/godotenv/autoload"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func main() {
	// initial database
	db := database.NewDBConnection()

  utils.Log("Db connection %s SUCCESS", db.Name())

  err := database.Migrate(db)
  if err != nil {
    panic(err)
  }
  utils.Log("Success migration")

  // initial HTTP router handler
  initRouter()
}

func initRouter() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	ping := application.NewPing()
  routePing := handlers.NewPing(ping)

	r.Get("/ping", routePing.GetPing)

	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
