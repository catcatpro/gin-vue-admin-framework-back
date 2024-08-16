package cmd

import (
	"gin_vue_admin_framework/internal/routes"
	"log"
	"net/http"
)

func init() {

}

func startWebServer() {

	r := routes.Router

	webServer := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	if err := webServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("webserver listen error: %s\n", err)
	}

}

func RunServer() {
	InitSystem()
	startWebServer()
}
