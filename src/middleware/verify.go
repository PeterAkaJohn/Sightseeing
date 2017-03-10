package middleware

import (
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
