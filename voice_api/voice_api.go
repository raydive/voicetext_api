package voice_api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	ApiToken string
}

type Config struct {
	speaker       string
	emotion       string
	emotion_level int
	pitch         int
	speed         int
	volume        int
}

// デフォルトの設定を生成
func DefualtConfig(speaker string, emotion string) *Config {
	config := new(Config)
	config.speaker = speaker
	config.emotion = emotion
	config.emotion_level = 1
	config.pitch = 100
	config.speed = 100
	config.volume = 100

	return config
}

func (config *Config) urlValue(text string) (value url.Values) {
	value = make(url.Values)
	value.Add("text", text)
	value.Add("speaker", config.speaker)
	value.Add("emotion", config.emotion)
	value.Add("emotion_level", strconv.Itoa(config.emotion_level))
	value.Add("pitch", strconv.Itoa(config.pitch))
	value.Add("speed", strconv.Itoa(config.speed))
	value.Add("volume", strconv.Itoa(config.volume))

	return
}

func (client *Client) TTS(text string, config *Config) ([]byte, error) {
	endpoint := "https://api.voicetext.jp/v1/tts"
	values := config.urlValue(text)
	fmt.Println(values.Encode())
	req, err := http.NewRequest("POST", endpoint, bytes.NewBufferString(values.Encode()))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(client.ApiToken, "")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(values.Encode())))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// とりあえず抜ける
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		// とりあえず抜ける
		return nil, err
	}

	status := resp.StatusCode
	fmt.Println(status)
	if status != http.StatusOK {
		return nil, nil
	}

	return body, nil
}
