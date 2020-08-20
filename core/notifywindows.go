// +build windows,!linux,!freebsd,!netbsd,!openbsd,!darwin,!js

package core

import (
	"path/filepath"

	"github.com/go-toast/toast"
)

func Notify(message string) {
	file_location, _ := filepath.Abs("ui/freemegb.png")
	notification := &toast.Notification{
		AppID:   "FreeMe!GB",
		Title:   "Unknown Instruction",
		Message: message,
		Icon:    file_location,
		Actions: []toast.Action{
			{"system", "Dismiss", "dismiss"},
		},
	}
	err := notification.Push()
	if err != nil {
		Logger.Log(LogTypes.ERROR, err)
	}
}
