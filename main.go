package main

import (
	"net/http"

  "github.com/asaskevich/govalidator"
  "github.com/bismarkanes/web-go/application"
  "github.com/bismarkanes/web-go/infrastructure/database"
  "github.com/bismarkanes/web-go/infrastructure/persistence"
  "github.com/bismarkanes/web-go/infrastructure/utils"
  "github.com/bismarkanes/web-go/interfaces/handlers"
  "github.com/go-chi/chi"
  "github.com/go-chi/chi/middleware"
  "os"

  _ "github.com/joho/godotenv/autoload"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func main() {
	// initial database
	db := database.NewDBConnection()

  utils.Log("Db connection %s SUCCESS", db.Name())

  // migration
  err := database.Migrate(db)
  if err != nil {
    panic(err)
  }
  utils.Log("Success migration")

  // initial HTTP router handler
  r := chi.NewRouter()

  r.Use(middleware.Logger)

  r.Get("/ping", handlers.GetPing)

  // users
  userPersist := persistence.NewUsersRepo(db)
  userApp := application.NewUsers(userPersist)
  userHandler := handlers.NewUsersHandler(userApp)

  r.Route("/users", func(r chi.Router) {
    r.Get("/", userHandler.GetUser)
    r.Route("/{userID}", func(r chi.Router) {
      r.Get("/", nil)
      r.Post("/", nil)
      r.Patch("/", nil)
      r.Delete("/", nil)
    })
  })

  http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
