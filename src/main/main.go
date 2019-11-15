package main

import (
	"fmt"
	"engine"
//	"reflect"
	"time"
	"log"
//	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

var (
	topLabel    *gtk.Label
	bottomLabel *gtk.Label
	nSets       = 1
)

func UI() {	
	gtk.Init(nil)
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	win.Add(windowWidget())
	
	win.ShowAll()
	go gtk.Main()
}

func windowWidget() *gtk.Widget {

	return &grid.Container.Widget
}



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
