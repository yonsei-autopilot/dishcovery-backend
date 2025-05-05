package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/util/codec"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/util/token"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/fail"
)

type ContextKey string // custom type to avoid context collision

const ContextKeyUserId ContextKey = "userID"

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if shouldBypass(r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			codec.FailureFromFail(w, &fail.TokenNotInHeader)
			return
		}

		rawToken := strings.TrimPrefix(authHeader, "Bearer ")
		userId, fail := token.VerifyAccessToken(rawToken)
		if fail != nil {
			codec.FailureFromFail(w, fail)
			return
		}

		ctx := context.WithValue(r.Context(), ContextKeyUserId, userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func shouldBypass(path string) bool {
	return strings.HasPrefix(path, "/auth") || strings.HasPrefix(path, "/health")
}
