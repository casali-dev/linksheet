package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/casali-dev/linksheet/auth"
	"github.com/casali-dev/linksheet/db"
	"github.com/casali-dev/linksheet/repositories"
	"github.com/golang-jwt/jwt/v5"
)

func validateJWT(tokenString string) (jwt.MapClaims, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, jwt.ErrInvalidKeyType
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		apiKey := r.Header.Get("X-API-Key")

		if token != "" && strings.HasPrefix(token, "Bearer ") {
			claims, err := validateJWT(strings.TrimPrefix(token, "Bearer "))
			if err != nil {
				WriteError(w, http.StatusUnauthorized, "Invalid JWT")
				return
			}

			authorID := claims["sub"].(string)
			ctx := auth.WithAuthInfo(r.Context(), auth.AuthInfo{
				AuthorID: authorID,
				AuthType: "jwt",
			})
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		if apiKey != "" {
			repo := repositories.NewAPIKeyRepository(db.DB)
			authorID, err := repo.GetAuthorIDByKey(apiKey)
			if err != nil || authorID == "" {
				WriteError(w, http.StatusUnauthorized, "Invalid API Key")
				return
			}
			ctx := auth.WithAuthInfo(r.Context(), auth.AuthInfo{
				AuthorID: authorID,
				AuthType: "apikey",
			})
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		WriteError(w, http.StatusUnauthorized, "Missing Authorization")
	})
}
