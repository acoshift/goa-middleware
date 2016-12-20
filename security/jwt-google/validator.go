package jwtGoogle

import (
	"net/http"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
	"golang.org/x/net/context"
)

type contextKey int

const (
	uidKey contextKey = iota + 1
)

func some(arr []string, value string) bool {
  for _, x := range arr {
    if
  }
}

func validate(opts Options) goa.Middleware {
	m, _ := goa.NewMiddleware(func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, r *http.Request) error {
			token := jwt.ContextJWT(ctx)
			if token == nil {
				return jwt.ErrJWTError("token not found")
			}
			claims := token.Claims.(jwtgo.MapClaims)

      ok := false
      for _, issuer := range opts.Issuers {
        if claims.VerifyIssuer(opts.Issuers, true) {
          ok = true
          break
        }
      }
      if !ok {
				return jwt.ErrJWTError("wrong issuer")
      }

      ok = false
      for _, aud := range opts.Audiences {
        if claims.VerifyAudience(aud, true) {
          ok = true
          break
        }
      }
      if !ok {
        return jwt.ErrJWTError("wrong audience")
      }

			uid := claims["sub"].(string)
			nctx := context.WithValue(ctx, uidKey, uid)
			return h(nctx, rw, r)
		}
	})
	return m
}
