// Package main is the entry-point to FreeMe!GB
//
// // By ioncloud64 (Nathan Martin)
package main

// AppID is the GTK Application ID string
const AppID = "com.ioncloud64.freemegb"

// UITextView is a wrapper containing a GTK TextView
type UITextView struct {
	TextView *gtk.TextView
}

// main is the entry point of FreeMe!GB.
func main() {}

// IsWindow converts a GObject to a GTK Window.
func IsWindow(obj glib.IObject) (*gtk.Window, error) {}

// IsMenuItem converts a GObject to a GTK MenuItem.
func IsMenuItem(obj glib.IObject) (*gtk.MenuItem, error) {}

// IsFileChooserDialog converts a GObject to a GTK FileChooserDialog.
func IsFileChooserDialog(obj glib.IObject) (*gtk.FileChooserDialog, error) {}

// IsAboutDialog converts a GObject to a GTK AboutDialog.
func IsAboutDialog(obj glib.IObject) (*gtk.AboutDialog, error) {}

// IsButton converts a GObject to a GTK Button.
func IsButton(obj glib.IObject) (*gtk.Button, error) {}

// IsListStore converts a GObject to a GTK ListStore.
func IsListStore(obj glib.IObject) (*gtk.ListStore, error) {}

// IsTreeView converts a GObject to a GTK TreeView.
func IsTreeView(obj glib.IObject) (*gtk.TreeView, error) {}

// IsProgressBar converts a GObject to a GTK ProgressBar.
func IsProgressBar(obj glib.IObject) (*gtk.ProgressBar, error) {}

// IsTextView converts a GObject to a GTK TextView.
func IsTextView(obj glib.IObject) (*gtk.TextView, error) {}

// UIErrorCheck checks a previous Is* function for any UI errors.
func UIErrorCheck(err error) {}
