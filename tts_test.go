package tts

import (
	"testing"
)

func TestCreateNewPollyEngine(t *testing.T) {
	tts := NewPolly("testKey", "testSecret")
	if tts == nil {
		t.Error("polly TTS result is empty after poplly.New()")
	}
}
