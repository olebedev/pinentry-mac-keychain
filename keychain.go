package main

import (
	"github.com/apex/log"
	"github.com/keybase/go-keychain"
)

func KeychainItemInsert(pin, service, account string) error {
	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetService(service)
	item.SetAccount(account)
	item.SetData([]byte(pin))
	item.SetSynchronizable(keychain.SynchronizableNo)
	item.SetAccessible(keychain.AccessibleWhenUnlocked)
	ctx := log.WithField("service", service).WithField("account", account)
	err := keychain.AddItem(item)
	if err != nil {
		ctx.WithError(err).Error("Unable to insert keychain item")
		return err
	}
	ctx.Info("Keychain item was insert")
	return nil
}

func KeychainItemGet(service, account string) (string, error) {
	query := keychain.NewItem()
	query.SetSecClass(keychain.SecClassGenericPassword)
	query.SetService(service)
	query.SetAccount(account)
	query.SetMatchLimit(keychain.MatchLimitOne)
	query.SetReturnData(true)
	results, err := keychain.QueryItem(query)
	ctx := log.WithField("service", service).WithField("account", account)
	if err != nil {
		ctx.WithError(err).Error("Unable to retrieve keychain item")
		return "", err
	} else if len(results) != 1 {
		ctx.Info("Unable to retrieve keychain item; not found")
		return "", keychain.ErrorItemNotFound
	} else {
		ctx.Info("Keychain item retrieved")
		return string(results[0].Data), nil
	}
}
