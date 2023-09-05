package middleware

import (
	"context"
	"duongGoGo/infra/caching"
	"duongGoGo/utils/tokenutil"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/redis/go-redis/v9"
	"net/http"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		rds := caching.GetRedisClient()

		accessTokenString, err := c.Cookie("ACCESS_TOKEN")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		accessToken, err := tokenutil.ParseToken(accessTokenString, tokenutil.ACCESS)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		if claims, ok := accessToken.Claims.(jwt.MapClaims); ok && accessToken.Valid {
			userId := claims["sub"]
			accessTokenRedisKey := tokenutil.CreateRedisKeyForToken(userId.(string), accessTokenString)
			_, err := rds.Get(ctx, accessTokenRedisKey).Result()

			if err != nil {
				if err == redis.Nil {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "token not valid"})
				} else {
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})
				}
			}
			c.Set("userId", userId)

			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func CSRFMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		rds := caching.GetRedisClient()
		csrfTokenRequest := c.GetHeader("X-CSRF-Token")

		csrfToken, err := tokenutil.ParseToken(csrfTokenRequest, tokenutil.CSRF)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		if claims, ok := csrfToken.Claims.(jwt.MapClaims); ok && csrfToken.Valid {
			userId := claims["sub"]
			csrfTokenRedisKey := tokenutil.CreateRedisKeyForToken(userId.(string), csrfTokenRequest)
			_, err := rds.Get(ctx, csrfTokenRedisKey).Result()

			if err != nil {
				if err == redis.Nil {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "token not valid"})
				} else {
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})
				}
			}
			c.Set("userId", userId)

			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
