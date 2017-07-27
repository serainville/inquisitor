package printer

import (
	"github.com/serainville/gologger/logger"
	"github.com/serainville/gologger/printer/console"
	"github.com/serainville/gologger/printer/es"
	"github.com/serainville/gologger/printer/file"
	"time"
)

func Print(log logger.LogInstance, packageName string, fileName string, lineNumber int, funcName string, time time.Time) {
	if log.LoggerInit.PrinterType == "console" {
		console.ConsolePrinter(log, packageName, fileName, lineNumber, funcName, time)
	} else if log.LoggerInit.PrinterType == "basic-console" {
		console.ConsoleBasicPrinter(log, time)
	} else if log.LoggerInit.PrinterType == "file" {
		file.FilePrinter(log, packageName, fileName, lineNumber, funcName, time)
	} else if log.LoggerInit.PrinterType == "es" {
		es.ESPrinter(log, packageName, fileName, lineNumber, funcName, time)
	}
}
