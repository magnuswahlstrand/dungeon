package audio

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
	"github.com/hajimehoshi/ebiten/audio/wav"
	"github.com/magnuswahlstrand/dungeon/assets"
)

const (
	sampleRate = 44100
)

var HookSound []byte
var DeathSound []byte
var SpawnSound []byte
var MissedSound []byte
var VictorySound []byte

// Player represents the current audio state.
type Player struct {
	audioContext *audio.Context
	audioPlayer  *audio.Player
	current      time.Duration
	total        time.Duration
	seBytes      []byte
	volume128    int
}

var globalAudioContext *audio.Context

func LoadResources() error {

	var err error
	globalAudioContext, err = audio.NewContext(sampleRate)
	if err != nil {
		log.Fatal(err)
	}

	const bytesPerSample = 4 // TODO: This should be defined in audio package
	path := "assets/audio/spawn.mp3"

	// Sound pack 01
	s, err := mp3.Decode(globalAudioContext, audio.BytesReadSeekCloser(assets.LookupFatal(path)))
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadAll(s)
	if err != nil {
		log.Fatal(err)
	}
	SpawnSound = bytes

	// From platformer_jumping
	path = "assets/audio/jump_11.wav"
	s3, err := wav.Decode(globalAudioContext, audio.BytesReadSeekCloser(assets.LookupFatal(path)))
	if err != nil {
		log.Fatal("failed to load", err)
	}
	b, err := ioutil.ReadAll(s3)
	if err != nil {
		log.Fatal("failed to read", err)
	}
	HookSound = b

	// From 8-Bit Sound Library
	path = "assets/audio/hit.mp3"
	s2, err := mp3.Decode(globalAudioContext, audio.BytesReadSeekCloser(assets.LookupFatal(path)))
	if err != nil {
		log.Fatal("failed to load", err)
	}
	b, err = ioutil.ReadAll(s2)
	if err != nil {
		log.Fatal("failed to read", err)
	}
	DeathSound = b

	// Sound pack 01
	path = "assets/audio/missed.mp3"
	s2, err = mp3.Decode(globalAudioContext, audio.BytesReadSeekCloser(assets.LookupFatal(path)))
	if err != nil {
		log.Fatal("failed to load", err)
	}
	b, err = ioutil.ReadAll(s2)
	if err != nil {
		log.Fatal("failed to read", err)
	}
	MissedSound = b

	// Sound pack 01
	path = "assets/audio/victory.mp3"
	s2, err = mp3.Decode(globalAudioContext, audio.BytesReadSeekCloser(assets.LookupFatal(path)))
	if err != nil {
		log.Fatal("failed to load", err)
	}
	b, err = ioutil.ReadAll(s2)
	if err != nil {
		log.Fatal("failed to read", err)
	}
	VictorySound = b

	return nil
}

func Play(b []byte, volume ...float64) {
	fmt.Println("play audio!", len(b))
	if b == nil || len(b) == 0 {
		log.Println("tried to play empty bytes")
		return
	}
	tmpP, err := audio.NewPlayerFromBytes(globalAudioContext, b)
	if err != nil {
		log.Fatal(err)
	}
	if len(volume) > 0 {
		tmpP.SetVolume(volume[0])
	}
	tmpP.Play()
}
