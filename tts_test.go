package tts

import (
	"testing"
)

func TestCreateNewPollyEngine(t *testing.T) {
	tts := NewPollyTTS("fakeKey", "fakeSecret")
	if tts == nil {
		t.Error("polly TTS result is empty after NewPollyTTS()")
	}
}

func TestCreateNewGoogleCloudEngine(t *testing.T) {
	tts := NewGoogleTTS("fakeAPIkey")
	if tts == nil {
		t.Error("google cloud TTS result is empty after NewGoogleTTS()")
	}
}
