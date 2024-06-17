package context

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

type GinContextKey struct{}

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), GinContextKey{}, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(GinContextKey{})

	if ginContext == nil {
		return nil, fmt.Errorf("could not access gin context")
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		return nil, fmt.Errorf("gin.Context could not be type casted")
	}

	return gc, nil
}
