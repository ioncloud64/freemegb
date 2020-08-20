// +build windows,!linux,!freebsd,!netbsd,!openbsd,!darwin,!js

package core

import (
	"path/filepath"

	"github.com/go-toast/toast"
)

// Notify is a function to display notifications across all platforms (Mac OS X, Linux, and MS Windows 10)
// This particular version is designed for Windows.
// It utilizes the go-toast/toast library to display native notifications on Windows 10
func Notify(message string) {
	fileLocation, _ := filepath.Abs("ui/freemegb.png")
	notification := &toast.Notification{
		AppID:   "FreeMe!GB",
		Title:   "Unknown Instruction",
		Message: message,
		Icon:    fileLocation,
		Actions: []toast.Action{
			{Type: "system", Label: "Dismiss", Arguments: "dismiss"},
		},
	}
	err := notification.Push()
	if err != nil {
		Logger.Log(LogTypes.ERROR, err)
	}
}
