package tts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const googleAPI = "https://texttospeech.googleapis.com"

type GoogleEngine struct {
	apiKey      string
	apiEndpoint string
	client      *http.Client
	request     googleTTSRequest
}

type (
	googleTTSRequest struct {
		Input       input       `json:"input,required"`
		Voice       voice       `json:"voice,required"`
		AudioConfig audioConfig `json:"audioConfig,required"`
	}
	input struct {
		Text string `json:"text,required"`
		SSML string `json:"ssml,omitempty"`
	}
	voice struct {
		LanguageCode string `json:"languageCode,required"`
		Name         string `json:"name,omitempty"`
		SSMLGender   string `json:"ssmlGender,omitempty"` // https://cloud.google.com/text-to-speech/docs/reference/rest/v1/SsmlVoiceGender
	}
	audioConfig struct {
		AudioEncoding    string   `json:"audioEncoding,required"` // https://cloud.google.com/text-to-speech/docs/reference/rest/v1/text/synthesize#AudioEncoding
		SpeakingRate     int      `json:"speakingRate,omitempty"`
		Pitch            int      `json:"pitch,omitempty"`
		VolumeGainDb     int      `json:"volumeGainDb,omitempty"`
		SampleRateHertz  int      `json:"sampleRateHertz,omitempty"`
		EffectsProfileID []string `json:"effectsProfileId,omitempty"`
	}
)

func NewGoogleTTS(apiKey string) TTS {
	return &GoogleEngine{
		apiKey:      apiKey,
		apiEndpoint: googleAPI,
		client:      &http.Client{},
		request: googleTTSRequest{
			Voice: voice{
				LanguageCode: "en-US",
				Name:         "",
			},
			AudioConfig: audioConfig{
				AudioEncoding: "MP3",
				SpeakingRate:  22050,
			},
		}}
}

func (t *GoogleEngine) Format(format Format) {
	switch format {
	case MP3:
		t.request.AudioConfig.AudioEncoding = "MP3"
	case OGG:
		t.request.AudioConfig.AudioEncoding = "OGG_OPUS"
	}
}

func (t *GoogleEngine) SampleRate(rate Rate) {
	t.request.AudioConfig.SampleRateHertz = int(rate)
}

func (t *GoogleEngine) Voice(voice Voice) {
	t.request.Voice.Name = fmt.Sprintf("%s", voice)
}

func (t *GoogleEngine) Language(lang Language) {
	t.request.Voice.LanguageCode = fmt.Sprintf("%s", lang)
}

func (t *GoogleEngine) Speech(text string) ([]byte, error) {
	t.request.Input.Text = text

	b, err := json.Marshal(t.request)
	if err != nil {
		return nil, err
	}

	r, _ := http.NewRequest("POST", fmt.Sprintf("%s/v1/text:synthesize?key=%s", googleAPI, t.apiKey), bytes.NewReader(b))
	r.Header.Set("Content-Type", "application/json")

	res, err := t.client.Do(r)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	} else if res.StatusCode == 400 {
		return nil, ErrEngineBadRequest{Msg: string(data)}
	} else if res.StatusCode == 403 {
		return nil, ErrEngineAuthorizationFailed{Msg: string(data)}
	} else if res.StatusCode != 200 {
		return nil, fmt.Errorf("api got http error with code `%d` and body `%q`", res.StatusCode, data)
	}

	return data, nil
}

func (t *GoogleEngine) EndpointAPI(url string) {
	t.apiEndpoint = url
}
