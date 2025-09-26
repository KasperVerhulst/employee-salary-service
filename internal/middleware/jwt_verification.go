package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/KasperVerhulst/SalaryService/config"
	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	Company string `json:"company"`
}

type UserCtxKey struct{}

func JWTMiddleware(next http.Handler, cfg *config.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Return 401 Unauthorized if no Authorization header is present
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// 2. The header should be in the format "Bearer <token>"
		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || strings.ToLower(headerParts[0]) != "bearer" {
			http.Error(w, "Authorization header must be in Bearer {token} format", http.StatusUnauthorized)
			return
		}

		jwtTokenString := headerParts[1]

		var token *jwt.Token
		var parseErr error
		var claims *CustomClaims
		var ok bool
		if cfg.VerifySignature {

			// Create the keyfunc.Keyfunc to fetch the public key for the signature the jwk_uri
			jwks, err := keyfunc.NewDefault([]string{cfg.JWKURI})
			if err != nil {
				log.Println("Failed to create JWKS from resource: %v", err)
				http.Error(w, "Authorization header required", http.StatusInternalServerError)
				return
			}

			// Parse the JWT.
			token, parseErr = jwt.NewParser().ParseWithClaims(jwtTokenString, &CustomClaims{}, jwks.Keyfunc)
			//unknown alg, signature not valid or issue with clains
			if parseErr != nil {
				http.Error(w, "Invalid JWT", http.StatusUnauthorized)
				return
			}
			claims = token.Claims.(*CustomClaims)

		} else {
			// Parse the JWT without verifying the signature
			// expiration, nbf and iat are not validated either
			token, _, parseErr = jwt.NewParser().ParseUnverified(jwtTokenString, &CustomClaims{})
			if parseErr != nil {
				http.Error(w, "Invalid JWT format", http.StatusUnauthorized)
				return
			}

			claims, ok = token.Claims.(*CustomClaims)
			if !ok {
				http.Error(w, "Invalid JWT claims", http.StatusUnauthorized)
				return
			}

			// Manually validate the time-based claims (exp, nbf, iat).
			// The signature is still NOT checked.
			v := jwt.NewValidator()
			if err := v.Validate(claims); err != nil {
				http.Error(w, "Invalid JWT claims", http.StatusUnauthorized)
				return
			}
		}

		// Add the company from the token claims to the request context
		company := claims.Company
		ctx := context.WithValue(r.Context(), UserCtxKey{}, company)

		// end middleware, send to handler
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
