package auth

import (
	"context"
)

type AuthInfo struct {
	AuthorID string
	AuthType string
}

type contextKey string

const authContextKey = contextKey("authInfo")

func WithAuthInfo(ctx context.Context, info AuthInfo) context.Context {
	return context.WithValue(ctx, authContextKey, info)
}

func GetAuthInfo(ctx context.Context) (AuthInfo, bool) {
	info, ok := ctx.Value(authContextKey).(AuthInfo)
	return info, ok
}
