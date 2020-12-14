package statusmsg

type SuccessMessage struct {
	Success string `json:"success"`
}

type StatusMessage struct {
	Code    int32  `json:"errorCode,omitempty"`
	Message string `json:"errorMessage"`
}

type Status struct {
	HttpStatus int
	StatusMsg  StatusMessage
}

var (
	StatusSuccess = SuccessMessage{
		Success: "ok",
	}

	StatusMsgProcessError = StatusMessage{}

	StatusMsgInvalidArg = StatusMessage{
		Code:    92001,
		Message: "bad_params",
	}

	StatusMsgUserNotFound = StatusMessage{
		Code:    10010,
		Message: "user_not_found",
	}

	StatusMsgUserExists = StatusMessage{
		Code:    10011,
		Message: "user_exists",
	}
	StatusMsgUserPasswordNotEqual = StatusMessage{
		Code:    10012,
		Message: "password_not_equal",
	}

	StatusMsgCompanyExists = StatusMessage{
		Code:    10012,
		Message: "company_exists",
	}

	StatusMsgInvalidOAuthCode = StatusMessage{
		Code:    11003,
		Message: "invalid_oauth_code",
	}
)

func (status *StatusMessage) WithMessage(message string) *StatusMessage {
	return &StatusMessage{
		Code:    status.Code,
		Message: status.Message + " ERR: " + message,
	}
}
