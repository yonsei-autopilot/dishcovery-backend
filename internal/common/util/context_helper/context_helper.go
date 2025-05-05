package contextHelper

import (
	"context"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/fail"
)

type ContextKey string // custom type to avoid context collision

const ContextKeyUserId ContextKey = "userID"

func GetUserId(ctx context.Context) (string, *fail.Fail) {
	id, ok := ctx.Value(ContextKeyUserId).(string)
	if !ok {
		return "", &fail.UserIdNotInContext
	}
	return id, nil
}

func SaveUserId(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, ContextKeyUserId, id)
}
