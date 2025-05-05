package fail

import "net/http"

var (
	// controller layer fail
	InvalidJsonBody         = Fail{Code: "3001", Message: "Invalid Json Body", Status: http.StatusBadRequest}
	RequestValidationFailed = Fail{Code: "3002", Message: "Request Validation Failed", Status: http.StatusBadRequest}
	InvalidImage            = Fail{Code: "3003", Message: "Invalid Image", Status: http.StatusBadRequest}
	ImageReadFailed         = Fail{Code: "3004", Message: "Image Read Failed", Status: http.StatusBadRequest}
	InvalidImageFormat      = Fail{Code: "3005", Message: "Invalid Image Format", Status: http.StatusBadRequest}
	UnsupportedImageFormat  = Fail{Code: "3006", Message: "Unsupported Image Format", Status: http.StatusBadRequest}
	UserIdNotInContext      = Fail{Code: "3007", Message: "User Id Not In Context", Status: http.StatusBadRequest}

	// business layer fail
	UserNotGoogleAuthenticated  = Fail{Code: "4001", Message: "User Not Google Authenticated", Status: http.StatusUnauthorized}
	UserNotFullyRegistered      = Fail{Code: "4002", Message: "User Not Fully Registered", Status: http.StatusNotFound}
	FailedSavingUser            = Fail{Code: "4003", Message: "Failed Saving User", Status: http.StatusInternalServerError}
	UserNotFound                = Fail{Code: "4004", Message: "User Not Found", Status: http.StatusNotFound}
	FailedDescriptionGeneration = Fail{Code: "4005", Message: "Failed Description Generation", Status: http.StatusInternalServerError}

	// token related fail
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
