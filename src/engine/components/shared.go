package components

import (
	"log"
	"os"
	"strings"
	"time"
)

var LogFilename = strings.ReplaceAll("freemegb_"+time.Now().Format("January 01, 2006")+".log", " ", "_")

var LogFile, err = os.OpenFile(LogFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

var Logger = log.New(LogFile, "FreeMe!GB ", log.LstdFlags)

var ROMref []byte
