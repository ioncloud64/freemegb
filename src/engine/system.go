package engine

import (
	"time"
	"log"
	"io/ioutil"
	"engine/components"

	"github.com/gotk3/gotk3/gtk"
//	"github.com/gotk3/gotk3/glib"
)

/*
	System Structure
	================
	CPU Structure
	---> Instructions Array
	---> Registers Structure
	---> MMU Structure
	GPU Structure
	---> Graphics Processing
	---> Screen Pipeline
*/
type SystemType struct {
	CPU CPUType
	ROM ROMType
}

var System = SystemType {
	CPU: CPU,
	ROM: ROM,
}

func (system *SystemType) LoadROM(location string, romListStore *gtk.ListStore, 
	romTreeView *gtk.TreeView, romProgressBar *gtk.ProgressBar,
	menuDebug *gtk.MenuItem, menuRun *gtk.MenuItem) {
	log.Println("ROM: Loading")
	before := time.Now()
	rom, err := ioutil.ReadFile(location)
	if err == nil {
		ROM.data = rom
		ROM.BuildROMModel()
		romTreeView.SetModel(nil)
		romListStore.Clear()
		romModelLength := len(ROM.model)
		percentStep := int(float64(0.01) * float64(romModelLength))
		for i := 0; i < romModelLength; i++ {
			row		:= ROM.model[i].([]string)
			iter	:= romListStore.Append()
			err		:= romListStore.Set(iter,
										[]int{0,1},
										[]interface{} {row[0], row[1]})
			if err != nil {
				log.Println(err)
			}
			if i % percentStep == 0 {
				romProgressBar.SetFraction(float64(i) / float64(romModelLength))
//				log.Println(romProgressBar.GetFraction())
//				romProgressBar.Pulse()
			} else if i == romModelLength - 1 {
				romProgressBar.SetFraction(1)
			}
		}
		romTreeView.SetModel(romListStore)
		components.ROMref = ROM.data
	} else {
		log.Println("ROM: Error loading")
		after := time.Now()
		log.Println(after.Sub(before))
		return
	}
	after := time.Now()
	log.Println("ROM: Loaded in", after.Sub(before))
	menuDebug.SetSensitive(true)
	menuRun.SetSensitive(true)
	return
}

