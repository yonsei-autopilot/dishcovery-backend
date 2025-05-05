package fail

import "net/http"

var (
	// controller layer fail
	InvalidJsonBody         = Fail{Code: "3001", Message: "Invalid Json Body", Status: http.StatusBadRequest}
	RequestValidationFailed = Fail{Code: "3002", Message: "Request Validation Failed", Status: http.StatusBadRequest}

	// business layer fail
	UserNotGoogleAuthenticated = Fail{Code: "4001", Message: "User Not Google Authenticated", Status: http.StatusUnauthorized}
	UserNotFullyRegistered     = Fail{Code: "4002", Message: "User Not Fully Registered", Status: http.StatusNotFound}
	FailedSavingUser           = Fail{Code: "4003", Message: "Failed Saving User", Status: http.StatusInternalServerError}
	UserNotRegistered          = Fail{Code: "4004", Message: "User Not Registered", Status: http.StatusNotFound}

	// token related
	FailedCreatingToken   = Fail{Code: "5001", Message: "Failed Creating Token", Status: http.StatusInternalServerError}
	SigningMethodMismatch = Fail{Code: "5002", Message: "Signing Method Mismatch", Status: http.StatusUnauthorized}
	InvalidClaims         = Fail{Code: "5003", Message: "Invalid Claims", Status: http.StatusUnauthorized}
	TokenExpired          = Fail{Code: "5004", Message: "Token Expired", Status: http.StatusUnauthorized}
	InvalidIssuer         = Fail{Code: "5005", Message: "Invalid Issuer", Status: http.StatusUnauthorized}
	InvalidSubject        = Fail{Code: "5006", Message: "Invalid Subject", Status: http.StatusUnauthorized}
	TokenNotInHeader      = Fail{Code: "5007", Message: "Token Not In Header", Status: http.StatusUnauthorized}
)

type Fail struct {
	Code    string
	Message string
	Status  int
}

func (e *Fail) Error() string {
	return e.Message
}
