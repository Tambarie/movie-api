package shared

var (
	NoResourceFound         = "this resource does not exist"
	NoRecordFound           = "sorry, no record found"
	NoErrorsFound           = "no errors at the moment"
	INVALID_MESSAGE_ERROR   = "The message format read from the given topic is invalid"
	VALIDATION_ERROR        = "The request has validation errors"
	REQUEST_NOT_FOUND       = "The requested resource was NOT found"
	GENERIC_ERROR           = "Generic error occurred. See stacktrace for details"
	AUTHORIZATION_ERROR     = "You do NOT have adequate permission to access this resource"
	DUPLICATE_ENTRY_ERROR   = "Duplicate entry detected."
	STRICT_VALIDATION_ERROR = "The request failed strict validation test."
	EXPIRED_OTP_CODE_ERROR  = "The given OTP code has expired"
	NO_PRINCIPAL            = "Principal identifier NOT provided"
)
