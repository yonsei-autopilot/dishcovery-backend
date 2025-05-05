package middleware

import (
	"net/http"
	"strings"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/util/codec"
	contextHelper "github.com/yonsei-autopilot/smart-menu-backend/internal/common/util/context_helper"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/util/token"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/fail"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if shouldBypass(r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			codec.Failure(w, &fail.TokenNotInHeader)
			return
		}

		rawToken := strings.TrimPrefix(authHeader, "Bearer ")
		userId, fail := token.VerifyAccessToken(rawToken)
		if fail != nil {
			codec.Failure(w, fail)
			return
		}

		ctx := contextHelper.SaveUserId(r.Context(), userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func shouldBypass(path string) bool {
	return strings.HasPrefix(path, "/auth") || strings.HasPrefix(path, "/health")
}
