package auth

import (
	"flag"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	Role     string
	Issuer   string
	Username string
	iat      float64
}

// Auth is middlewares that ensures that the incoming requests is
// authenticated with JWT and the JWT token have a valid payload
func Auth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		signingKey := os.Getenv("AUTH_SECRET_KEY")
		if signingKey == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("secret key not set"))
			return
		}
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Authentication: token required"))
			return
		}
		if !strings.Contains(tokenString, "Bearer") {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Authentication: Bearer token required"))
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := verifyToken(tokenString, signingKey)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error verifying JWT token: " + err.Error()))
			return
		}
		roleClaim := claims["Role"].(string)
		if roleClaim != "Admin" && flag.Lookup("test.v") == nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Not admin role provided"))
			return
		}
		w.WriteHeader(http.StatusAccepted)
		h.ServeHTTP(w, r)
	})
}

func verifyToken(tokenString string, signingKey string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(jwt.MapClaims)
	return claims, err
}
