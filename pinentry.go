package main

import (
	"github.com/apex/log"
	"github.com/foxcpp/go-assuan/common"
	"github.com/foxcpp/go-assuan/pinentry"
	"github.com/keybase/go-keychain"
)

// GetPIN returns a PIN code.
//
// First, it tries to get the PIN from the KeyChain and return it. If not
// found, it request it from the user via the pinentry-mac, store it in
// KeyChain and then return it back, as a result.
//
// Note: it doesn't try to interfere the requests that are not recognised as the
// standard GnuPG tool requests for a PIN code. This is done via RegExp over the
// description that is (not?)provided. Otherwise, pass the execution through,
// down to the the original `pinentry.GetPIN`.
func GetPIN(s pinentry.Settings) (string, *common.Error) {
	// Check if this is a standard GnuPG PIN request
	m := SCSNRe.FindStringSubmatch(s.Desc)
	if m == nil {
		client, err := pinentry.LaunchCustom(pmp)
		if err != nil {
			log.WithError(err).Error("Launch custom pinentry")
			return "", &common.Error{
				Message: err.Error(),
			}
		}
		fields := (log.Fields)(Apply(&client, s))
		ctx := log.WithFields(fields).WithField("gnupg", false)

		str, err := client.GetPIN()
		if err != nil {
			ctx.WithError(err).Error("GetPIN")
			return "", &common.Error{
				Message: err.Error(),
			}
		}
		ctx.Info("GetPIN")
		return str, nil
	}

	// Try to retrieve the PIN
	pin, err := KeychainItemGet(RecordName(m[1]), m[2])
	if err != nil && err != keychain.ErrorItemNotFound {
		return "", &common.Error{
			Message: err.Error(),
		}
	}

	if err == nil && pin != "" {
		return pin, nil
	}

	client, err := pinentry.LaunchCustom(pmp)
	if err != nil {
		log.WithError(err).Error("Launch custom pinentry")
		return "", &common.Error{
			Message: err.Error(),
		}
	}
	// Double check the PIN before create the record
	(&s).RepeatPrompt = "Repeat PIN"
	(&s).Desc = s.Desc + "\n\nNote: the PIN will be saved in the macOs keychain"
	fields := (log.Fields)(Apply(&client, s))
	ctx := log.WithFields(fields).WithField("gnupg", true)

	str, err := client.GetPIN()
	if err != nil {
		ctx.WithError(err).Error("GetPIN")
		return "", &common.Error{
			Message: err.Error(),
		}
	} else {
		ctx.Info("GetPIN")
	}

	err = KeychainItemInsert(str, RecordName(m[1]), m[2])
	if err != nil {
		return "", &common.Error{
			Message: err.Error(),
		}
	}

	return str, nil
}

// Confirm is a passthru function that doesn't do anything
// in addition to the original pinentry executable
func Confirm(s pinentry.Settings) (bool, *common.Error) {
	client, err := pinentry.LaunchCustom(pmp)
	if err != nil {
		log.WithError(err).Error("Launch custom pinentry")
		return false, &common.Error{
			Message: err.Error(),
		}
	}

	fields := (log.Fields)(Apply(&client, s))
	ctx := log.WithFields(fields).WithField("gnupg", false)

	err = client.Confirm()
	if err != nil {
		ctx.WithError(err).Error("Confirm")
		return false, &common.Error{
			Message: err.Error(),
		}
	}
	ctx.Info("Confirm")
	return true, nil
}

// Message is a passthru function that doesn't do anything
// in addition to the original pinentry executable
func Message(s pinentry.Settings) *common.Error {
	client, err := pinentry.LaunchCustom(pmp)
	if err != nil {
		log.WithError(err).Error("Launch custom pinentry")
		return &common.Error{
			Message: err.Error(),
		}
	}
	fields := (log.Fields)(Apply(&client, s))
	ctx := log.WithFields(fields).WithField("gnupg", false)
	client.Message()
	ctx.Info("Message")
	return nil
}
