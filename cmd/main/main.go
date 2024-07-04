package main

import (
    "log"
    "github.com/joho/godotenv"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/SoroushBeigi/library-backend-go.git/pkg/routes"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  r:=mux.NewRouter()
  routes.RegisterLibraryRoutes(r)
  http.Handle("/",r)
  log.Fatal(http.ListenAndServe("localhost:9010",r))
}