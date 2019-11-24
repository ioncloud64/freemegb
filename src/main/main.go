package main

import (
	//	"fmt"
	"engine"
	//	"reflect"
	"errors"
	"log"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

const AppID = "com.ioncloud64.freemegb"

func main() {
	var System = engine.System

	UI(&System)
}

func UI(System *engine.SystemType) {
	app, err := gtk.ApplicationNew(AppID, glib.APPLICATION_FLAGS_NONE)
	UIErrorCheck(err)

	app.Connect("startup", func() {
		log.Println("FreeMe!GB is starting up...")
	})
	app.Connect("activate", func() {
		log.Println("FreeMe!GB is activated...")

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
			go System.CPU.Run(true)
		})

		// Debug MenuItem
		menuDebugObj, err := builder.GetObject("menuDebug")
		UIErrorCheck(err)

		menuDebug, err := IsMenuItem(menuDebugObj)
		UIErrorCheck(err)

		// Run MenuItem
		menuRunObj, err := builder.GetObject("menuRun")
		UIErrorCheck(err)

		menuRun, err := IsMenuItem(menuRunObj)
		UIErrorCheck(err)

		menuRun.Connect("activate", func() {
			go System.CPU.Run(false)
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

			var result = romFileChooserDialog.Run()
			if result == gtk.RESPONSE_ACCEPT {
				ROMfile := romFileChooserDialog.GetFilename()
				log.Println(ROMfile)
				romFileChooserDialog.Close()
				// Do not block UI execution
				go System.LoadROM(string(ROMfile), romListStore, romTreeStore, romProgressBar, menuDebug, menuRun)
			} else if result == gtk.RESPONSE_CANCEL {
				log.Println("Cancelling")
				romFileChooserDialog.Close()
			}

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
		log.Println("FreeMe!GB is shutting down...")
	})
	
	app.Run(nil)
}

func IsWindow(obj glib.IObject) (*gtk.Window, error) {
	// Make type assertion (as per gtk.go).
	if win, ok := obj.(*gtk.Window); ok {
		return win, nil
	}
	return nil, errors.New("not a *gtk.Window")
}

func IsMenuItem(obj glib.IObject) (*gtk.MenuItem, error) {
	// Make type assertion (as per gtk.go).
	if item, ok := obj.(*gtk.MenuItem); ok {
		return item, nil
	}
	return nil, errors.New("not a *gtk.MenuItem")
}

func IsFileChooserDialog(obj glib.IObject) (*gtk.FileChooserDialog, error) {
	// Make type assertion (as per gtk.go).
	if dialog, ok := obj.(*gtk.FileChooserDialog); ok {
		return dialog, nil
	}
	return nil, errors.New("not a *gtk.FileChooserDialog")
}

func IsButton(obj glib.IObject) (*gtk.Button, error) {
	// Make type assertion (as per gtk.go).
	if button, ok := obj.(*gtk.Button); ok {
		return button, nil
	}
	return nil, errors.New("not a *gtk.Button")
}

func IsListStore(obj glib.IObject) (*gtk.ListStore, error) {
	// Make type assertion (as per gtk.go).
	if item, ok := obj.(*gtk.ListStore); ok {
		return item, nil
	}
	return nil, errors.New("not a *gtk.ListStore")
}

func IsTreeView(obj glib.IObject) (*gtk.TreeView, error) {
	// Make type assertion (as per gtk.go).
	if item, ok := obj.(*gtk.TreeView); ok {
		return item, nil
	}
	return nil, errors.New("not a *gtk.TreeView")
}

func IsProgressBar(obj glib.IObject) (*gtk.ProgressBar, error) {
	// Make type assertion (as per gtk.go).
	if item, ok := obj.(*gtk.ProgressBar); ok {
		return item, nil
	}
	return nil, errors.New("not a *gtk.ProgressBar")
}

func UIErrorCheck(err error) {
	if err != nil {
		// panic for any errors.
		log.Panic(err)
	}
}

func SetupUI(*gtk.Window) error {
	return nil
}
