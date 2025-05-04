package fail

import "net/http"

var (
	UserNotGoogleAuthenticated = Fail{Code: "4001", Message: "User Not Google Authenticated", Status: http.StatusUnauthorized}
	UserNotFullyRegistered     = Fail{Code: "4002", Message: "User Not Fully Registered", Status: http.StatusNotFound}
	FailedSavingUser           = Fail{Code: "4003", Message: "Failed Saving User", Status: http.StatusInternalServerError}
)

type Fail struct {
	Code    string
	Message string
	Status  int
}

func (e *Fail) Error() string {
	return e.Message
}
