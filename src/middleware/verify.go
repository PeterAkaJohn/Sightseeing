package middleware

import (
	"context"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

//MySigningKey : used for jwt authentication
var mySigningKey = []byte("secrets")

func VerifyToken(r *http.Request) (jwt.MapClaims, error) {
	token, err := jwt.Parse(r.Header.Get("Authorization"), func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return mySigningKey, nil
	})
	if token == nil {
		log.Print(err)
		return nil, err
	}
	if err != nil && !token.Valid {
		log.Print(err)
		return nil, err
	}
	return token.Claims.(jwt.MapClaims), nil
}

func VerifyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := jwt.Parse(r.Header.Get("Authorization"), func(token *jwt.Token) (interface{}, error) {
			// since we only use the one private key to sign the tokens,
			// we also only use its public counter part to verify
			return mySigningKey, nil
		})
		if token == nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if err != nil && !token.Valid {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		ctx := newContextWithUserID(r.Context(), r, token.Claims.(jwt.MapClaims))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func newContextWithUserID(ctx context.Context, req *http.Request, claims jwt.MapClaims) context.Context {
	userIDFromClaims := claims["id"].(float64)
	userID := int64(userIDFromClaims)
	return context.WithValue(ctx, 0, userID)
}

func UserIDFromContext(ctx context.Context) int64 {
	return ctx.Value(0).(int64)
}

//AddFavorite Handler to retrieve Context use userID := UserIDFromContext(r.Context())
