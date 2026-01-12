package middleware

func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := extractBearerToken(c)
		claims, err := VerifyJWT(token, secret)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()
	}
}
