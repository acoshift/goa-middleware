package jwtGoogle

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
)

// New creates new jwt google middleware
func New(opts Options, scheme *goa.JWTSecurity) goa.Middleware {
	return jwt.New(NewJWTGoogleResolver(), validate(opts), scheme)
}
