package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/util"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/fail"
)

const (
	accessTokenDurationHours  float32 = 0.25   // 15분
	refreshTokenDurationHours float32 = 24 * 7 // 1주일
	issValue                  string  = "dishcovery"
)

var secretKey string

func InitializeSecretKey() {
	key, err := util.GetEnv("JWT_SECRET_KEY")
	if err != nil {
		panic("missing jwt secret key")
	}
	secretKey = key
}

func CreateTokens(id string) (string, string, *fail.Fail) {
	accessToken, fail := create(id, accessTokenDurationHours)
	if fail != nil {
		return "", "", fail
	}

	refreshToken, fail := create(id, refreshTokenDurationHours)
	if fail != nil {
		return "", "", fail
	}

	return accessToken, refreshToken, nil
}

func CreateRefreshToken(id string) (string, *fail.Fail) {
	refreshToken, fail := create(id, refreshTokenDurationHours)
	if fail != nil {
		return "", fail
	}
	return refreshToken, nil
}

func VerifyAccessToken(accessToken string) (string, *fail.Fail) {
	id, fail := verify(accessToken)
	if fail != nil {
		return "", fail
	}
	return id, nil
}

func VerifyRefreshToken(refreshToken string) (string, *fail.Fail) {
	id, fail := verify(refreshToken)
	if fail != nil {
		return "", fail
	}
	return id, nil
}

func create(id string, durationHours float32) (string, *fail.Fail) {
	expiration := time.Duration(durationHours * float32(time.Hour))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": issValue,
		"sub": id,
		"exp": time.Now().Add(expiration).Unix(),
		"iat": time.Now().Unix(),
	})

	result, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", &fail.FailedCreatingToken
	}
	return result, nil
}

func verify(tokenString string) (string, *fail.Fail) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return "", &fail.SigningMethodMismatch
	}

	id, fail := openClaims(token)
	if fail != nil {
		return "", fail
	}

	return id, nil
}

func openClaims(token *jwt.Token) (string, *fail.Fail) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", &fail.InvalidClaims
	}

	if iss, ok := claims["iss"].(string); !ok || iss != issValue {
		return "", &fail.InvalidIssuer
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return "", &fail.InvalidClaims
	}
	if int64(exp) < time.Now().Unix() {
		return "", &fail.TokenExpired
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return "", &fail.InvalidIssuer
	}

	return sub, nil
}
