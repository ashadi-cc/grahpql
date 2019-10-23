package router

import (
	"gql-ashadi/api/handler"
	"gql-ashadi/gql/query"
	"gql-ashadi/gql/schema"
	"gql-ashadi/service/logger"

	"github.com/go-chi/chi"
	graphql "github.com/graph-gophers/graphql-go"
)

//SetupQueryRouter set query router
func SetupQueryRouter(router *chi.Mux) {
	logger.GetLogger().Info("setup query router")
	s, err := schema.GetRootSchema()
	if err != nil {
		logger.GetLogger().Fatalf("reading schema error: %s", err.Error())
	}
	resolver := &query.Resolver{}
	schema := graphql.MustParseSchema(s, resolver)
	handle := &handler.GraphQL{Schema: schema}
	router.Handle("/query", handle)
}
