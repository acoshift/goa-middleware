package jwtGoogle

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa/middleware/security/jwt"
)

type jwtGoogleResolver struct {
	*sync.RWMutex
	keys []jwt.Key
	exp  time.Time
}

// NewJWTGoogleResolver create new jwt key resolver
func NewJWTGoogleResolver() jwt.KeyResolver {
	return &jwtGoogleResolver{RWMutex: &sync.RWMutex{}}
}

func (r *jwtGoogleResolver) fetchKeys() error {
	r.Lock()
	defer r.Unlock()
	resp, err := http.Get("https://www.googleapis.com/robot/v1/metadata/x509/securetoken@system.gserviceaccount.com")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	r.exp, _ = time.Parse(time.RFC1123, resp.Header.Get("Expires"))

	m := map[string]string{}
	if err = json.NewDecoder(resp.Body).Decode(&m); err != nil {
		return err
	}
	k := []jwt.Key{}
	for _, v := range m {
		p, _ := jwtgo.ParseRSAPublicKeyFromPEM([]byte(v))
		if p != nil {
			k = append(k, p)
		}
	}
	r.keys = k
	return nil
}

func (r *jwtGoogleResolver) SelectKeys(req *http.Request) []jwt.Key {
	r.RLock()
	if r.exp.IsZero() || r.exp.Before(time.Now()) || len(r.keys) == 0 {
		r.RUnlock()
		if err := r.fetchKeys(); err != nil {
			return nil
		}
		r.RLock()
	}
	defer r.RUnlock()
	return r.keys
}
