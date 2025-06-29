package main

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/adrg/libvlc-go/v3"
)

func main() {
	fmt.Println("RadioHopper is warming up...")

	_ = tview.NewBox()
    player, err := vlc.NewPlayer()
    if err != nil {
    	
    }
    defer player.Release()
    
    fmt.Println(player.IsPlaying())
}
