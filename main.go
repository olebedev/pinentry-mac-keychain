package main

import (
	"fmt"
	"io"
	"os"

	"github.com/foxcpp/go-assuan/pinentry"

	"github.com/apex/log"
	"github.com/apex/log/handlers/discard"
	"github.com/apex/log/handlers/json"
)

const pmp = "/usr/local/MacGPG2/libexec/pinentry-mac.app/Contents/MacOS/pinentry-mac"

// File path to write debug info to, use
//
//	go install -ldflags "-X main.logfile=$HOME/pinentry.log"
//
// to build with debug output.
var logfile = ""

func main() {
	var err error

	// Logging setup
	if logfile != "" {
		var f *os.File
		f, err = os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		defer func() {
			err = f.Close()
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}()
		log.SetHandler(json.New(f))
	} else {
		log.SetHandler(discard.New())
	}

	cbks := pinentry.Callbacks{
		// Where the hijacking happens
		GetPIN: GetPIN,

		// Pass through functions
		Confirm: Confirm,
		Msg:     Message,
	}

	err = pinentry.Serve(cbks, "pinentry-map-keychain")
	if err != nil && err != io.EOF {
		log.WithError(err).Error("Failed to run pinentry.Serve")
		os.Exit(1)
	}
}
