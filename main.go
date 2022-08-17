package main

import (
	"github.com/kaskichandrakant/go-log-trace/error_handler"
	"github.com/kaskichandrakant/go-log-trace/logger"
	"log"
	"path/filepath"
)

func main() {
	abs, fileErr := filepath.Abs("./log.txt")
	if error_handler.HandleError(fileErr) {
		return
	}
	lgr, lgrErr := logger.NewCustomLogger(abs, log.Ldate|log.Ltime|log.Lshortfile)
	if error_handler.HandleError(lgrErr) {
		return
	}
	lgr.AddCustom("Custom", "CUSTOM")

	lgr.Err("logging error message")
	lgr.Info("logging info message")
	lgr.Print("Custom", "some another message with different identifier")
}
