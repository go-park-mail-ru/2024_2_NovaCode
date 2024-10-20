package httpErrors

import "net/http"

const (
	StrBadRequest                 = "bad request"
	StrUnauthorized               = "unauthorized"
	StrForbidden                  = "forbidden"
	StrNotFound                   = "not found"
	StrInternalServerError        = "internal server error"
	StrUsernameAlreadyExists      = "user with that username already exists"
	StrEmailAlreadyExists         = "user with that email already exists"
	StrUsernameAvailabilityFailed = "failed to check username availability"
	StrEmailAvailabilityFailed    = "failed to check email availability"
	StrHashPasswordFailed         = "failed to hash password"
	StrCreateUserFailed           = "failed to create user"
	StrGenerateTokenFailed        = "failed to generate token"
	StrInvalidUsernamePassword    = "invalid username or password"
	StrCreateDTOFailed            = "Can't create DTO"
)

type RestError interface {
	Status() int
	Error() string
	Causes() interface{}
}

type RestErrorStruct struct {
	status int
	error  string
	causes interface{}
}

func (e RestErrorStruct) Status() int {
	return e.status
}

func (e RestErrorStruct) Error() string {
	return e.error
}

func (e RestErrorStruct) Causes() interface{} {
	return e.causes
}

func NewRestError(status int, err string, causes interface{}) RestError {
	return RestErrorStruct{
		status: status,
		error:  err,
		causes: causes,
	}
}

func NewBadRequestError(causes interface{}) RestError {
	return RestErrorStruct{
		status: http.StatusBadRequest,
		error:  StrBadRequest,
		causes: causes,
	}
}

func NewUnauthorizedError(causes interface{}) RestError {
	return RestErrorStruct{
		status: http.StatusUnauthorized,
		error:  StrUnauthorized,
		causes: causes,
	}
}

func NewForbiddenError(causes interface{}) RestError {
	return RestErrorStruct{
		status: http.StatusForbidden,
		error:  StrForbidden,
		causes: causes,
	}
}

func NewNotFoundError(causes interface{}) RestError {
	return RestErrorStruct{
		status: http.StatusNotFound,
		error:  StrNotFound,
		causes: causes,
	}
}

func NewInternalServerError(causes interface{}) RestError {
	return RestErrorStruct{
		status: http.StatusInternalServerError,
		error:  StrInternalServerError,
		causes: causes,
	}
}
