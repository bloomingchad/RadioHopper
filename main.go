package main

import (
	"fmt"
	"log"
	"time"

	"github.com/adrg/libvlc-go/v3"
	"github.com/rivo/tview"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// const ah_fm = "https://nl.ah.fm/live"
const _1fmdeephouse = "https://strm112.1.fm/deephouse_128"

func main() {
	fmt.Println("RadioHopper is warming up...")

	_ = tview.NewBox()
	err := vlc.Init("--no-video", "--aout=alsa") // , "--verbose=1")
	check(err)
	defer func() {
		err = vlc.Release()
		check(err)
	}()

	player, err := vlc.NewPlayer()
	check(err)
	defer func() {
		err = player.Release()
		check(err)
	}()

	media, err := player.LoadMediaFromURL(_1fmdeephouse)
	check(err)

	err = player.SetMedia(media)
	check(err)
	defer func() {
		err = media.Release()
		check(err)
	}()

	err = player.Play()
	check(err)

	fmt.Println("please wait milord...")
	for {
		time.Sleep(30 * time.Millisecond)
		if player.IsPlaying() {
			break
		}
	}

	fmt.Println("now playing...")
	for player.IsPlaying() {
		time.Sleep(1 * time.Second)

	}

	fmt.Println("music stopped")
}
