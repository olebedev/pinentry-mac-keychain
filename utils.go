package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/apex/log"
	"github.com/foxcpp/go-assuan/pinentry"
)

// Smart card serial number RegExp that matches the standard PIN request by GnuPG tool.
// https://go.dev/play/p/_mqQxYDdSGm
var SCSNRe = regexp.MustCompile("(?i)^" +
	"(?:Please\\sunlock\\sthe\\scard\\n\\nNumber\\:\\s)" +
	"([\\d\\s]*)" +
	"(?:\\nHolder\\:\\s(.+))?" +
	"(?:\\W|$)")

func RecordName(pin string) string {
	return fmt.Sprintf("smartcard/%s", strings.Replace(pin, " ", "", -1))
}

// TODO: create a PR to github.com/foxcpp/go-assuan/pinentry and link it here
func Apply(c *pinentry.Client, s pinentry.Settings) map[string]interface{} {
	fields := make(log.Fields)
	if s.Desc != "" {
		fields["Desc"] = s.Desc
		c.SetDesc(s.Desc)
	}
	if s.Prompt != "" {
		fields["Prompt"] = s.Prompt
		c.SetPrompt(s.Prompt)
	}
	if s.Error != "" {
		fields["Error"] = s.Error
		c.SetError(s.Error)
	}
	if s.OkBtn != "" {
		fields["OkBtn"] = s.OkBtn
		c.SetOkBtn(s.OkBtn)
	}
	if s.NotOkBtn != "" {
		fields["NotOkBtn"] = s.NotOkBtn
		c.SetNotOkBtn(s.NotOkBtn)
	}
	if s.CancelBtn != "" {
		fields["CancelBtn"] = s.CancelBtn
		c.SetCancelBtn(s.CancelBtn)
	}
	if s.Title != "" {
		fields["Title"] = s.Title
		c.SetTitle(s.Title)
	}
	if s.Timeout != 0 {
		fields["Timeout"] = s.Timeout
		c.SetTimeout(s.Timeout)
	}
	if s.RepeatPrompt != "" {
		fields["RepeatPrompt"] = s.RepeatPrompt
		c.SetRepeatPrompt(s.RepeatPrompt)
	}
	if s.RepeatError != "" {
		fields["RepeatError"] = s.RepeatError
		c.SetRepeatError(s.RepeatError)
	}
	if s.QualityBar != "" {
		fields["QualityBar"] = s.QualityBar
		c.SetQualityBar(s.QualityBar)
	}
	return fields
}
