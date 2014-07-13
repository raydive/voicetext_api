package main

import (
	"fmt"
	"github.com/raydive/voicetext_api/voice_api"
	"os"
)

func main() {
	fmt.Println("start")
	client := voice_api.Client{ ApiToken: os.Getenv("VOICETEXT_API_KEY") }
	config := voice_api.DefualtConfig("haruka", "happiness")

	body, err := client.TTS("おはようございます", config)
	if err != nil || body == nil {
		return
	}

	f, err := os.Create("hello.wav")
	defer f.Close()
	if err != nil {
		panic(err)
		return
	}

	if _, err := f.Write(body); err != nil {
		panic(err)
		return
	}

	fmt.Println("end")
}
