package tts

type Engine string
type Format int
type Voice string
type Language string
type Rate int
type VoiceID string

const (
	PollyEngine       Engine = "polly"
	GoogleCloudEngine Engine = "google cloud"
)
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
	// generates and returns audio content
	Speech(text string) ([]byte, error)

	Language(lang Language)
	Voice(voice Voice)
	Format(format Format)
	SampleRate(rate Rate)
}


//func NewTTS(engine Engine, credentials interface{}) (TTS, error) {
//	switch engine {
//	case PollyEngine:
//		_, ok := credentials.(aws4.Keys)
//		if !ok {
//			return nil, fmt.Errorf("credentials for Polly must be aws4.Keys stuct")
//		}
//		tts := NewPollyTTS(credentials.(aws4.Keys))
//		return tts, nil
//	case GoogleCloudEngine:
//		_, ok := credentials.(GoogleCredentials)
//		if !ok {
//			return nil, fmt.Errorf("credentials for Polly must be aws4.Keys stuct")
//		}
//		tts := NewPollyTTS(credentials.(aws4.Keys))
//		return tts, nil
//	}
//
//
//}
