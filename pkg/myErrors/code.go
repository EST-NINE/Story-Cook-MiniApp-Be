package myErrors

const (
	ERROR = iota
	SUCCESS

	ErrorInvalidParams
	ErrorAuthCheckTokenFail
	ErrorAuthCheckTokenTimeout
	ErrorAuthToken
	ErrorDatabase

	ErrorCreateUser
	ErrorNotExistUser

	ErrorNotExistStory
	ErrorNotExistTask
	ErrorNotExistAdmin
)
