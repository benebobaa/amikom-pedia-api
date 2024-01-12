package exception

type OtpError struct {
	Error string `json:"error"`
}

func NewOtpError(error string) OtpError {
	return OtpError{Error: error}
}

func NewEmailAlreadyVerified(error string) OtpError {
	return OtpError{Error: error}
}
