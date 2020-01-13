package tts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bmizerany/aws4"
	"io/ioutil"
	"net/http"
)

const pollyAPI = "https://polly.us-west-2.amazonaws.com"

type Polly struct {
	client      *aws4.Client
	endpointAPI string
	request     pollyRequest
}

type pollyRequest struct {
	LanguageCode string `json:"LanguageCode"`
	OutputFormat string `json:"OutputFormat"`
	SampleRate   string `json:"SampleRate"`
	Text         string `json:"Text"`
	VoiceId      string `json:"VoiceId"`
}

func NewPollyTTS(accessKey string, secretKey string) TTS {
	return &Polly{
		client: &aws4.Client{Keys: &aws4.Keys{
			AccessKey: accessKey,
			SecretKey: secretKey,
		}},
		endpointAPI: pollyAPI,
		request: pollyRequest{
			LanguageCode: "en-US",
			OutputFormat: "mp3",
			SampleRate:   "22050",
			VoiceId:      "Brian",
		}}
}

func (t *Polly) Format(format Format) {
	switch format {
	case MP3:
		t.request.OutputFormat = "mp3"
	case OGG:
		t.request.OutputFormat = "ogg_vorbis"
	}
}

func (t *Polly) SampleRate(rate Rate) {
	t.request.SampleRate = fmt.Sprintf("%d", rate)
}

func (t *Polly) Voice(voice Voice) {
	t.request.VoiceId = fmt.Sprintf("%s", voice)
}

func (t *Polly) Language(lang Language) {
	t.request.LanguageCode = fmt.Sprintf("%s", lang)
}

func (t *Polly) Speech(text string) ([]byte, error) {
	t.request.Text = text

	b, err := json.Marshal(t.request)
	if err != nil {
		return nil, err
	}

	r, _ := http.NewRequest("POST", t.endpointAPI+"/v1/speech", bytes.NewReader(b))
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

func (t *Polly) EndpointAPI(url string) {
	t.endpointAPI = url
}
