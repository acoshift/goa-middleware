package jwtGoogle

import (
	"acourse/app"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
)

// New creates new jwt google middleware
func New(opts Options) goa.Middleware {
	return jwt.New(NewJWTGoogleResolver(), validate(opts), app.NewJWTSecurity())
}
