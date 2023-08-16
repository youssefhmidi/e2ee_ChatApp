package models

import "github.com/golang-jwt/jwt/v4"

type JwtCustomClaims struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

type JwtMessageClaim struct {
	PublicKey     string `json:"public_key"`
	SignedMessage string `json:"message"`
	Id            int    `json:"id"`
	jwt.RegisteredClaims
}

// interface to make the JWT data readonly
type JwtService interface {
	//  //return the secret of the used JwtRequirement struct
	//	 //e.g: Getting the acces token secret will look like this
	//    GetSecret("access") === jwt.Secrets["access"]
	GetSecret(from string) string
	// return the Expiry as an int (multiplie it by time.Hour)
	GetExpiryTime(from string) int
}
