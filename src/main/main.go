package main

import (
	"fmt"
	"engine"
//	"reflect"
	"time"
	"log"
//	"os"
	"errors"
	
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

const AppID = "com.ioncloud64.freemegb"

func main() {
	var System = engine.System
	
	System.CPU.INSTRUCTIONS[0x00].Exec()
	System.CPU.INSTRUCTIONS[0x01].Exec()
	
	System.CPU.REGISTERS.Print()

	System.CPU.REGISTERS.PrintRegister16(System.CPU.REGISTERS.BC)
	System.CPU.REGISTERS.PrintRegister8(System.CPU.REGISTERS.B())
	
	fmt.Printf("0x%04X\n\n", System.CPU.REGISTERS.BC);
	System.CPU.REGISTERS.SetB(0x17)
	fmt.Printf("0x%04X\n\n", System.CPU.REGISTERS.BC);
	
	finished := make(chan bool)
	UI()
	
	fmt.Println("Pausing for 5 seconds...")
	time.Sleep(5 * time.Second)
	go System.CPU.Run(finished)
	fmt.Println("Waiting for System to finish...")
	
	<- finished
	
	fmt.Println("System finished!")
	
}

func UI() {
	app, err := gtk.ApplicationNew(AppID, glib.APPLICATION_FLAGS_NONE)
	UIErrorCheck(err)
	
	app.Connect("startup", func() {
			log.Println("FreeMe!GB is starting up...")
	})
	app.Connect("activate", func() {
			log.Println("FreeMe!GB  is activated...")
			
			builder, err := gtk.BuilderNewFromFile("ui/MainWindow.glade")
			UIErrorCheck(err)

			obj, err := builder.GetObject("MainWindow")
			UIErrorCheck(err)
			
			menuOpen, err := builder.GetObject("menuOpen")
			UIErrorCheck(err)
			
			open, err := IsMenuItem(menuOpen)
			UIErrorCheck(err)

			
			
			
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
					
					cancelButton.Connect("clicked", func(){
							romFileChooserDialog.Response(gtk.RESPONSE_CANCEL)
					})
					
					open, err := builder.GetObject("buttonOpen")
					UIErrorCheck(err)
					
					openButton, err := IsButton(open)
					UIErrorCheck(err)
					
					openButton.Connect("clicked", func(){
							romFileChooserDialog.Response(gtk.RESPONSE_ACCEPT)
					})

					var result = romFileChooserDialog.Run()
					if result == gtk.RESPONSE_ACCEPT {
						log.Println(romFileChooserDialog.GetFilename())
						romFileChooserDialog.Close()
					} else if result == gtk.RESPONSE_CANCEL {
						romFileChooserDialog.Close()
					}
					
			})
			
			win, err := IsWindow(obj)
			UIErrorCheck(err)
			
			win.Show()
			app.AddWindow(win)
	})
	app.Connect("shutdown", func() {
			log.Println("Closing FreeMe!GB")
	})

	
	go app.Run(nil)
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

func UIErrorCheck(e error) {
	if e != nil {
		// panic for any errors.
		log.Panic(e)
	}
}

