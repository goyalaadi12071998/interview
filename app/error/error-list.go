package errorclass

const (
	BadRequestValidationError = "BAD_REQUEST_VALIDATION_ERROR"
	BadRequestError           = "BAD_REQUEST_ERROR"
	InternalServerError       = "INTERNAL_SERVER_ERROR"
	RecordNotFound            = "BAD_REQUEST_RECORD_NOT_FOUND"
)

var errorList = map[string]*Error{
	BadRequestValidationError: {
		code: BadRequestValidationError,
		name: "BadRequestValidationError",
	},
	InternalServerError: {
		code: InternalServerError,
		name: "InternalServerError",
	},
	RecordNotFound: {
		code: RecordNotFound,
		name: "RecordNotFound",
	},
}
