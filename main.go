// Package main is the entry-point to FreeMe!GB
//
// By ioncloud64 (Nathan Martin)
package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ioncloud64/freemegb/core"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

// AppID is the GTK Application ID string
const AppID = "com.ioncloud64.freemegb"

// UITextView is a wrapper containing a GTK TextView
type UITextView struct {
	TextView *gtk.TextView
}

func (TV *UITextView) Write(data []byte) (n int, err error) {
	// TODO: Implement proper TextView Console
	// buff, err := TV.TextView.GetBuffer()
	// text, err := buff.GetText(buff.GetStartIter(),
	// 	buff.GetEndIter(), true)
	// buff.SetText(text + string(data))
	// TV.TextView.ScrollToIter(buff.GetEndIter(), 0.0, true, 0.5, 0.5)
	// var mark = buff.CreateMark("end", buff.GetEndIter(), false)
	// TV.TextView.ScrollToMark(mark, 0.0, false, 0.5, 0.5)
	return len(data), err
}

// main is the entry point of FreeMe!GB.
func main() {
	core.Init()
	var System = core.System

	fmt.Println(core.LogFilename)

	defer core.LogFile.Close()

	UI(&System)
}

// UI receives the core system and sets up the UI of FreeMe!GB.
func UI(System *core.SystemType) {
	app, err := gtk.ApplicationNew(AppID, glib.APPLICATION_FLAGS_NONE)
	UIErrorCheck(err)

	core.AppRef = app

	app.Connect("startup", func() {
		core.Logger.Println("FreeMe!GB is starting up...")
	})
	app.Connect("activate", func() {
		core.Logger.Println("FreeMe!GB is activated...")

		builder, err := gtk.BuilderNewFromFile("ui/MainWindow.glade")
		UIErrorCheck(err)

		obj, err := builder.GetObject("MainWindow")
		UIErrorCheck(err)

		// Open MenuItem
		menuOpen, err := builder.GetObject("menuOpen")
		UIErrorCheck(err)

		open, err := IsMenuItem(menuOpen)
		UIErrorCheck(err)

		// DebugStart MenuItem
		menuDebugStart, err := builder.GetObject("menuDebugStart")
		UIErrorCheck(err)

		debugStart, err := IsMenuItem(menuDebugStart)
		UIErrorCheck(err)

		debugStart.Connect("activate", func() {
			b, err := gtk.BuilderNewFromFile("ui/EmulatorWindow.glade")
			UIErrorCheck(err)

			obj, err := b.GetObject("EmulatorWindow")
			UIErrorCheck(err)

			emulatorWindow, err := IsWindow(obj)
			UIErrorCheck(err)

			gtkglarea, err := b.GetObject("GLArea")
			UIErrorCheck(err)

			glarea, err := IsGLArea(gtkglarea)
			UIErrorCheck(err)

			glarea.SetRequiredVersion(4, 6)

			go System.CPU.Run(true)

			emulatorWindow.Show()

			glarea.MakeCurrent()

			// init GL
			if err := gl.Init(); err != nil {
				panic(err)
			}
			version := gl.GoStr(gl.GetString(gl.VERSION))
			fmt.Println("OpenGL version", version)
			major, minor := glarea.GetRequiredVersion()
			fmt.Println("MajorV", major, "MinorV", minor)

		})

		// Debug MenuItem
		menuDebugObj, err := builder.GetObject("menuDebug")
		UIErrorCheck(err)

		menuDebug, err := IsMenuItem(menuDebugObj)
		UIErrorCheck(err)

		// Console
		consoleObj, err := builder.GetObject("textViewConsole")
		UIErrorCheck(err)

		console, err := IsTextView(consoleObj)
		UIErrorCheck(err)

		var consoleUI = UITextView{
			TextView: console,
		}

		multiWriter := io.MultiWriter(core.Logger.Writer(), &consoleUI)

		core.Logger.SetOutput(multiWriter)
		// Run MenuItem
		menuRunObj, err := builder.GetObject("menuRun")
		UIErrorCheck(err)

		menuRun, err := IsMenuItem(menuRunObj)
		UIErrorCheck(err)

		menuRun.Connect("activate", func() {
			go System.CPU.Run(false)
		})

		menuAbout, err := builder.GetObject("menuItemAbout")
		UIErrorCheck(err)

		menuItemAbout, err := IsMenuItem(menuAbout)
		UIErrorCheck(err)

		menuItemAbout.Connect("activate", func() {
			builder, err := gtk.BuilderNewFromFile("ui/AboutDialog.glade")
			UIErrorCheck(err)

			obj, err := builder.GetObject("aboutDialog")
			UIErrorCheck(err)

			aboutDialog, err := IsAboutDialog(obj)
			UIErrorCheck(err)

			result := aboutDialog.Run()
			if result == gtk.RESPONSE_CLOSE || result == gtk.RESPONSE_DELETE_EVENT {
				aboutDialog.Close()
			}

		})

		menuSettings, err := builder.GetObject("menuSettings")
		UIErrorCheck(err)

		menuItemSettings, err := IsMenuItem(menuSettings)
		UIErrorCheck(err)

		menuItemSettings.Connect("activate", func() {
			builder, err := gtk.BuilderNewFromFile("ui/SettingsWindow.glade")
			UIErrorCheck(err)

			obj, err := builder.GetObject("SettingsWindow")
			UIErrorCheck(err)

			settingsWindow, err := IsWindow(obj)
			UIErrorCheck(err)

			settingsWindow.Show()
		})

		romList, err := builder.GetObject("romListStore")
		UIErrorCheck(err)

		romListStore, err := IsListStore(romList)
		UIErrorCheck(err)

		romTree, err := builder.GetObject("romTreeStore")
		UIErrorCheck(err)

		romTreeStore, err := IsTreeView(romTree)
		UIErrorCheck(err)

		romProgress, err := builder.GetObject("romProgressBar")
		UIErrorCheck(err)

		romProgressBar, err := IsProgressBar(romProgress)
		UIErrorCheck(err)

		// Setup Open Menu Item
		// Builds ROM Dialog details
		// Sets up the buttons for the ROM Dialog
		open.Connect("activate", func() {
			builder, err := gtk.BuilderNewFromFile("ui/RomFileChooserDialog.glade")
			UIErrorCheck(err)

			obj, err := builder.GetObject("RomFileChooserDialog")
			UIErrorCheck(err)

			romFileChooserDialog, err := IsFileChooserDialog(obj)
			UIErrorCheck(err)

			cancel, err := builder.GetObject("buttonCancel")
			UIErrorCheck(err)

			cancelButton, err := IsButton(cancel)
			UIErrorCheck(err)

			cancelButton.Connect("clicked", func() {
				romFileChooserDialog.Response(gtk.RESPONSE_CANCEL)
			})

			open, err := builder.GetObject("buttonOpen")
			UIErrorCheck(err)

			openButton, err := IsButton(open)
			UIErrorCheck(err)

			openButton.Connect("clicked", func() {
				romFileChooserDialog.Response(gtk.RESPONSE_ACCEPT)
			})

			romFileChooserDialog.Connect("file-activated", func() {
				romFileChooserDialog.Response(gtk.RESPONSE_ACCEPT)
			})

			var result = romFileChooserDialog.Run()
			if result == gtk.RESPONSE_ACCEPT {
				ROMfile := romFileChooserDialog.GetFilename()
				core.Logger.Println(ROMfile)
				romFileChooserDialog.Close()
				// Do not block UI execution
				go System.LoadROM(string(ROMfile), romListStore, romTreeStore, romProgressBar, menuDebug, menuRun)
			} else if result == gtk.RESPONSE_CANCEL {
				core.Logger.Println("Cancelling")
				romFileChooserDialog.Close()
			}

		})

		recentROMsmenu, err := builder.GetObject("recentRoms")
		UIErrorCheck(err)

		recentROMs := recentROMsmenu.(*gtk.RecentChooserMenu)

		recentROMs.Connect("item-activated", func() {
			// gotk3 doesn't provide ALL GTK bindings, I am using the slicing operator to cut off file:///
			// This gets a uri, i.e. with url escape characters for spaces and special characters
			// strings.ReplaceAll() is required for this operation
			var romLoc string = strings.ReplaceAll(recentROMs.GetCurrentUri()[7:], "%20", " ")
			fmt.Println(romLoc)
			go System.LoadROM(romLoc, romListStore, romTreeStore, romProgressBar, menuDebug, menuRun)
		})

		menuQuit, err := builder.GetObject("menuQuit")
		UIErrorCheck(err)

		quit, err := IsMenuItem(menuQuit)
		UIErrorCheck(err)

		quit.Connect("activate", func() {
			app.Quit()
		})

		win, err := IsWindow(obj)
		UIErrorCheck(err)

		win.Show()
		app.AddWindow(win)
	})
	app.Connect("shutdown", func() {
		core.Logger.Println("FreeMe!GB is shutting down...")
	})

	app.Run(os.Args[1:])
}

// IsWindow converts a GObject to a GTK Window.
func IsWindow(obj glib.IObject) (*gtk.Window, error) {
	// Make type assertion (as per gtk.go).
	if win, ok := obj.(*gtk.Window); ok {
		return win, nil
	}
	return nil, errors.New("not a *gtk.Window")
}

// IsMenuItem converts a GObject to a GTK MenuItem.
func IsMenuItem(obj glib.IObject) (*gtk.MenuItem, error) {
	// Make type assertion (as per gtk.go).
	if item, ok := obj.(*gtk.MenuItem); ok {
		return item, nil
	}
	return nil, errors.New("not a *gtk.MenuItem")
}

// IsFileChooserDialog converts a GObject to a GTK FileChooserDialog.
func IsFileChooserDialog(obj glib.IObject) (*gtk.FileChooserDialog, error) {
	// Make type assertion (as per gtk.go).
	if dialog, ok := obj.(*gtk.FileChooserDialog); ok {
		return dialog, nil
	}
	return nil, errors.New("not a *gtk.FileChooserDialog")
}

// IsAboutDialog converts a GObject to a GTK AboutDialog.
func IsAboutDialog(obj glib.IObject) (*gtk.AboutDialog, error) {
	// Make type assertion (as per gtk.go).
	if dialog, ok := obj.(*gtk.AboutDialog); ok {
		return dialog, nil
	}
	return nil, errors.New("not a *gtk.AboutDialog")
}

// IsButton converts a GObject to a GTK Button.
func IsButton(obj glib.IObject) (*gtk.Button, error) {
	// Make type assertion (as per gtk.go).
	if button, ok := obj.(*gtk.Button); ok {
		return button, nil
	}
	return nil, errors.New("not a *gtk.Button")
}

// IsListStore converts a GObject to a GTK ListStore.
func IsListStore(obj glib.IObject) (*gtk.ListStore, error) {
	// Make type assertion (as per gtk.go).
	if item, ok := obj.(*gtk.ListStore); ok {
		return item, nil
	}
	return nil, errors.New("not a *gtk.ListStore")
}

// IsTreeView converts a GObject to a GTK TreeView.
func IsTreeView(obj glib.IObject) (*gtk.TreeView, error) {
	// Make type assertion (as per gtk.go).
	if item, ok := obj.(*gtk.TreeView); ok {
		return item, nil
	}
	return nil, errors.New("not a *gtk.TreeView")
}

// IsProgressBar converts a GObject to a GTK ProgressBar.
func IsProgressBar(obj glib.IObject) (*gtk.ProgressBar, error) {
	// Make type assertion (as per gtk.go).
	if item, ok := obj.(*gtk.ProgressBar); ok {
		return item, nil
	}
	return nil, errors.New("not a *gtk.ProgressBar")
}

// IsTextView converts a GObject to a GTK TextView.
func IsTextView(obj glib.IObject) (*gtk.TextView, error) {
	// Make type assertion (as per gtk.go).
	if item, ok := obj.(*gtk.TextView); ok {
		return item, nil
	}
	return nil, errors.New("not a *gtk.TextView")
}

// IsGLArea converts a GObject to a GTK GLArea.
func IsGLArea(obj glib.IObject) (*gtk.GLArea, error) {
	// Make type assertion (as per gtk.go).
	if item, ok := obj.(*gtk.GLArea); ok {
		return item, nil
	}
	return nil, errors.New("not a *gtk.GLArea")
}

// UIErrorCheck checks a previous Is* function for any UI errors.
func UIErrorCheck(err error) {
	if err != nil {
		// panic for any errors.
		core.Logger.Panic(err)
	}
}
