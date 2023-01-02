package helpers

import (
	"context"
)

type UserIdentityContextKey string

const userIdentityKey UserIdentityContextKey = "userIdentity"

func AddUserIdentity(ctx context.Context, userIdentity string) context.Context {
	return context.WithValue(ctx, userIdentityKey, userIdentity)
}

func GetUserIdentity(ctx context.Context) string {
	return ctx.Value(userIdentityKey).(string)
}
