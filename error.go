package tts

import "fmt"

type ErrEngineBadRequest struct {
	Msg string
}

func (e ErrEngineBadRequest) Error() string {
	return fmt.Sprintf("tts got bad request error with message `%s`", e.Msg)
}

// check if error is BadRequest response from TTS engine
func IsErrEngineBadRequest(err error) bool {
	_, ok := err.(ErrEngineBadRequest)
	return ok
}
