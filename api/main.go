package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
)

type LoginPayload struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {

	// Clé secrète pour la signature HMAC (HS256)
	secretKey := os.Getenv("SECRET_KEY_JWT")
	if len(secretKey) == 0 {
		secretKey = "toto"
	}
	signingKey := []byte(secretKey)

	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	// Define a simple GET endpoint
	// curl -X POST http://localhost:8080/login -d '{"email":"monemail","password":"pass"}'
	r.POST("/login", func(c *gin.Context) {
		var payload LoginPayload
		// Bind the JSON payload to the LoginPayload struct
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		payloadClaims := jwt.New()
		payloadClaims.Set(jwt.SubjectKey, `https://github.com/lestrrat-go/jwx/jwt`)
		payloadClaims.Set(jwt.AudienceKey, `Golang Users`)
		payloadClaims.Set(jwt.IssuedAtKey, time.Unix(time.Now().Add(time.Hour*5).Round(time.Second).Unix(), 0))
		payloadClaims.Set(`privateClaimKey`, `Hello, World!`)

		signed, err := jwt.Sign(payloadClaims, jwa.HS256, signingKey)
		if err != nil {
			log.Printf("failed to sign token: %s", err)
			return
		}

		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"token": string(signed),
		})
	})

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run()
}
