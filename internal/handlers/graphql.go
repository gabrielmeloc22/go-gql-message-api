package handlers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gabrielmeloc22/go-message-api/internal/context"
	"github.com/gabrielmeloc22/go-message-api/internal/generated"
	"github.com/gabrielmeloc22/go-message-api/internal/resolvers"
	"github.com/gin-gonic/gin"
)

func GraphqlHandler() gin.HandlerFunc {
	c := generated.Config{Resolvers: &resolvers.Resolver{}}

	c.Directives.Authenticated = context.AuthDirectiveHandler

	h := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
