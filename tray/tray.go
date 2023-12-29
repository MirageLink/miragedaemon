package tray

import (
	"fmt"
	"github.com/getlantern/systray"
	"os"
)

func OnReady() {
	systray.SetTitle("guh")
	mQuitOrig := systray.AddMenuItem("Quit", "guh ")
	go func() {
		<-mQuitOrig.ClickedCh
		systray.Quit()
	}()
}

func OnExit() {
	fmt.Println("does this run")
	os.Exit(0)
}
