package context

import (
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gabrielmeloc22/go-message-api/internal/config"
	"github.com/gabrielmeloc22/go-message-api/internal/model"

	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Auth() gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := strings.Split(c.Request.Header.Get("Authorization"), " ")

		if len(authHeader) == 2 {
			tokenString := authHeader[1]

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
				}

				return []byte("my-secret"), nil
			})

			if err != nil {
				c.Error(fmt.Errorf("error while parsing jwt"))
				return
			}

			userId, ok := token.Claims.(jwt.MapClaims)["userId"]

			if !ok {
				c.Error(fmt.Errorf("invalid jwt"))
				return
			}
			c.Set("currentUser", userId)
		}
	}
}

func AuthDirectiveHandler(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	currUser := GetCurrentUser(ctx)

	if currUser == nil {
		return nil, fmt.Errorf("not authorized")
	}

	return next(ctx)
}

func GetCurrentUser(ctx context.Context) *model.User {
	gc, _ := GinContextFromContext(ctx)

	userId, exists := gc.Get("currentUser")

	if exists {
		var user model.User
		config.Db.Where("id = ?", userId).Find(&user)

		return &user
	}

	return nil
}
