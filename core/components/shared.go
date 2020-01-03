package components

import (
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gotk3/gotk3/gtk"
)

var UserHome, err = os.UserHomeDir()

var LogFilename string

var LogFile *os.File

var Logger *log.Logger

var ROMref []byte

var AppRef *gtk.Application

func Init() {
	os.MkdirAll(path.Join(UserHome,".freemegb"), os.FileMode(0755))
	LogFilename = UserHome+"/.freemegb/"+strings.ReplaceAll("freemegb_"+time.Now().Format("January 2, 2006")+".log", " ", "_")
	LogFile, err = os.OpenFile(LogFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	Logger = log.New(LogFile, "FreeMe!GB ", log.LstdFlags)
}
