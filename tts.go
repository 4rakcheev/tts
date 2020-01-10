package tts

type Format int
type Voice string
type Language string
type Rate int
type VoiceID string

const (
	MP3 Format = iota
	OGG
)

const (
	R8000  Rate = 8000
	R16000 Rate = 16000
	R22050 Rate = 22050
)

type TTS interface {
	Speech(text string) ([]byte, error)

	Language(lang Language)
	Voice(voice Voice)
	Format(format Format)
	SampleRate(rate Rate)
}
