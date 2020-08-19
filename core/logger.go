package core

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

type LoggerType struct {
	InternalLogger *log.Logger
}

type LogType struct {
	INFO      byte
	ERROR     byte
	WARNING   byte
	COMPLETED byte
}

var LogTypes = LogType{
	INFO:      0,
	ERROR:     1,
	WARNING:   2,
	COMPLETED: 3,
}

var Logger LoggerType

func (logger *LoggerType) Log(Ltype byte, v ...interface{}) {
	switch Ltype {
	case 0: // INFO
		logger.InternalLogger.SetPrefix("<span foreground='#555555' font='bold'>[INFO]:\t\t\t</span>")
		break
	case 1: // ERROR
		logger.InternalLogger.SetPrefix("<span foreground='#770000' font='bold'>[ERROR]:\t\t</span>")
		break
	case 2: // WARNING
		logger.InternalLogger.SetPrefix("<span foreground='#00BBCC' font='bold'>[WARNING]:\t\t</span>")
		break
	case 3: // COMPLETED
		logger.InternalLogger.SetPrefix("<span foreground='#00BB00' font='bold'>[COMPLETED]:\t</span>")
		break
	}
	logger.InternalLogger.Println(v...)
	switch Ltype {
	case 0: // INFO
		c := color.New(color.FgHiBlack).Add(color.Bold)
		c.Printf("%-15s", "[INFO]:")
		fmt.Println(v...)
		break
	case 1: // ERROR
		c := color.New(color.FgHiRed).Add(color.Bold)
		c.Printf("%-15s", "[ERROR]:")
		fmt.Println(v...)
		break
	case 2: // WARNING
		c := color.New(color.FgHiYellow).Add(color.Bold)
		c.Printf("%-15s", "[WARNING]:")
		fmt.Println(v...)
		break
	case 3: // COMPLETED
		c := color.New(color.FgHiGreen).Add(color.Bold)
		c.Printf("%-15s", "[COMPLETED]:")
		fmt.Println(v...)
		break
	}
}

func (logger *LoggerType) Logf(Ltype byte, format string, v ...interface{}) {
	switch Ltype {
	case 0: // INFO
		logger.InternalLogger.SetPrefix("<span foreground='#555555' font='bold'>[INFO]:\t\t\t</span>")
		break
	case 1: // ERROR
		logger.InternalLogger.SetPrefix("<span foreground='#770000' font='bold'>[ERROR]:\t\t</span>")
		break
	case 2: // WARNING
		logger.InternalLogger.SetPrefix("<span foreground='#00BBCC' font='bold'>[WARNING]:\t\t</span>")
		break
	case 3: // COMPLETED
		logger.InternalLogger.SetPrefix("<span foreground='#00BB00' font='bold'>[COMPLETED]:\t</span>")
		break
	}
	logger.InternalLogger.Printf(format, v...)
	switch Ltype {
	case 0: // INFO
		c := color.New(color.FgHiBlack).Add(color.Bold)
		c.Printf("%-15s", "[INFO]:")
		fmt.Printf(format, v...)
		break
	case 1: // ERROR
		c := color.New(color.FgHiRed).Add(color.Bold)
		c.Printf("%-15s", "[ERROR]:")
		fmt.Printf(format, v...)
		break
	case 2: // WARNING
		c := color.New(color.FgHiYellow).Add(color.Bold)
		c.Printf("%-15s", "[WARNING]:")
		fmt.Printf(format, v...)
		break
	case 3: // COMPLETED
		c := color.New(color.FgHiGreen).Add(color.Bold)
		c.Printf("%-15s", "[COMPLETED]:")
		fmt.Printf(format, v...)
		break
	}
}

func (logger *LoggerType) Panic(v ...interface{}) {
	logger.InternalLogger.SetPrefix("<span foreground='#FF0000' font='bold'>[ERROR]</span>")
	logger.InternalLogger.Panic(v...)
	fmt.Printf("[ERROR] %s\n", v...)
}
