// @title Notes API
// @version 1.0
// @description REST API for notes
// @BasePath /api/v1
package main

import (
	_ "example.com/pr12/docs"
	"example.com/pr12/internal/httpapi"
	"example.com/pr12/internal/repo"
	"example.com/pr12/internal/service"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	repository := repo.NewMemory()
	h := &httpapi.Handler{Service: service.New(repository)}
	m := h.Router()
	m.Handle("GET /docs/", httpSwagger.WrapHandler)
	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	log.Fatal((&http.Server{Addr: addr, Handler: m, ReadHeaderTimeout: 5 * time.Second}).ListenAndServe())
}
