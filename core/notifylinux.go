// +build linux

package core

import (
	"github.com/esiqveland/notify"
	"github.com/godbus/dbus"
)

// Notify is a function to display notifications across all platforms (Mac OS X, Linux, and MS Windows 10)
// This particular version is designed for Linux.
// It utilizes the godbus/dbus and esiqveland/notify packages to display native notifications.
func Notify(message string) {
	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}
	iconName := "ioncloud64-freemegb"
	notif := notify.Notification{
		AppName:       "FreeMe!GB",
		ReplacesID:    uint32(0),
		AppIcon:       iconName,
		Summary:       "Unknown Instruction",
		Body:          message,
		Actions:       []string{"cancel", "Cancel", "open", "Open"}, // tuples of (action_key, label)
		Hints:         map[string]dbus.Variant{"desktop-entry": dbus.MakeVariant("freemegb")},
		ExpireTimeout: int32(5000),
	}

	notify.SendNotification(conn, notif)
}
