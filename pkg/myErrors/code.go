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
	ErrorNotEnoughMoney

	ErrorNotExistStory
	ErrorNotExistTask
	ErrorNotExistAdmin
)
