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
)

type Fail struct {
	Code    string
	Message string
	Status  int
}

func (e *Fail) Error() string {
	return e.Message
}
