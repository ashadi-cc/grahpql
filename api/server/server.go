package server

import (
	"gql-ashadi/api/router"
	"gql-ashadi/datastore/pg"
	"gql-ashadi/service/config"
	"gql-ashadi/service/logger"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func InitServer() {
	logger.GetLogger().Info("setup database connection")
	//setup database connection
	pg.MustConnect()

	//setup router and middleware
	routes := chi.NewRouter()
	setupMiddleware(routes)
	setupRouter(routes)
	logger.GetLogger().Info("App running at port", config.GetConfig().AppPort)
	logger.GetLogger().Fatal(http.ListenAndServe(config.GetConfig().AppPort, routes))
}

func setupRouter(rr *chi.Mux) {
	router.SetupQueryRouter(rr)
}

func setupMiddleware(router *chi.Mux) {
	router.Use(
		middleware.DefaultCompress, // compress results, mostly gzipping assets and json
		middleware.StripSlashes,    // match paths with a trailing slash, strip it, and continue routing through the mux
		middleware.Recoverer,       // recover from panics without crashing server
	)

	if config.GetConfig().DebugMode {
		router.Use(middleware.Logger) // log api request calls
	}
}
