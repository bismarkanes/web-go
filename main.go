package main

import (
	"net/http"

	"github.com/asaskevich/govalidator"
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

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("PONG"))
	})

	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
