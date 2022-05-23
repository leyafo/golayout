package helper

import (
	"net/http"

	"github.com/go-chi/jwtauth/v5"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("!fj22hello+world_!#*"), nil)

	// // For debugging/example purposes, we generate and print
	// // a sample jwt token with claims `user_id:123` here:
	// _, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 123})
	// fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
}

func GetJWTAuth() *jwtauth.JWTAuth {
	return tokenAuth
}

func GetClaims(r *http.Request) map[string]interface{} {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		return make(map[string]interface{})
	}
	return claims
}

func GetUserID(r *http.Request) uint64 {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		return 0
	}
	return claims["user_id"].(uint64)
}
