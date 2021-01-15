package main

import (
	"net/http"

    "github.com/asaskevich/govalidator"
    "github.com/bismarkanes/web-go/application"
    "github.com/bismarkanes/web-go/infrastructure/config"
    "github.com/bismarkanes/web-go/infrastructure/database"
    "github.com/bismarkanes/web-go/infrastructure/persistence"
    "github.com/bismarkanes/web-go/infrastructure/utilities"
    "github.com/bismarkanes/web-go/interfaces/handlers"
    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"

    _ "github.com/joho/godotenv/autoload"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func main() {
    conf := config.NewConfig()

	// initial database
	db := database.NewDBConnection()

    utilities.Log("Db connection %s SUCCESS", db.Name())

    // migration
    err := database.Migrate(db)
    if err != nil {
        panic(err)
    }
    utilities.Log("Success migration")

    // initial HTTP router handler
    r := chi.NewRouter()

    r.Use(middleware.Logger)

    r.Get("/ping", handlers.GetPing)

    // users
    userPersist := persistence.NewUserRepo(db)
    userApp := application.NewUser(userPersist)
    userHandler := handlers.NewUserHandler(userApp)

    r.Route("/users", func(r chi.Router) {
        r.Get("/", userHandler.GetAllUser)
        r.Post("/", userHandler.CreateUser)
        r.Route("/{userID}", func(r chi.Router) {
            r.Get("/", userHandler.GetUser)
            r.Patch("/", userHandler.UpdateUser)
            r.Delete("/", userHandler.DeleteUser)
        })
    })

    http.ListenAndServe(":"+conf.AppPort, r)
}
