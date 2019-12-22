package github.com/ioncloud64/freemegb/engine/components

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/gotk3/gotk3/gtk"
)

var LogFilename = strings.ReplaceAll("freemegb_"+time.Now().Format("January 2, 2006")+".log", " ", "_")

var LogFile, err = os.OpenFile(LogFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

var Logger = log.New(LogFile, "FreeMe!GB ", log.LstdFlags)

var ROMref []byte

var AppRef *gtk.Application
