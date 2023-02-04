package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"vborys/schema"
)

func main() {
	schema, err := graphql.NewSchema(schema.DefineSchema())
	if err != nil {
		log.Panic("Schema definition error", err)
	}
	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})
	http.Handle("/graphql", h)
	log.Print("Server started on port 8080")
	err2 := http.ListenAndServe(":8080", nil)
	if err2 != nil {
		log.Panic("Server not started!", err2)
	}
}
