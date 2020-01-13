package tts

import "fmt"

type (
	ErrEngineBadRequest struct {
		Msg string
	}
	ErrEngineAuthorizationFailed struct {
		Msg string
	}
)

func (e ErrEngineBadRequest) Error() string {
	return fmt.Sprintf("tts enginge got `bad request` error with message `%s`", e.Msg)
}
func (e ErrEngineAuthorizationFailed) Error() string {
	return fmt.Sprintf("tts engine has authorizatoin problem: `%s`", e.Msg)
}

func IsErrEngineBadRequest(err error) bool {
	_, ok := err.(ErrEngineBadRequest)
	return ok
}
func IsErrEngineAuthorizationFailed(err error) bool {
	_, ok := err.(ErrEngineAuthorizationFailed)
	return ok
}
