package jwtGoogle

import "golang.org/x/net/context"

type contextKey int

const (
	uidKey contextKey = iota + 1
)

// ContextUserID retrieves JWT Subject from a context
func ContextUserID(ctx context.Context) string {
	return ctx.Value(uidKey).(string)
}
