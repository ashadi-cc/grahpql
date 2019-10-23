package query

import (
	"context"
	"gql-ashadi/gql/schema"
	"log"
	"testing"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/gqltesting"
)

var (
	rootSchema *graphql.Schema
	ctx        = context.Background()
)

func TestSetRootSchema(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	sc, err := schema.GetRootSchema()
	if err != nil {
		log.Fatal(err)
	}

	rootSchema = graphql.MustParseSchema(sc, &Resolver{})
}

func TestUserQuery(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  rootSchema,
			Query: `
				{
					user(email:"test@1.com") {
						id
						email
					}
				}
			`,
			ExpectedResult: `
				{
					"user": null
				}
			`,
		},
	})
}
