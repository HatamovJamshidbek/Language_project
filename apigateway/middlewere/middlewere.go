package middlewere

import (
	"api_service/api/token"
	"fmt"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, fmt.Errorf("authorization header is required"))
			return
		}

		valid, err := token.ValidateToken(auth)
		if err != nil || !valid {
			fmt.Println("salom")
			fmt.Println("333333", auth)
			c.AbortWithStatusJSON(http.StatusUnauthorized, fmt.Errorf("invalid token: %s", err))
			return
		}

		claims, err := token.ExtractClaims(auth)
		if err != nil || !valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, fmt.Errorf("invalid token claims: %s", err))
			return
		}

		c.Set("claims", claims)
		fmt.Println(claims)
		c.Next()
	}
}

func CasbinMiddleware(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("saladsfdsfadsfadsjfadsfijadsif")
		// Claimsni olish
		claims, exists := c.Get("claims")
		if !exists {
			fmt.Println("Claims not found in context")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		jwtClaims, ok := claims.(jwt.MapClaims)
		if !ok {
			fmt.Println("Claims are not of type jwt.MapClaims")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Claims ma'lumotlarini olish
		sub, ok := jwtClaims["role"].(string)
		if !ok {
			fmt.Println("Role claim is missing or not a string")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		obj := c.FullPath()
		act := c.Request.Method
		fmt.Println(sub, obj, act)

		allowed, err := enforcer.Enforce(sub, obj, act)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error occurred during authorization"})
			return
		}

		if !allowed {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}

		c.Next()
	}
}
