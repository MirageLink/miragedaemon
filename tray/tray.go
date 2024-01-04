package tray

import (
	"embed"
	"fmt"
	"github.com/getlantern/systray"
	"os"
)

//go:embed icon.ico
var content embed.FS

func OnReady() {
	icon := getIcon("icon.ico")
	systray.SetIcon(icon)
	systray.SetTitle("mirage")
	mQuitOrig := systray.AddMenuItem("Quit", "guh")
	go func() {
		<-mQuitOrig.ClickedCh
		systray.Quit()
	}()
}

func OnExit() {
	os.Exit(0)
}

// I stole this lol
func getIcon(s string) []byte {
	b, err := content.ReadFile(s)
	if err != nil {
		fmt.Print(err)
	}
	return b
}
